package goapidoc

import (
	"regexp"
	"strings"
)

type apiTypeKind int8

const (
	apiPrimeKind apiTypeKind = iota + 1
	apiObjectKind
	apiArrayKind
)

// Example: Obj<integer#int64, string[]>[]
type apiType struct {
	name string
	kind apiTypeKind

	prime  *apiPrime
	object *apiObject
	array  *apiArray
}

// Example: string or string#date-time
type apiPrime struct {
	typ    string
	format string
}

// Example: xxx or xxx<yyy, zzz>
type apiObject struct {
	typ      string
	generics []*apiType
}

// Example: xxx[][]
type apiArray struct {
	item *apiType
}

var (
	typeNameRe    = regexp.MustCompile(`^[a-zA-Z0-9_]+(?:(?:<(.+)>)|(?:#[a-zA-Z0-9\-_]*))?(?:\[])*$`) // xxx(?:<(yyy)>|#zzz)?(?:[])*
	genericNameRe = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

// checkTypeName checks given type name from Param, Response and Definition.
func checkTypeName(typ string) {
	// re
	if !typeNameRe.MatchString(typ) {
		panic("Invalid type `" + typ + "`")
	}
	// <(.+)>
	for _, subTyp := range typeNameRe.FindStringSubmatch(typ)[1:] {
		genParts := strings.Split(subTyp, ",")
		genStrings := make([]string, 0, 2) // cap defaults to 2
		for idx, part := range genParts {
			if strings.Count(part, "<") == strings.Count(part, ">") {
				// T1, T2<T3>
				genStrings = append(genStrings, part)
			} else {
				// T1, T2<T3, T4> => T1 | T2<T3 | T4> => T1 | T2<T3,T4>
				if idx+1 < len(genParts) {
					genParts[idx+1] = genParts[idx] + "," + genParts[idx+1]
				} else {
					panic("Invalid type `" + typ + "`")
				}
			}
		}
		for _, genTyp := range genStrings {
			if genTyp = strings.TrimSpace(genTyp); genTyp != "" {
				checkTypeName(genTyp)
			}
		}
	}
}

// parseApiType parses type string to three kinds of apiType.
func parseApiType(typ string) *apiType {
	typ = strings.TrimSpace(typ)

	// 1. array: X[] | X[][][]
	if strings.HasSuffix(typ, "[]") {
		item := parseApiType(typ[:len(typ)-2]) // X | X[][]
		return &apiType{
			name:  typ,
			kind:  apiArrayKind,
			array: &apiArray{item: item},
		}
	}

	// 2. prime with format: X# | X#Y
	if strings.Contains(typ, "#") {
		for _, prime := range []string{INTEGER, NUMBER, STRING, BOOLEAN, FILE, ARRAY, OBJECT} {
			if strings.HasPrefix(typ, prime) {
				if prime == ARRAY || prime == OBJECT {
					panic("Use array or object in type `" + typ + "` invalidly")
				}
				fmtIdx := strings.Index(typ, "#") + 1
				fmt := ""
				if fmtIdx < len(typ) {
					fmt = typ[fmtIdx:]
				}
				return &apiType{
					name:  typ,
					kind:  apiPrimeKind,
					prime: &apiPrime{typ: prime, format: fmt},
				}
			}
		}
	}

	// 3. object with generic: X<Y<Z>> | X<Y, Z<A, B<C>>>
	if strings.Contains(typ, "<") {
		genIdx := strings.Index(typ, "<") + 1
		genParts := strings.Split(typ[genIdx:len(typ)-1], ",") // Y<Z> || {Y, Z<A, B<C>>}
		genStrings := make([]string, 0, 2)                     // cap defaults to 2
		for idx, part := range genParts {
			part = strings.TrimSpace(part)
			if strings.Count(part, "<") == strings.Count(part, ">") { // Y<Z>
				genStrings = append(genStrings, part)
			} else { // Z<A, B<C>>
				if idx+1 < len(genParts) {
					genParts[idx+1] = genParts[idx] + "," + genParts[idx+1] // -> Z<A + , + B<C>>
				} else {
					panic("Invalid type `" + typ + "`")
				}
			}
		}
		genTypes := make([]*apiType, 0, len(genStrings))
		for _, gen := range genStrings {
			genTypes = append(genTypes, parseApiType(gen))
		}
		object := typ[:genIdx-1]
		switch object {
		case INTEGER, NUMBER, STRING, BOOLEAN, FILE, ARRAY, OBJECT:
			panic("Invalid type `" + typ + "`")
		}
		return &apiType{
			name:   typ,
			kind:   apiObjectKind,
			object: &apiObject{typ: object, generics: genTypes},
		}
	}

	// 4. prime without format: X
	switch typ {
	case INTEGER, NUMBER, STRING, BOOLEAN, FILE:
		return &apiType{
			name:  typ,
			kind:  apiPrimeKind,
			prime: &apiPrime{typ: typ, format: defaultFormat(typ)},
		}
	case ARRAY, OBJECT:
		panic("Use array or object as type invalidly")
	}

	// 5. object without generic: X
	if strings.Contains(typ, "#") {
		panic("Invalid type `" + typ + "`")
	}
	return &apiType{
		name:   typ,
		kind:   apiObjectKind,
		object: &apiObject{typ: typ, generics: []*apiType{}},
	}
}

// defaultFormat returns the default format for given type.
func defaultFormat(typ string) string {
	if typ == INTEGER {
		return INT32
	}
	if typ == NUMBER {
		return DOUBLE
	}
	return ""
}

// prehandleDefinition prehandles and returns a cloned Definition with new generic type name.
func prehandleDefinition(definition *Definition) *Definition {
	// check and deduplicate generic names
	generics := make([]string, 0, len(definition.generics))
	for _, gen := range definition.generics {
		contained := false
		for _, newGen := range generics {
			if newGen == gen {
				contained = true
				break
			}
		}
		if !contained {
			if !genericNameRe.MatchString(gen) { // a-zA-Z0-9_
				panic("Invalid generic type `" + gen + "`")
			}
			generics = append(generics, gen)
		}
	}

	// clone definition
	out := &Definition{
		name:       definition.name,
		desc:       definition.desc,
		generics:   generics,
		properties: make([]*Property, 0, len(definition.properties)),
	}
	for _, prop := range definition.properties {
		out.properties = append(out.properties, &Property{
			name:       prop.name,
			typ:        prop.typ,
			required:   prop.required,
			desc:       prop.desc,
			allowEmpty: prop.allowEmpty,
			defaul:     prop.defaul,
			example:    prop.example,
			enums:      prop.enums,
			minLength:  prop.minLength,
			maxLength:  prop.maxLength,
			minimum:    prop.minimum,
			maximum:    prop.maximum,
		})
	}

	// update generic type name
	for idx, gen := range out.generics {
		newGen := "«" + gen + "»"
		re := regexp.MustCompile(`(^|[,\s<])` + gen + `([,\s<>\[]|$)`) // {,\s<>[]} => {,\s<} xxx {,\s<>[}
		for _, prop := range out.properties {
			for { // replace all
				curr := prop.typ
				prop.typ = strings.ReplaceAll(prop.typ, " ", "")
				prop.typ = re.ReplaceAllString(prop.typ, "$1"+newGen+"$2") // T -> «T»
				prop.typ = strings.ReplaceAll(prop.typ, ",", ", ")
				if prop.typ == curr {
					break
				}
			}
		}
		out.generics[idx] = newGen
	}
	return out
}

// prehandleDefinitionList prehandles and returns final Definition list with given definition list and type list.
func prehandleDefinitionList(allDefinitions []*Definition, allTypes []string) []*Definition {
	// extract generic definitions from given definitions
	genericDefs := make(map[string]*Definition)
	finalMap := newOrderedMap(len(allDefinitions)) // map[string]*Definition
	for _, definition := range allDefinitions {
		if len(definition.generics) != 0 {
			genericDefs[definition.name] = definition
		} else {
			finalMap.Set(definition.name, definition) // plain definition
		}
	}

	// extract more definitions from given types
	var extractFn func(typ string)
	extractFn = func(typ string) {
		at := parseApiType(typ)
		for at.kind == apiArrayKind {
			at = at.array.item
		}
		if at.kind != apiObjectKind || len(at.object.generics) == 0 {
			return
		}

		// check and get object with generic
		genDef, ok := genericDefs[at.object.typ]
		if !ok {
			panic("Invalid type `" + typ + "`") // type has no generic defined
		}
		if len(at.object.generics) != len(genDef.generics) {
			panic("Invalid type `" + typ + "`") // generic parameter length not matched
		}

		// new definition need to be added
		newDef := &Definition{
			name:       genDef.name, // TypeName<GenericName, ...>
			desc:       genDef.desc,
			generics:   make([]string, 0), // empty
			properties: make([]*Property, 0, len(genDef.properties)),
		}
		for _, prop := range genDef.properties {
			newDef.properties = append(newDef.properties, &Property{
				name:       prop.name,
				typ:        prop.typ, // << need to extract recurrently
				required:   prop.required,
				desc:       prop.desc,
				allowEmpty: prop.allowEmpty,
				defaul:     prop.defaul,
				example:    prop.example,
				enums:      prop.enums,
				minLength:  prop.minLength,
				maxLength:  prop.maxLength,
				minimum:    prop.minimum,
				maximum:    prop.maximum,
			})
		}

		// replace to spec name for new definition
		specNames := make([]string, 0, len(genDef.generics))
		for idx, genName := range genDef.generics {
			specName := at.object.generics[idx].name
			specNames = append(specNames, specName)
			for _, prop := range newDef.properties {
				prop.typ = strings.ReplaceAll(prop.typ, genName, specName) // «T» -> XXX, replace directly
			}
		}
		newDef.name += "<" + strings.Join(specNames, ", ") + ">" // TypeName -> TypeName<GenericName, ...>

		// extract recurrently and append to finalMap
		for _, prop := range newDef.properties {
			if !finalMap.Has(prop.typ) {
				extractFn(prop.typ) // << extract property type recurrently
			}
		}
		finalMap.Set(newDef.name, newDef)
	}
	for _, typ := range allTypes {
		extractFn(typ)
	}

	// convert ordered-map to definition slice
	outDefs := make([]*Definition, 0, len(finalMap.Keys()))
	for _, key := range finalMap.Keys() {
		def := finalMap.MustGet(key).(*Definition)
		outDefs = append(outDefs, def)
	}
	return outDefs
}
