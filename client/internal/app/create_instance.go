package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type createInstancePostRequestResponse struct {
	IpAddress string `json:"ip_address"`
}

const (
	DEFAULT_HOST = "192.168.0.9:1323"
)

func runCreateInstanceCmd(_ context.Context) error {
	url := fmt.Sprintf("http://%s/create-instance", DEFAULT_HOST) // TODO: https化

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(make([]byte, 0)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var res createInstancePostRequestResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	fmt.Printf("created virtual machine. IP address: %s\n", res.IpAddress)

	return nil
}
