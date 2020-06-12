package grpcservice

import (
	pb "github.com/linkingthing/pg-agent/pkg/proto"
)

type PGHandler struct{}

func newPGHandler() *PGHandler {
	return &PGHandler{}
}

func (h *PGHandler) UpdatePostgresqlConf(req *pb.UpdatePostgresqlConfRequest) error {
	return updatePGConfFile(req.GetPort(), req.GetIsMaster())
}

func (h *PGHandler) UpdatePGHBAConf(req *pb.UpdatePGHBAConfRequest) error {
	return updateHBAConfFile(req.GetAnotherIp(), req.GetUser())
}

func (h *PGHandler) CreateRecoveryConf(req *pb.CreateRecoveryConfRequest) error {
	return createRecoveryConfFile(req.GetAnotherIp(), req.GetPort())
}

func (h *PGHandler) DeleteRecoveryConf(req *pb.DeleteRecoveryConfRequest) error {
	return deleteRecoveryConfFile()
}
