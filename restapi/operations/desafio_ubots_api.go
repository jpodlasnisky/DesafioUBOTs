// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
	"std/github.com/jpodlasnisky/DesafioUBOTs/restapi/operations/cliente"
	"std/github.com/jpodlasnisky/DesafioUBOTs/restapi/operations/vinho"
)

// NewDesafioUbotsAPI creates a new DesafioUbots instance
func NewDesafioUbotsAPI(spec *loads.Document) *DesafioUbotsAPI {
	return &DesafioUbotsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ClienteClienteMaiorCompraUnicaHandler: cliente.ClienteMaiorCompraUnicaHandlerFunc(func(params cliente.ClienteMaiorCompraUnicaParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation cliente.ClienteMaiorCompraUnica has not yet been implemented")
		}),
		ClienteClientesGastoTotalHandler: cliente.ClientesGastoTotalHandlerFunc(func(params cliente.ClientesGastoTotalParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation cliente.ClientesGastoTotal has not yet been implemented")
		}),
		ClienteClientesMaisFieisHandler: cliente.ClientesMaisFieisHandlerFunc(func(params cliente.ClientesMaisFieisParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation cliente.ClientesMaisFieis has not yet been implemented")
		}),
		VinhoRecomendacaoVinhoHandler: vinho.RecomendacaoVinhoHandlerFunc(func(params vinho.RecomendacaoVinhoParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation vinho.RecomendacaoVinho has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		APIKeyAuth: func(token string) (*models.Token, error) {
			return nil, errors.NotImplemented("api key auth (api_key) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*DesafioUbotsAPI Essa API foi desenvolvida para atender os requisitos solicitados no desafio prático da UBOTs. */
type DesafioUbotsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	APIKeyAuth func(string) (*models.Token, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ClienteClienteMaiorCompraUnicaHandler sets the operation handler for the cliente maior compra unica operation
	ClienteClienteMaiorCompraUnicaHandler cliente.ClienteMaiorCompraUnicaHandler
	// ClienteClientesGastoTotalHandler sets the operation handler for the clientes gasto total operation
	ClienteClientesGastoTotalHandler cliente.ClientesGastoTotalHandler
	// ClienteClientesMaisFieisHandler sets the operation handler for the clientes mais fieis operation
	ClienteClientesMaisFieisHandler cliente.ClientesMaisFieisHandler
	// VinhoRecomendacaoVinhoHandler sets the operation handler for the recomendacao vinho operation
	VinhoRecomendacaoVinhoHandler vinho.RecomendacaoVinhoHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *DesafioUbotsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DesafioUbotsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DesafioUbotsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DesafioUbotsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DesafioUbotsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DesafioUbotsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DesafioUbotsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DesafioUbotsAPI
func (o *DesafioUbotsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.ClienteClienteMaiorCompraUnicaHandler == nil {
		unregistered = append(unregistered, "cliente.ClienteMaiorCompraUnicaHandler")
	}
	if o.ClienteClientesGastoTotalHandler == nil {
		unregistered = append(unregistered, "cliente.ClientesGastoTotalHandler")
	}
	if o.ClienteClientesMaisFieisHandler == nil {
		unregistered = append(unregistered, "cliente.ClientesMaisFieisHandler")
	}
	if o.VinhoRecomendacaoVinhoHandler == nil {
		unregistered = append(unregistered, "vinho.RecomendacaoVinhoHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DesafioUbotsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DesafioUbotsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.APIKeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *DesafioUbotsAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *DesafioUbotsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *DesafioUbotsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DesafioUbotsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the desafio ubots API
func (o *DesafioUbotsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DesafioUbotsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clienteMaiorCompraUnica"] = cliente.NewClienteMaiorCompraUnica(o.context, o.ClienteClienteMaiorCompraUnicaHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clientesGastoTotal"] = cliente.NewClientesGastoTotal(o.context, o.ClienteClientesGastoTotalHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clientesMaisFieis/{tamanhoLista}"] = cliente.NewClientesMaisFieis(o.context, o.ClienteClientesMaisFieisHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/recomendacaoVinho/{cpfCliente}"] = vinho.NewRecomendacaoVinho(o.context, o.VinhoRecomendacaoVinhoHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DesafioUbotsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *DesafioUbotsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DesafioUbotsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DesafioUbotsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *DesafioUbotsAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
