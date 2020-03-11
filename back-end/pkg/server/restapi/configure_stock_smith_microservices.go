// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"pkg/server/restapi/operations"
	"pkg/server/restapi/operations/news_api"
	"pkg/server/restapi/operations/stock_api"
)

//go:generate swagger generate server --target ../../server --name StockSmithMicroservices --spec ../../../idl/swagger.yaml

func configureFlags(api *operations.StockSmithMicroservicesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.StockSmithMicroservicesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.StockAPIGetStocksQueryStockHandler == nil {
		api.StockAPIGetStocksQueryStockHandler = stock_api.GetStocksQueryStockHandlerFunc(func(params stock_api.GetStocksQueryStockParams) middleware.Responder {
			return middleware.NotImplemented("operation stock_api.GetStocksQueryStock has not yet been implemented")
		})
	}
	if api.StockAPIGetStocksQueryTopHandler == nil {
		api.StockAPIGetStocksQueryTopHandler = stock_api.GetStocksQueryTopHandlerFunc(func(params stock_api.GetStocksQueryTopParams) middleware.Responder {
			return middleware.NotImplemented("operation stock_api.GetStocksQueryTop has not yet been implemented")
		})
	}
	if api.NewsAPIPostNewsQueryHandler == nil {
		api.NewsAPIPostNewsQueryHandler = news_api.PostNewsQueryHandlerFunc(func(params news_api.PostNewsQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation news_api.PostNewsQuery has not yet been implemented")
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
