@echo off

mkdir pkg
mkdir pkg\client

swagger generate server -f idl\swagger.yaml -t pkg
swagger generate client -f idl\swagger.yaml -t pkg