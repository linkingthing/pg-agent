package grpcservice

import (
	"fmt"
	"os"
)

const PGStandbySignalFile = "/data/standby.signal"

func createStandbySignalFile() error {
	_, err := os.Create(PGStandbySignalFile)
	if err != nil {
		return fmt.Errorf("create standby.signal failed: %s", err.Error())
	}

	return nil
}

func deleteStandbySignalFile() error {
	if _, err := os.Stat(PGStandbySignalFile); err == nil {
		return os.Remove(PGStandbySignalFile)
	}

	return nil
}
