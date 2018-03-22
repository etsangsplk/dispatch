///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vmware/dispatch/pkg/function-manager/gen/restapi/operations/runner"
	"github.com/vmware/dispatch/pkg/function-manager/gen/restapi/operations/store"
)

// NewFunctionManagerAPI creates a new FunctionManager instance
func NewFunctionManagerAPI(spec *loads.Document) *FunctionManagerAPI {
	return &FunctionManagerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		StoreAddFunctionHandler: store.AddFunctionHandlerFunc(func(params store.AddFunctionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StoreAddFunction has not yet been implemented")
		}),
		StoreDeleteFunctionHandler: store.DeleteFunctionHandlerFunc(func(params store.DeleteFunctionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StoreDeleteFunction has not yet been implemented")
		}),
		StoreGetFunctionHandler: store.GetFunctionHandlerFunc(func(params store.GetFunctionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StoreGetFunction has not yet been implemented")
		}),
		RunnerGetFunctionRunsHandler: runner.GetFunctionRunsHandlerFunc(func(params runner.GetFunctionRunsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation RunnerGetFunctionRuns has not yet been implemented")
		}),
		StoreGetFunctionsHandler: store.GetFunctionsHandlerFunc(func(params store.GetFunctionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StoreGetFunctions has not yet been implemented")
		}),
		RunnerGetRunHandler: runner.GetRunHandlerFunc(func(params runner.GetRunParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation RunnerGetRun has not yet been implemented")
		}),
		RunnerGetRunsHandler: runner.GetRunsHandlerFunc(func(params runner.GetRunsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation RunnerGetRuns has not yet been implemented")
		}),
		RunnerRunFunctionHandler: runner.RunFunctionHandlerFunc(func(params runner.RunFunctionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation RunnerRunFunction has not yet been implemented")
		}),
		StoreUpdateFunctionHandler: store.UpdateFunctionHandlerFunc(func(params store.UpdateFunctionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation StoreUpdateFunction has not yet been implemented")
		}),

		// Applies when the "Cookie" header is set
		CookieAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (cookie) Cookie from header param [Cookie] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*FunctionManagerAPI VMware Dispatch Function Manager
 */
type FunctionManagerAPI struct {
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
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// CookieAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Cookie provided in the header
	CookieAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// StoreAddFunctionHandler sets the operation handler for the add function operation
	StoreAddFunctionHandler store.AddFunctionHandler
	// StoreDeleteFunctionHandler sets the operation handler for the delete function operation
	StoreDeleteFunctionHandler store.DeleteFunctionHandler
	// StoreGetFunctionHandler sets the operation handler for the get function operation
	StoreGetFunctionHandler store.GetFunctionHandler
	// RunnerGetFunctionRunsHandler sets the operation handler for the get function runs operation
	RunnerGetFunctionRunsHandler runner.GetFunctionRunsHandler
	// StoreGetFunctionsHandler sets the operation handler for the get functions operation
	StoreGetFunctionsHandler store.GetFunctionsHandler
	// RunnerGetRunHandler sets the operation handler for the get run operation
	RunnerGetRunHandler runner.GetRunHandler
	// RunnerGetRunsHandler sets the operation handler for the get runs operation
	RunnerGetRunsHandler runner.GetRunsHandler
	// RunnerRunFunctionHandler sets the operation handler for the run function operation
	RunnerRunFunctionHandler runner.RunFunctionHandler
	// StoreUpdateFunctionHandler sets the operation handler for the update function operation
	StoreUpdateFunctionHandler store.UpdateFunctionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *FunctionManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FunctionManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FunctionManagerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FunctionManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FunctionManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FunctionManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FunctionManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FunctionManagerAPI
func (o *FunctionManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CookieAuth == nil {
		unregistered = append(unregistered, "CookieAuth")
	}

	if o.StoreAddFunctionHandler == nil {
		unregistered = append(unregistered, "store.AddFunctionHandler")
	}

	if o.StoreDeleteFunctionHandler == nil {
		unregistered = append(unregistered, "store.DeleteFunctionHandler")
	}

	if o.StoreGetFunctionHandler == nil {
		unregistered = append(unregistered, "store.GetFunctionHandler")
	}

	if o.RunnerGetFunctionRunsHandler == nil {
		unregistered = append(unregistered, "runner.GetFunctionRunsHandler")
	}

	if o.StoreGetFunctionsHandler == nil {
		unregistered = append(unregistered, "store.GetFunctionsHandler")
	}

	if o.RunnerGetRunHandler == nil {
		unregistered = append(unregistered, "runner.GetRunHandler")
	}

	if o.RunnerGetRunsHandler == nil {
		unregistered = append(unregistered, "runner.GetRunsHandler")
	}

	if o.RunnerRunFunctionHandler == nil {
		unregistered = append(unregistered, "runner.RunFunctionHandler")
	}

	if o.StoreUpdateFunctionHandler == nil {
		unregistered = append(unregistered, "store.UpdateFunctionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FunctionManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FunctionManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "cookie":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.CookieAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *FunctionManagerAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *FunctionManagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
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

// ProducersFor gets the producers for the specified media types
func (o *FunctionManagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
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
func (o *FunctionManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the function manager API
func (o *FunctionManagerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FunctionManagerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/function"] = store.NewAddFunction(o.context, o.StoreAddFunctionHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/function/{functionName}"] = store.NewDeleteFunction(o.context, o.StoreDeleteFunctionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/function/{functionName}"] = store.NewGetFunction(o.context, o.StoreGetFunctionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/function/{functionName}/runs"] = runner.NewGetFunctionRuns(o.context, o.RunnerGetFunctionRunsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/function"] = store.NewGetFunctions(o.context, o.StoreGetFunctionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runs/{runName}"] = runner.NewGetRun(o.context, o.RunnerGetRunHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runs"] = runner.NewGetRuns(o.context, o.RunnerGetRunsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/function/{functionName}/runs"] = runner.NewRunFunction(o.context, o.RunnerRunFunctionHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/function/{functionName}"] = store.NewUpdateFunction(o.context, o.StoreUpdateFunctionHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FunctionManagerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *FunctionManagerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *FunctionManagerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *FunctionManagerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
