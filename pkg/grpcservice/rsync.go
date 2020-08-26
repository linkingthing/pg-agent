package grpcservice

import (
	"fmt"

	pb "github.com/linkingthing/pg-agent/pkg/proto"
	"github.com/linkingthing/pg-agent/pkg/util"
)

const RsyncCmd = "rsync -az --password-file=/etc/rsyncd.passwd --delete /data/ rsync@%s::postgres"

func rsyncPostgresqlData(req *pb.RsyncPostgresqlDataRequest) error {
	return util.ExecCommand(fmt.Sprintf(RsyncCmd, req.GetAddress()))
}
