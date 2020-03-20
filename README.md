# Proxy Server

An example proxy server microservice for evaluating Skaffold. 

## Development

`skaffold dev --default-repo <path-to-registry>`

### Compile Protos
`protoc -I todo/ todo/todo.proto --go_out=plugins=grpc:todo`