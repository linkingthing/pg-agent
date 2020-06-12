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

func (s *PGService) CreateRecoveryConf(ctx context.Context, req *pb.CreateRecoveryConfRequest) (*pb.PGResponse, error) {
	if err := s.handler.CreateRecoveryConf(req); err != nil {
		return &pb.PGResponse{Succeed: false}, err
	} else {
		return &pb.PGResponse{Succeed: true}, nil
	}
}

func (s *PGService) DeleteRecoveryConf(ctx context.Context, req *pb.DeleteRecoveryConfRequest) (*pb.PGResponse, error) {
	if err := s.handler.DeleteRecoveryConf(req); err != nil {
		return &pb.PGResponse{Succeed: false}, err
	} else {
		return &pb.PGResponse{Succeed: true}, nil
	}
}
