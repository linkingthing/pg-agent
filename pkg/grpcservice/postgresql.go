package grpcservice

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	pb "github.com/linkingthing/pg-agent/pkg/proto"
)

const PGConfFile = "/data/postgresql.conf"

var PGConfRegexp = regexp.MustCompile(`^\s*#?\s*(listen_addresses|port|wal_level|archive_mode|archive_command|max_wal_senders|wal_keep_segments|hot_standby|primary_conninfo|synchronous_standby_names)\s*=\s*`)

func updatePGConfFile(req *pb.UpdatePostgresqlConfRequest) error {
	pgConfContent, err := getFileContent(PGConfFile, PGConfRegexp)
	if err != nil {
		return fmt.Errorf("get pg config file content failed: %s", err.Error())
	}

	pgConfContentSuffix, err := compileTemplate(PG_CONF, map[string]interface{}{
		"Host":     req.GetHost(),
		"User":     req.GetUser(),
		"Password": req.GetPassword(),
		"Port":     req.GetPort(),
	})
	if err != nil {
		return fmt.Errorf("generate pg config suffix from tempalte failed: %s", err.Error())
	}

	pgConfContent += pgConfContentSuffix
	if req.GetIsMaster() {
		pgConfContent += PG_EXTENSIONS_FOR_MASTER_CONF
	}

	if err := ioutil.WriteFile(PGConfFile, []byte(pgConfContent), 0600); err != nil {
		return fmt.Errorf("write pg config file failed: %s", err.Error())
	}

	return nil
}

func getFileContent(file string, reg *regexp.Regexp) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	content := ""
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "#~lx") {
			break
		}

		if len(reg.FindStringSubmatch(txt)) == 0 {
			content += fmt.Sprintf("%s\n", txt)
		}
	}

	return content, nil
}
