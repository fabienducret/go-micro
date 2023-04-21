package main

import (
	"broker/ports"
)

func Log(lr ports.Logger, payload ports.LogPayload) (*jsonResponse, error) {
	entry := ports.Log(payload)

	result, err := lr.Log(entry)
	if err != nil {
		return nil, err
	}

	return logSentPayload(result), nil
}

func logSentPayload(result string) *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: result,
	}
}
