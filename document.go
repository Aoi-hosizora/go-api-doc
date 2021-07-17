package goapidoc

// ========
// Document
// ========

// Document represents an api document information.
type Document struct {
	host     string
	basePath string
	info     *Info

	option      *Option
	operations  []*Operation
	definitions []*Definition
}

// NewDocument creates a default Document with given arguments.
func NewDocument(host, basePath string, info *Info) *Document {
	return &Document{host: host, basePath: basePath, info: info}
}

func (d *Document) GetHost() string               { return d.host }
func (d *Document) GetBasePath() string           { return d.basePath }
func (d *Document) GetInfo() *Info                { return d.info }
func (d *Document) GetOption() *Option            { return d.option }
func (d *Document) GetOperations() []*Operation   { return d.operations }
func (d *Document) GetDefinitions() []*Definition { return d.definitions }

// Cleanup cleans up Document.
func (d *Document) Cleanup() *Document {
	d.host = ""
	d.basePath = ""
	d.info = nil
	d.option = nil
	d.operations = nil
	d.definitions = nil
	return d
}

// Host sets the host in Document.
func (d *Document) Host(host string) *Document {
	d.host = host
	return d
}

// BasePath sets the basePath in Document.
func (d *Document) BasePath(basePath string) *Document {
	d.basePath = basePath
	return d
}

// Info sets the info in Document.
func (d *Document) Info(info *Info) *Document {
	d.info = info
	return d
}

// Option sets the option in Document.
func (d *Document) Option(option *Option) *Document {
	d.option = option
	return d
}

// Operations sets the whole operations in Document.
func (d *Document) Operations(operations ...*Operation) *Document {
	d.operations = operations
	return d
}

// AddOperations adds some operations into Document.
func (d *Document) AddOperations(operations ...*Operation) *Document {
	d.operations = append(d.operations, operations...)
	return d
}

// Definitions sets the whole definitions in Document.
func (d *Document) Definitions(definitions ...*Definition) *Document {
	d.definitions = definitions
	return d
}

// AddDefinitions adds some definitions into Document.
func (d *Document) AddDefinitions(definitions ...*Definition) *Document {
	d.definitions = append(d.definitions, definitions...)
	return d
}

// ====
// Info
// ====

// Info represents a basic api information of Document.
type Info struct {
	title   string
	desc    string
	version string

	termsOfService string
	license        *License
	contact        *Contact
}

// NewInfo creates a default Info with given arguments.
func NewInfo(title, desc, version string) *Info {
	return &Info{title: title, desc: desc, version: version}
}

func (i *Info) GetTitle() string          { return i.title }
func (i *Info) GetDesc() string           { return i.desc }
func (i *Info) GetVersion() string        { return i.version }
func (i *Info) GetTermsOfService() string { return i.termsOfService }
func (i *Info) GetLicense() *License      { return i.license }
func (i *Info) GetContact() *Contact      { return i.contact }

// Title sets the title in Info.
func (i *Info) Title(title string) *Info {
	i.title = title
	return i
}

// Desc sets the desc in Info.
func (i *Info) Desc(desc string) *Info {
	i.desc = desc
	return i
}

// Version sets the version in Info.
func (i *Info) Version(version string) *Info {
	i.version = version
	return i
}

// TermsOfService sets the termsOfService in Info.
func (i *Info) TermsOfService(terms string) *Info {
	i.termsOfService = terms
	return i
}

// License sets the license in Info.
func (i *Info) License(license *License) *Info {
	i.license = license
	return i
}

// Contact sets the contact in Info.
func (i *Info) Contact(contact *Contact) *Info {
	i.contact = contact
	return i
}

// =======
// License
// =======

// License represents an api license information of Document.
type License struct {
	name string
	url  string
}

// NewLicense creates a default License with given arguments.
func NewLicense(name, url string) *License {
	return &License{name: name, url: url}
}

func (l *License) GetName() string { return l.name }
func (l *License) GetUrl() string  { return l.url }

// Name sets the name in License.
func (l *License) Name(name string) *License {
	l.name = name
	return l
}

// Url sets the url in License.
func (l *License) Url(url string) *License {
	l.url = url
	return l
}

// =======
// Contact
// =======

// Contact represents an api contact information of Document.
type Contact struct {
	name  string
	url   string
	email string
}

// NewContact creates a default Contact with given arguments.
func NewContact(name, url, email string) *Contact {
	return &Contact{name: name, url: url, email: email}
}

func (c *Contact) GetName() string  { return c.name }
func (c *Contact) GetUrl() string   { return c.url }
func (c *Contact) GetEmail() string { return c.email }

// Name sets the name in Contact.
func (c *Contact) Name(name string) *Contact {
	c.name = name
	return c
}

// Url sets the url in Contact.
func (c *Contact) Url(url string) *Contact {
	c.url = url
	return c
}

// Email sets the email in Contact.
func (c *Contact) Email(email string) *Contact {
	c.email = email
	return c
}

// ======
// Option
// ======

// Option represents an extra options of Document.
// TODO BREAK CHANGES
type Option struct {
	schemes              []string
	consumes             []string
	produces             []string
	tags                 []*Tag
	securities           []*Security
	externalDocs         *ExternalDocs
	additionalDoc        string
	routesAliases        map[string]string
	routesAdditionalDocs map[string]string
}

// NewOption creates an empty document Option.
func NewOption() *Option {
	return &Option{}
}

func (o *Option) GetSchemes() []string                       { return o.schemes }
func (o *Option) GetConsumes() []string                      { return o.consumes }
func (o *Option) GetProduces() []string                      { return o.produces }
func (o *Option) GetTags() []*Tag                            { return o.tags }
func (o *Option) GetSecurities() []*Security                 { return o.securities }
func (o *Option) GetExternalDocs() *ExternalDocs             { return o.externalDocs }
func (o *Option) GetAdditionalDoc() string                   { return o.additionalDoc }
func (o *Option) GetRoutesAliases() map[string]string        { return o.routesAliases }
func (o *Option) GetRoutesAdditionalDocs() map[string]string { return o.routesAdditionalDocs }

// Schemes sets the whole schemes in Option.
func (o *Option) Schemes(schemes ...string) *Option {
	o.schemes = schemes
	return o
}

// AddSchemes adds some tags schemes into Option.
func (o *Option) AddSchemes(schemes ...string) *Option {
	o.schemes = append(o.schemes, schemes...)
	return o
}

// Consumes sets the whole consumes in Option.
func (o *Option) Consumes(consumes ...string) *Option {
	o.consumes = consumes
	return o
}

// AddConsumes adds some consumes into Option.
func (o *Option) AddConsumes(consumes ...string) *Option {
	o.consumes = append(o.consumes, consumes...)
	return o
}

// Produces sets the whole produces in Option.
func (o *Option) Produces(produces ...string) *Option {
	o.produces = produces
	return o
}

// AddProduces adds some produces into Option.
func (o *Option) AddProduces(produces ...string) *Option {
	o.produces = append(o.produces, produces...)
	return o
}

// Tags sets the whole tags in Option.
func (o *Option) Tags(tags ...*Tag) *Option {
	o.tags = tags
	return o
}

// AddTags adds some tags into Option.
func (o *Option) AddTags(tags ...*Tag) *Option {
	o.tags = append(o.tags, tags...)
	return o
}

// Securities sets the whole securities in Option.
func (o *Option) Securities(securities ...*Security) *Option {
	o.securities = securities
	return o
}

// AddSecurities adds some securities into Option.
func (o *Option) AddSecurities(securities ...*Security) *Option {
	o.securities = append(o.securities, securities...)
	return o
}

// ExternalDocs sets the externalDocs in Option.
func (o *Option) ExternalDocs(docs *ExternalDocs) *Option {
	o.externalDocs = docs
	return o
}

// AdditionalDoc sets the additional document in Option, this is only supported in API Blueprint.
func (o *Option) AdditionalDoc(doc string) *Option {
	o.additionalDoc = doc
	return o
}

// RoutesAliases sets the whole routes' aliases in Option, this is only supported in API Blueprint.
func (o *Option) RoutesAliases(aliases map[string]string) *Option {
	o.routesAliases = aliases
	return o
}

// AddRoutesAlias adds a routes' alias in Option, this is only supported in API Blueprint.
func (o *Option) AddRoutesAlias(route, alias string) *Option {
	if o.routesAliases == nil {
		o.routesAliases = make(map[string]string, 0)
	}
	o.routesAliases[route] = alias
	return o
}

// RoutesAdditionalDocs sets the routes' additional documents in Option, this is only supported in API Blueprint.
func (o *Option) RoutesAdditionalDocs(aliases map[string]string) *Option {
	o.routesAdditionalDocs = aliases
	return o
}

// AddRoutesAlias adds a routes' additional document in Option, this is only supported in API Blueprint.
func (o *Option) AddRoutesAdditionalDoc(route, doc string) *Option {
	if o.routesAdditionalDocs == nil {
		o.routesAdditionalDocs = make(map[string]string, 0)
	}
	o.routesAdditionalDocs[route] = doc
	return o
}

// ===
// Tag
// ===

// Tag represents an api tag information of Document.
type Tag struct {
	name         string
	desc         string
	externalDocs *ExternalDocs
}

// NewTag creates a default Tag with given arguments.
func NewTag(name, desc string) *Tag {
	return &Tag{name: name, desc: desc}
}

func (t *Tag) GetName() string                { return t.name }
func (t *Tag) GetDesc() string                { return t.desc }
func (t *Tag) GetExternalDocs() *ExternalDocs { return t.externalDocs }

// Name sets the name in Tag.
func (t *Tag) Name(name string) *Tag {
	t.name = name
	return t
}

// Desc sets the desc in Tag.
func (t *Tag) Desc(desc string) *Tag {
	t.desc = desc
	return t
}

// ExternalDocs sets the externalDocs in Tag.
func (t *Tag) ExternalDocs(docs *ExternalDocs) *Tag {
	t.externalDocs = docs
	return t
}

// ========
// Security
// ========

// Security represents an api security definition information of Document.
type Security struct {
	title string
	typ   string
	desc  string

	in   string // only for apiKey
	name string // only for apiKey

	flow             string            // only for oauth2
	authorizationUrl string            // only for oauth2
	tokenUrl         string            // only for oauth2
	scopes           map[string]string // only for oauth2
}

// NewSecurity creates a default Security with given arguments.
// TODO BREAK CHANGES
func NewSecurity(title string, typ string) *Security {
	return &Security{title: title, typ: typ}
}

// NewApiKeySecurity creates an apiKey authentication Security with given arguments.
func NewApiKeySecurity(title, in, name string) *Security {
	return &Security{title: title, typ: APIKEY, in: in, name: name}
}

// NewBasicSecurity creates a basic authentication Security with given arguments.
func NewBasicSecurity(title string) *Security {
	return &Security{title: title, typ: BASIC}
}

// NewOAuth2Security creates an oauth2 authentication Security with given arguments.
func NewOAuth2Security(title string, flow string) *Security {
	return &Security{title: title, typ: OAUTH2, flow: flow}
}

func (s *Security) GetTitle() string             { return s.title }
func (s *Security) GetType() string              { return s.typ }
func (s *Security) GetDesc() string              { return s.desc }
func (s *Security) GetInLoc() string             { return s.in }
func (s *Security) GetName() string              { return s.name }
func (s *Security) GetFlow() string              { return s.flow }
func (s *Security) GetAuthorizationUrl() string  { return s.authorizationUrl }
func (s *Security) GetTokenUrl() string          { return s.tokenUrl }
func (s *Security) GetScopes() map[string]string { return s.scopes }

// Title sets the title in Security.
func (s *Security) Title(title string) *Security {
	s.title = title
	return s
}

// Type sets the authentication type in Security.
func (s *Security) Type(typ string) *Security {
	s.typ = typ
	return s
}

// Desc sets the desc in Security.
func (s *Security) Desc(desc string) *Security {
	s.desc = desc
	return s
}

// InLoc sets the in-location in Security.
func (s *Security) InLoc(in string) *Security {
	s.in = in
	return s
}

// Name sets the name in Security.
func (s *Security) Name(name string) *Security {
	s.name = name
	return s
}

// Flow sets the flow in Security.
func (s *Security) Flow(flow string) *Security {
	s.flow = flow
	return s
}

// AuthorizationUrl sets the authorizationUrl in Security.
func (s *Security) AuthorizationUrl(authorizationUrl string) *Security {
	s.authorizationUrl = authorizationUrl
	return s
}

// TokenUrl sets the tokenUrl in Security.
func (s *Security) TokenUrl(tokenUrl string) *Security {
	s.tokenUrl = tokenUrl
	return s
}

// Scopes sets the whole scopes in Security.
func (s *Security) Scopes(scopes map[string]string) *Security {
	s.scopes = scopes
	return s
}

// AddScope add a scope into Security.
func (s *Security) AddScope(scope, desc string) *Security {
	if s.scopes == nil {
		s.scopes = make(map[string]string)
	}
	s.scopes[scope] = desc
	return s
}

// ============
// ExternalDocs
// ============

// ExternalDocs represents an additional external documentation of Document.
type ExternalDocs struct {
	desc string
	url  string
}

// NewExternalDocs creates a default ExternalDocs with given arguments.
func NewExternalDocs(desc, url string) *ExternalDocs {
	return &ExternalDocs{desc: desc, url: url}
}

func (e *ExternalDocs) GetDesc() string { return e.desc }
func (e *ExternalDocs) GetUrl() string  { return e.url }

// Desc sets the desc in ExternalDocs.
func (e *ExternalDocs) Desc(desc string) *ExternalDocs {
	e.desc = desc
	return e
}

// Url sets the url in ExternalDocs.
func (e *ExternalDocs) Url(url string) *ExternalDocs {
	e.url = url
	return e
}

// ===============
// global document
// ===============

// _document represents a global Document.
var _document = NewDocument("", "", nil)

// SetDocument sets the basic information for global Document.
func SetDocument(host, basePath string, info *Info) *Document {
	_document.host = host
	_document.basePath = basePath
	_document.info = info
	return _document
}

func GetHost() string               { return _document.GetHost() }
func GetBasePath() string           { return _document.GetBasePath() }
func GetInfo() *Info                { return _document.GetInfo() }
func GetOption() *Option            { return _document.GetOption() }
func GetOperations() []*Operation   { return _document.GetOperations() }
func GetDefinitions() []*Definition { return _document.GetDefinitions() }

// CleanupDocument cleans up Document.
func CleanupDocument() *Document {
	return _document.Cleanup()
}

// SetHost sets the host in global Document.
func SetHost(host string) *Document {
	return _document.Host(host)
}

// SetBasePath sets the basePath in global Document.
func SetBasePath(basePath string) *Document {
	return _document.BasePath(basePath)
}

// SetInfo sets the info in global Document.
func SetInfo(info *Info) *Document {
	return _document.Info(info)
}

// SetOption sets the option in Document.
func SetOption(option *Option) *Document {
	return _document.Option(option)
}

// SetOperations sets the whole operations in global Document.
func SetOperations(operations ...*Operation) *Document {
	return _document.Operations(operations...)
}

// AddOperations adds some operations into global Document.
func AddOperations(operations ...*Operation) *Document {
	return _document.AddOperations(operations...)
}

// SetDefinitions sets the whole definitions in global Document.
func SetDefinitions(definitions ...*Definition) *Document {
	return _document.Definitions(definitions...)
}

// AddDefinitions adds some definitions into global Document.
func AddDefinitions(definitions ...*Definition) *Document {
	return _document.AddDefinitions(definitions...)
}
