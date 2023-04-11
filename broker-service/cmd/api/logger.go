package main

import (
	"broker/ports"
)

func Log(lr ports.Logger, name, data string) (*jsonResponse, error) {
	entry := ports.Log{
		Name: name,
		Data: data,
	}

	err := lr.Log(entry)
	if err != nil {
		return nil, err
	}

	payload := &jsonResponse{
		Error:   false,
		Message: "logged",
	}

	return payload, nil
}
