package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/restapi/operations/stock_api"
)

func HandleStockQueryStock(params stock_api.GetStocksmithAPIStocksQueryStockParams) middleware.Responder {
	return middleware.NotImplemented("operation stock_api.GetStocksmithAPIStocksQueryStock has not yet been implemented")
}


func HandleStockQueryTop(params stock_api.GetStocksmithAPIStocksQueryTopParams) middleware.Responder {
	return middleware.NotImplemented("operation stock_api.GetStocksmithAPIStocksQueryTop has not yet been implemented")
}