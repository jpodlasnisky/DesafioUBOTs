// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"std/github.com/jpodlasnisky/DesafioUBOTs/casadevinhos"
	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
	"std/github.com/jpodlasnisky/DesafioUBOTs/restapi/operations"
	"std/github.com/jpodlasnisky/DesafioUBOTs/restapi/operations/cliente"
	"std/github.com/jpodlasnisky/DesafioUBOTs/restapi/operations/vinho"
)

//go:generate swagger generate server --target ..\..\DesafioUBOTs --name DesafioUbots --spec ..\swagger.json --principal models.Token

func configureFlags(api *operations.DesafioUbotsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DesafioUbotsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	// Executa quando no header existe a "api_key"
	api.APIKeyAuth = func(token string) (*models.Token, error) {
		if token == "chavedeseguranca" {
			prin := models.Token(token)
			return &prin, nil
		}
		return nil, errors.New(401, "Unauthorized")
	}
	api.ClienteClientesMaisFieisHandler = cliente.ClientesMaisFieisHandlerFunc(func(params cliente.ClientesMaisFieisParams, principal *models.Token) middleware.Responder {
		return cliente.NewClientesMaisFieisOK().WithPayload(casadevinhos.RetornaClientesMaisFieis())
	})

	api.ClienteClienteMaiorCompraUnicaHandler = cliente.ClienteMaiorCompraUnicaHandlerFunc(func(params cliente.ClienteMaiorCompraUnicaParams, principal *models.Token) middleware.Responder {
		return cliente.NewClienteMaiorCompraUnicaOK().WithPayload(casadevinhos.RetornaMaiorCompraUnica())
	})
	// clientesGastoTotal
	api.ClienteClientesGastoTotalHandler = cliente.ClientesGastoTotalHandlerFunc(func(params cliente.ClientesGastoTotalParams, principal *models.Token) middleware.Responder {
		return cliente.NewClientesGastoTotalOK().WithPayload(casadevinhos.OrdenaClientesMaiorValorTotalEmCompras())
	})

	api.VinhoRecomendacaoVinhoHandler = vinho.RecomendacaoVinhoHandlerFunc(func(params vinho.RecomendacaoVinhoParams, principal *models.Token) middleware.Responder {
		return vinho.NewRecomendacaoVinhoOK().WithPayload(casadevinhos.RetornaVinhoMaisCompradoPeloCliente(params.CpfCliente))
	})

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.ClienteClienteMaiorCompraUnicaHandler == nil {
		api.ClienteClienteMaiorCompraUnicaHandler = cliente.ClienteMaiorCompraUnicaHandlerFunc(func(params cliente.ClienteMaiorCompraUnicaParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation cliente.ClienteMaiorCompraUnica has not yet been implemented")
		})
	}
	if api.ClienteClientesGastoTotalHandler == nil {
		api.ClienteClientesGastoTotalHandler = cliente.ClientesGastoTotalHandlerFunc(func(params cliente.ClientesGastoTotalParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation cliente.ClientesGastoTotal has not yet been implemented")
		})
	}
	if api.ClienteClientesMaisFieisHandler == nil {
		api.ClienteClientesMaisFieisHandler = cliente.ClientesMaisFieisHandlerFunc(func(params cliente.ClientesMaisFieisParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation cliente.ClientesMaisFieis has not yet been implemented")
		})
	}
	if api.VinhoRecomendacaoVinhoHandler == nil {
		api.VinhoRecomendacaoVinhoHandler = vinho.RecomendacaoVinhoHandlerFunc(func(params vinho.RecomendacaoVinhoParams, principal *models.Token) middleware.Responder {
			return middleware.NotImplemented("operation vinho.RecomendacaoVinho has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
