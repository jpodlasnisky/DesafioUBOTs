// Code generated by go-swagger; DO NOT EDIT.

package cliente

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
)

// ClienteMaiorCompraUnicaHandlerFunc turns a function with the right signature into a cliente maior compra unica handler
type ClienteMaiorCompraUnicaHandlerFunc func(ClienteMaiorCompraUnicaParams, *models.Token) middleware.Responder

// Handle executing the request and returning a response
func (fn ClienteMaiorCompraUnicaHandlerFunc) Handle(params ClienteMaiorCompraUnicaParams, principal *models.Token) middleware.Responder {
	return fn(params, principal)
}

// ClienteMaiorCompraUnicaHandler interface for that can handle valid cliente maior compra unica params
type ClienteMaiorCompraUnicaHandler interface {
	Handle(ClienteMaiorCompraUnicaParams, *models.Token) middleware.Responder
}

// NewClienteMaiorCompraUnica creates a new http.Handler for the cliente maior compra unica operation
func NewClienteMaiorCompraUnica(ctx *middleware.Context, handler ClienteMaiorCompraUnicaHandler) *ClienteMaiorCompraUnica {
	return &ClienteMaiorCompraUnica{Context: ctx, Handler: handler}
}

/*ClienteMaiorCompraUnica swagger:route GET /clienteMaiorCompraUnica cliente clienteMaiorCompraUnica

Retorna o cliente com maior valor gasto em uma única compra

*/
type ClienteMaiorCompraUnica struct {
	Context *middleware.Context
	Handler ClienteMaiorCompraUnicaHandler
}

func (o *ClienteMaiorCompraUnica) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewClienteMaiorCompraUnicaParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Token
	if uprinc != nil {
		principal = uprinc.(*models.Token) // this is really a models.Token, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
