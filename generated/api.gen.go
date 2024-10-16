// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// EstateDronePlanResponse defines model for EstateDronePlanResponse.
type EstateDronePlanResponse struct {
	Distance int `json:"distance"`
}

// EstateRequest defines model for EstateRequest.
type EstateRequest struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}

// EstateResponse defines model for EstateResponse.
type EstateResponse struct {
	Id string `json:"id"`
}

// EstateStatsResponse defines model for EstateStatsResponse.
type EstateStatsResponse struct {
	Count  int `json:"count"`
	Max    int `json:"max"`
	Median int `json:"median"`
	Min    int `json:"min"`
}

// EstateTreeRequest defines model for EstateTreeRequest.
type EstateTreeRequest struct {
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

// EstateTreeResponse defines model for EstateTreeResponse.
type EstateTreeResponse struct {
	Id string `json:"id"`
}

// PostEstateJSONRequestBody defines body for PostEstate for application/json ContentType.
type PostEstateJSONRequestBody = EstateRequest

// PostEstateIdTreeJSONRequestBody defines body for PostEstateIdTree for application/json ContentType.
type PostEstateIdTreeJSONRequestBody = EstateTreeRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new estate
	// (POST /estate)
	PostEstate(ctx echo.Context) error
	// Get dron plan for the estate
	// (GET /estate/{id}/drone-plan)
	GetEstateIdDronePlan(ctx echo.Context, id string) error
	// Get stats of estate
	// (GET /estate/{id}/stats)
	GetEstateIdStats(ctx echo.Context, id string) error
	// Create a tree data for estate
	// (POST /estate/{id}/tree)
	PostEstateIdTree(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostEstate converts echo context to params.
func (w *ServerInterfaceWrapper) PostEstate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEstate(ctx)
	return err
}

// GetEstateIdDronePlan converts echo context to params.
func (w *ServerInterfaceWrapper) GetEstateIdDronePlan(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEstateIdDronePlan(ctx, id)
	return err
}

// GetEstateIdStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetEstateIdStats(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEstateIdStats(ctx, id)
	return err
}

// PostEstateIdTree converts echo context to params.
func (w *ServerInterfaceWrapper) PostEstateIdTree(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEstateIdTree(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/estate", wrapper.PostEstate)
	router.GET(baseURL+"/estate/:id/drone-plan", wrapper.GetEstateIdDronePlan)
	router.GET(baseURL+"/estate/:id/stats", wrapper.GetEstateIdStats)
	router.POST(baseURL+"/estate/:id/tree", wrapper.PostEstateIdTree)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RVTU/rOhD9K9G8t0xJ+kGfyJJ3EeriSgjYIRbGnjZGiW3sCVBV/e9XtttCSgNBAl2x",
	"i5LjzJkzx2dWwHVttEJFDooVOF5izcLjmbXaXqIzWjn0L4zVBi1JDJ9rdI4twgdaGoQCHFmpFrBep2Dx",
	"oZEWBRQ3O+BtugXqu3vkBOsUzhwxwl9WK7yomOquJqQjpnj4gs+sNhVCMcrz3T+lIlygfVN9d7C7/CU+",
	"NOjobdEK1YLKVsnhgYopPEnRA7fHbPP37en3+HWpIkWrKDA+PZlPxHTApnw6mBxP/hvcDUcnAzbmYza9",
	"y49ZPof0g3lJ8Q6XK2Lkuglx3ShqcTqoV82ee4BQSKZ64OSHoL0OI8tII57fVevu/Npit1NKlIuy3fj4",
	"INnnHm5aftZJvoslpFsWH/XwSTcNvZtGX+omD5JqrsMlkxw3dBSrPer37NpTJkmBxhV7kmSshhQe0Tqp",
	"FRQwPMqPco/SBhUzEgoYh1cpGEZl6CfD0HRoVMep+XYZSa1mAgq40I6iMBBJo6NTLZbRyYowepkZU0ke",
	"jmX3TquXnPRP/1qcQwH/ZC9Bmm1SNGvny7qtDdkGw4s4j0B5lOdfXnwz7lBdoONWGooiRkTCLTJCkbiG",
	"c3Ru3lTV0ks7+UoyrWVygMspE4ndCZWCa+qa2SUU8H+gl7BE4VOyGalHbMabraRYZ8KvkIGpYmAs8MC0",
	"z3Ez7JnYLZxgF8tqJLQOipsV+DAJFoJ060cpYH9u6au29/1/++0zfbsvDwgaQIlXpDXYxCI1VqHYU/kc",
	"KfEixhNzbRMqsVNu/+j6KB32xc9Uub3quq9P0KK3xhGt553SksVekTUTPs2/TdrvCsPXa/SvBGJrB/78",
	"UPR2SQQjFu7szlQejPZxa4rGVlBASWSKLKs0Z1Xp7bW+Xf8JAAD//xr//isBDAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
