package main

import (
	"github.com/linkingthing/pg-agent/pkg/grpcserver"
)

func main() {
	s, err := grpcserver.New()
	if err != nil {
		panic("new grpc server failed: " + err.Error())
	}

	if err := s.Run(); err != nil {
		s.Stop()
		panic("run grpc server failed: " + err.Error())
	}
}
