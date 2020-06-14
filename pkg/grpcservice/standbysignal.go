package grpcservice

import (
	"fmt"
	"os"
)

const PGStandbySignalFile = "/data/standby.signal"

func createStandbySignalFile() error {
	f, err := os.Create(PGStandbySignalFile)
	if err != nil {
		return fmt.Errorf("create standby.signal failed: %s", err.Error())
	}

	if _, err := f.WriteString(PG_STANDBY_SIGNAL); err != nil {
		return fmt.Errorf("write standby.signal config to file failed: %s", err.Error())
	}

	return nil
}

func deleteStandbySignalFile() error {
	if _, err := os.Stat(PGStandbySignalFile); err == nil {
		return os.Remove(PGStandbySignalFile)
	}

	return nil
}
