package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service- registry service responded with a code %v", res.StatusCode)
	}
	return nil
}

func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(http.MethodDelete,
		ServicesURL,
		bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	)
}
