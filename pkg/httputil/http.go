package httputil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DoHTTPRequest(client *http.Client, req *http.Request, out interface{}) error {
	if client == nil {
		return fmt.Errorf("client must not be nil")
	}
	if req == nil {
		return fmt.Errorf("req must not be nil")
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return fmt.Errorf("read respsonse error %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d: %s", resp.StatusCode, data)
	}
	if out != nil {
		if err = json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("unmarshal error: %w - %s", err, data)
		}
	}
	return nil
}
