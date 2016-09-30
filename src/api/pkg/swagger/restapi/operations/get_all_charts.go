package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllChartsHandlerFunc turns a function with the right signature into a get all charts handler
type GetAllChartsHandlerFunc func(GetAllChartsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllChartsHandlerFunc) Handle(params GetAllChartsParams) middleware.Responder {
	return fn(params)
}

// GetAllChartsHandler interface for that can handle valid get all charts params
type GetAllChartsHandler interface {
	Handle(GetAllChartsParams) middleware.Responder
}

// NewGetAllCharts creates a new http.Handler for the get all charts operation
func NewGetAllCharts(ctx *middleware.Context, handler GetAllChartsHandler) *GetAllCharts {
	return &GetAllCharts{Context: ctx, Handler: handler}
}

/*GetAllCharts swagger:route GET /v1/charts getAllCharts

get all charts from all repos

*/
type GetAllCharts struct {
	Context *middleware.Context
	Handler GetAllChartsHandler
}

func (o *GetAllCharts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetAllChartsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
