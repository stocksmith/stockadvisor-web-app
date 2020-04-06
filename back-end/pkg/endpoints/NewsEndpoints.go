package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/restapi/operations/news_api"
)

func HandleNewsQuery(params news_api.GetStocksmithAPINewsQueryParams) middleware.Responder {
	return middleware.NotImplemented("operation news_api.GetStocksmithAPINewsQuery has not yet been implemented")
}