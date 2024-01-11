/*
 * Service
 *
 * API service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package example

import (
	"net/http"
	"strings"
)

// ResourceAPIController binds http requests to an api service and writes the service results to the http response
type ResourceAPIController struct {
	service      ResourceAPIServicer
	errorHandler ErrorHandler
}

// ResourceAPIOption for how the controller is set up.
type ResourceAPIOption func(*ResourceAPIController)

// WithResourceAPIErrorHandler inject ErrorHandler into controller
func WithResourceAPIErrorHandler(h ErrorHandler) ResourceAPIOption {
	return func(c *ResourceAPIController) {
		c.errorHandler = h
	}
}

// NewResourceAPIController creates a default api controller
func NewResourceAPIController(s ResourceAPIServicer, opts ...ResourceAPIOption) Router {
	controller := &ResourceAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ResourceAPIController
func (c *ResourceAPIController) Routes() Routes {
	return Routes{
		"Create": Route{
			strings.ToUpper("Post"),
			"/resource",
			c.Create,
		},
		"Delete": Route{
			strings.ToUpper("Delete"),
			"/resource",
			c.Delete,
		},
	}
}

// Create - Create a resource
func (c *ResourceAPIController) Create(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Create(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Delete - Delete a resource
func (c *ResourceAPIController) Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var idParam string
	if query.Has("id") {
		param := query.Get("id")

		idParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "id"}, nil)
		return
	}
	result, err := c.service.Delete(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}