package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetDataSomeKeyPathHandlerFunc turns a function with the right signature into a get data some key path handler
type GetDataSomeKeyPathHandlerFunc func(GetDataSomeKeyPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataSomeKeyPathHandlerFunc) Handle(params GetDataSomeKeyPathParams) middleware.Responder {
	return fn(params)
}

// GetDataSomeKeyPathHandler interface for that can handle valid get data some key path params
type GetDataSomeKeyPathHandler interface {
	Handle(GetDataSomeKeyPathParams) middleware.Responder
}

// NewGetDataSomeKeyPath creates a new http.Handler for the get data some key path operation
func NewGetDataSomeKeyPath(ctx *middleware.Context, handler GetDataSomeKeyPathHandler) *GetDataSomeKeyPath {
	return &GetDataSomeKeyPath{Context: ctx, Handler: handler}
}

/*GetDataSomeKeyPath swagger:route GET /data/someKeyPath getDataSomeKeyPath

returns the config data added

whenever Director needs to retrieve a value it will use GET action.


*/
type GetDataSomeKeyPath struct {
	Context *middleware.Context
	Handler GetDataSomeKeyPathHandler
}

func (o *GetDataSomeKeyPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDataSomeKeyPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}