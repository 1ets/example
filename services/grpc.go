package services

import (
	"example/app/adapters/protobuf"
	"example/app/adapters/servers"

	"google.golang.org/grpc"
)

// func RouteGrpcService(gs *grpc.Server) {
// 	protobuf.RegisterApiAccountServer(gs, &servers.ApiAccountServer{})
// }

// Route gRPC into server adapter
func GrpcRouter(route *grpc.Server) {
	protobuf.RegisterExampleServiceServer(route, &servers.GrpcExample{})
}
