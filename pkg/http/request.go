package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func Request(endpoint string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.render.com/v1/%s", endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request for endpoint %s: %w", endpoint, err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken()))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting endpoint %s: %w", endpoint, err)
	}

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading request for endpoint %s: %w", endpoint, err)
	}

	return jsonString, nil
}

func Create(endpoint string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.render.com/v1/%s", endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating request for endpoint %s: %w", endpoint, err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken()))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting endpoint %s: %w", endpoint, err)
	}

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading request for endpoint %s: %w", endpoint, err)
	}

	return jsonString, nil
}

func apiToken() string {
	api := viper.GetString("api")
	if api == "" {
		panic("you must set a token before you can make requests")
	}
	return api
}