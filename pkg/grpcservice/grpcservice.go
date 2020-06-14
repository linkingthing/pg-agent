package grpcservice

import (
	"context"

	pb "github.com/linkingthing/pg-agent/pkg/proto"
)

type PGService struct {
	handler *PGHandler
}

func New() *PGService {
	return &PGService{
		handler: newPGHandler(),
	}
}

func (s *PGService) UpdatePostgresqlConf(ctx context.Context, req *pb.UpdatePostgresqlConfRequest) (*pb.PGResponse, error) {
	if err := s.handler.UpdatePostgresqlConf(req); err != nil {
		return &pb.PGResponse{Succeed: false}, err
	} else {
		return &pb.PGResponse{Succeed: true}, nil
	}
}

func (s *PGService) UpdatePGHBAConf(ctx context.Context, req *pb.UpdatePGHBAConfRequest) (*pb.PGResponse, error) {
	if err := s.handler.UpdatePGHBAConf(req); err != nil {
		return &pb.PGResponse{Succeed: false}, err
	} else {
		return &pb.PGResponse{Succeed: true}, nil
	}
}

func (s *PGService) CreateStandbySignal(ctx context.Context, req *pb.CreateStandbySignalRequest) (*pb.PGResponse, error) {
	if err := s.handler.CreateStandbySignal(req); err != nil {
		return &pb.PGResponse{Succeed: false}, err
	} else {
		return &pb.PGResponse{Succeed: true}, nil
	}
}

func (s *PGService) DeleteStandbySignal(ctx context.Context, req *pb.DeleteStandbySignalRequest) (*pb.PGResponse, error) {
	if err := s.handler.DeleteStandbySignal(req); err != nil {
		return &pb.PGResponse{Succeed: false}, err
	} else {
		return &pb.PGResponse{Succeed: true}, nil
	}
}
