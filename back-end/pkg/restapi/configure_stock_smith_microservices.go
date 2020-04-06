// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/endpoints"

	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/restapi/operations"
	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/restapi/operations/news_api"
	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/restapi/operations/stock_api"
)

//go:generate swagger generate server --target ..\..\pkg --name StockSmithMicroservices --spec ..\..\idl\swagger.yaml
var scrapers []NewsScraper

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

	if api.NewsAPIGetStocksmithAPINewsQueryHandler == nil {
		api.NewsAPIGetStocksmithAPINewsQueryHandler = news_api.GetStocksmithAPINewsQueryHandlerFunc(endpoints.HandleNewsQuery)	
	}

	if api.StockAPIGetStocksmithAPIStocksQueryStockHandler == nil {
		api.StockAPIGetStocksmithAPIStocksQueryStockHandler = stock_api.GetStocksmithAPIStocksQueryStockHandlerFunc(endpoints.HandleStockQueryStock)
	}
	if api.StockAPIGetStocksmithAPIStocksQueryTopHandler == nil {
		api.StockAPIGetStocksmithAPIStocksQueryTopHandler = stock_api.GetStocksmithAPIStocksQueryTopHandlerFunc(endpoints.HandleStockQueryTop)
	}

	api.PreServerShutdown = handlePreServerShutdown

	api.ServerShutdown = handleServerShutdown

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func handlePreServerShutdown() {

}

func handleServerShutdown() {

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
