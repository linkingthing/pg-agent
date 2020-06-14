package grpcservice

import (
	pb "github.com/linkingthing/pg-agent/pkg/proto"
)

type PGHandler struct{}

func newPGHandler() *PGHandler {
	return &PGHandler{}
}

func (h *PGHandler) UpdatePostgresqlConf(req *pb.UpdatePostgresqlConfRequest) error {
	return updatePGConfFile(req)
}

func (h *PGHandler) UpdatePGHBAConf(req *pb.UpdatePGHBAConfRequest) error {
	return updateHBAConfFile(req)
}

func (h *PGHandler) CreateStandbySignal(req *pb.CreateStandbySignalRequest) error {
	return createStandbySignalFile()
}

func (h *PGHandler) DeleteStandbySignal(req *pb.DeleteStandbySignalRequest) error {
	return deleteStandbySignalFile()
}
