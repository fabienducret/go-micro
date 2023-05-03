package project

import (
	"net/http"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	payload := "{\"action\":\"auth\",\"auth\":{\"email\":\"admin@example.com\",\"password\":\"verysecret\"}}"

	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/handle", strings.NewReader(payload))

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Invalid status code %v", resp.StatusCode)
	}
}
