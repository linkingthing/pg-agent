package grpcservice

import (
	"fmt"
	"os"
	"os/exec"

	pb "github.com/linkingthing/pg-agent/pkg/proto"
)

const RsyncCmd = "rsync -az --password-file=/etc/rsyncd.passwd --delete /data/ rsync@%s::postgres"

func rsyncPostgresqlData(req *pb.RsyncPostgresqlDataRequest) error {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(RsyncCmd, req.GetAddress()))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
