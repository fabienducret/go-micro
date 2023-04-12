package main

import (
	"broker/ports"
)

func Log(lr ports.Logger, payload ports.LogPayload) (*jsonResponse, error) {
	entry := ports.Log{
		Name: payload.Name,
		Data: payload.Data,
	}

	err := lr.Log(entry)
	if err != nil {
		return nil, err
	}

	return logSentPayload(), nil
}

func logSentPayload() *jsonResponse {
	return &jsonResponse{
		Error:   false,
		Message: "logged",
	}
}
