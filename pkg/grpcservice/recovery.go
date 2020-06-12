package grpcservice

import (
	"fmt"
	"os"
)

const PGRecoverFile = "/data/recovery.conf"

func createRecoveryConfFile(host string, dbPort uint32) error {
	f, err := os.Create(PGRecoverFile)
	if err != nil {
		return fmt.Errorf("create recovery.conf failed: %s", err.Error())
	}

	pgRecoverFilePathContent, err := compileTemplate(PG_RECOVERY_CONF, map[string]interface{}{"Host": host, "Port": dbPort})
	if err != nil {
		return fmt.Errorf("generate recovery config from tempalte failed: %s", err.Error())
	}

	if _, err := f.WriteString(pgRecoverFilePathContent); err != nil {
		return fmt.Errorf("write recovery config to file failed: %s", err.Error())
	}

	return nil
}

func deleteRecoveryConfFile() error {
	if _, err := os.Stat(PGRecoverFile); err != nil {
		return err
	}

	return os.Remove(PGRecoverFile)
}
