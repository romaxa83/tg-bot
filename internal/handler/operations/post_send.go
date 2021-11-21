// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostSendHandlerFunc turns a function with the right signature into a post send handler
type PostSendHandlerFunc func(PostSendParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSendHandlerFunc) Handle(params PostSendParams) middleware.Responder {
	return fn(params)
}

// PostSendHandler interface for that can handle valid post send params
type PostSendHandler interface {
	Handle(PostSendParams) middleware.Responder
}

// NewPostSend creates a new http.Handler for the post send operation
func NewPostSend(ctx *middleware.Context, handler PostSendHandler) *PostSend {
	return &PostSend{Context: ctx, Handler: handler}
}

/* PostSend swagger:route POST /send postSend

Send message to all subscribers

*/
type PostSend struct {
	Context *middleware.Context
	Handler PostSendHandler
}

func (o *PostSend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostSendParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
