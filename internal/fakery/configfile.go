package fakery

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

type headers map[string]string

type request struct {
	Url     string  `json:"url"`
	Headers headers `json:"headers"`
	Method  string  `json:"method"`
}

type response struct {
	Status  int     `json:"status"`
	Headers headers `json:"headers"`
	Latency int     `json:"latency"`
	Body    string  `json:"body"`
}

type FakeryEndpoint struct {
	Request  request
	Response response
}

type fakeryServerConfig struct {
	Endpoints []FakeryEndpoint
}

func CreateNewFakeryServerConfig(file string) (*fakeryServerConfig, error) {
	var config fakeryServerConfig
	var err error = nil

	switch path.Ext(file) {
	case ".json":
		config, err = createFromJSON(file)
	default:
		return &fakeryServerConfig{}, fmt.Errorf("Invalid file format. Must be json.")
	}

	return &config, err
}

func createFromJSON(path string) (fakeryServerConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return fakeryServerConfig{}, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return fakeryServerConfig{}, err
	}

	var config fakeryServerConfig
	err = json.Unmarshal(b, &config.Endpoints)
	if err != nil {
		return fakeryServerConfig{}, err
	}

	return config, nil
}
