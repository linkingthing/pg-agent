package grpcservice

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

const PGHBAConfFile = "/data/pg_hba.conf"

var PGHBAConfRegexp = regexp.MustCompile(`^\s*#?\s*(local|host)\s*(replication)`)

func updateHBAConfFile(anotherIP, dbUser string) error {
	pgHBAConfContent, err := getFileContent(PGHBAConfFile, PGHBAConfRegexp)
	if err != nil {
		return fmt.Errorf("get pg hba config content failed: %s", err.Error())
	}

	pgHBAContentSuffix, err := compileTemplate(PG_HBA_CONF, map[string]interface{}{"IP": anotherIP, "User": dbUser})
	if err != nil {
		return fmt.Errorf("generate pg hba config suffix from tempalte failed: %s", err.Error())
	}

	pgHBAConfContent += pgHBAContentSuffix
	if err = ioutil.WriteFile(PGHBAConfFile, []byte(pgHBAConfContent), 0600); err != nil {
		return fmt.Errorf("write pg hba config file failed: %s", err.Error())
	}

	return nil
}
