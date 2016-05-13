package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteSingleClusterHandlerFunc turns a function with the right signature into a delete single cluster handler
type DeleteSingleClusterHandlerFunc func(DeleteSingleClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSingleClusterHandlerFunc) Handle(params DeleteSingleClusterParams) middleware.Responder {
	return fn(params)
}

// DeleteSingleClusterHandler interface for that can handle valid delete single cluster params
type DeleteSingleClusterHandler interface {
	Handle(DeleteSingleClusterParams) middleware.Responder
}

// NewDeleteSingleCluster creates a new http.Handler for the delete single cluster operation
func NewDeleteSingleCluster(ctx *middleware.Context, handler DeleteSingleClusterHandler) *DeleteSingleCluster {
	return &DeleteSingleCluster{Context: ctx, Handler: handler}
}

/*DeleteSingleCluster swagger:route DELETE /clusters/{name} clusters deleteSingleCluster

Delete the specified cluster

*/
type DeleteSingleCluster struct {
	Context *middleware.Context
	Handler DeleteSingleClusterHandler
}

func (o *DeleteSingleCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteSingleClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}