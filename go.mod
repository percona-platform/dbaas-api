module github.com/percona-platform/dbaas-api

go 1.14

// some dependecies should be synced with Dockerfile

require (
	github.com/golang/protobuf v1.4.3
	github.com/mwitkow/go-proto-validators v0.3.2
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
