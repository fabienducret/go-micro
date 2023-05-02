package adapters

import "authentication/ports"

type testLogger struct{}

func NewTestLogger() *testLogger {
	return &testLogger{}
}

func (r *testLogger) Log(toLog ports.Log) error {
	return nil
}
