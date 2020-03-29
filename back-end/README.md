# Notes
- Every file inside of back-end/pkg/ will be overwritten when generating the server and the client except for back-end/pkg/server/restapi/configure_stock_smith_microservices.go
- To generate the server and client, install [go-swagger]{https://github.com/go-swagger/go-swagger}

## Generating the server
- from stockadvisor-web-app/back-end, run: `swagger generate server -f idl\swagger.yaml -t pkg\server`

## Generating the client
- from stockadvisor-web-app/back-end, run: `swagger generate client -f idl\swagger.yaml -t pkg\client`

## Building the client and the server
- from stockadvisor-web-app/back-end, run: `go build ./...`
