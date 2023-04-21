package main

import (
	"broker/ports"
)

func Log(lr ports.Logger, payload ports.LogPayload) (*jsonResponse, error) {
	entry := ports.Log(payload)

	reply, err := lr.Log(entry)
	if err != nil {
		return nil, err
	}

	return logSentPayload(reply), nil
}

func logSentPayload(reply string) *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: reply,
	}
}
