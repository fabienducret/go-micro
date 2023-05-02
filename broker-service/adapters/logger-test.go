package adapters

import (
	"broker/ports"
	"fmt"
)

type testLogger struct{}

func NewTestLogger() *testLogger {
	return &testLogger{}
}

func (l *testLogger) Log(toLog ports.Log) (string, error) {
	return fmt.Sprintf("Log handled for:%s", toLog.Name), nil
}
