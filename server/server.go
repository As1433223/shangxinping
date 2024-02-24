package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"server/logic"
	"server/proto"
)

func Server() {
	GrpcServer := grpc.NewServer()
	server := &logic.ServerRpc{}
	proto.RegisterServerRpcServer(GrpcServer, server)
	Listen, err := net.Listen("tcp", ":5566")
	if err != nil {
		panic(err)
	}
	grpc_health_v1.RegisterHealthServer(GrpcServer, health.NewServer())
	GrpcServer.Serve(Listen)
}
