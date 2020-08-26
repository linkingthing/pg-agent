package grpcserver

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/linkingthing/pg-agent/pkg/grpcservice"
	"github.com/linkingthing/pg-agent/pkg/proto"
	"github.com/linkingthing/pg-agent/pkg/util"
)

const ListenAddr = "0.0.0.0:8888"

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

func New() (*GRPCServer, error) {
	if err := runRsyncDaemon(); err != nil {
		return nil, fmt.Errorf("run rsync daemon failed: %s", err.Error())
	}

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

func runRsyncDaemon() error {
	if _, err := os.Stat("/var/run/rsyncd.pid"); os.IsNotExist(err) == false {
		util.ExecCommand("kill -9 $(cat /var/run/rsyncd.pid)")
		util.ExecCommand("rm -rf /var/run/rsyncd.pid")
	}

	return util.ExecCommand("rsync --daemon")
}
