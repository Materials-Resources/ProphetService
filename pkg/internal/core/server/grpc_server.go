package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const defaultHost = "0.0.0.0"

type GrpcServer interface {
	Start()
	Stop()
}

type grpcServer struct {
	Listener net.Listener
	server   *grpc.Server
}

func NewGrpcServer(server *grpc.Server) GrpcServer {
	listener, err := net.Listen(
		"tcp",
		"0.0.0.0:50058",
	)
	if err != nil {
		log.Fatalf(
			"could not create listner: %s\n",
			err.Error(),
		)
		return nil
	}

	return &grpcServer{
		Listener: listener,
		server:   server,
	}
}

func (s grpcServer) Start() {
	if err := s.server.Serve(s.Listener); err != nil {
		log.Fatalf(
			"failed to start grpc server: %s\n",
			err.Error(),
		)
	}
	log.Printf("listening on...")
}

func (s grpcServer) Stop() {

}
