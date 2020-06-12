package grpcserver

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/linkingthing/pg-agent/pkg/grpcservice"
	"github.com/linkingthing/pg-agent/pkg/proto"
)

const ListenAddr = "0.0.0.0:8888"

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

func New() (*GRPCServer, error) {
	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		return nil, fmt.Errorf("create listener with addr 0.0.0.0:8888 failed: %s", err.Error())
	}

	grpcServer := &GRPCServer{
		server:   grpc.NewServer(),
		listener: listener,
	}

	proto.RegisterPGManagerServer(grpcServer.server, grpcservice.New())
	return grpcServer, nil
}

func (s *GRPCServer) Run() error {
	return s.server.Serve(s.listener)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
