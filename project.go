package identity_platform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Project struct {
	c *http.Client
}

func NewProject(client *http.Client) *Project {
	return &Project{
		c: client,
	}
}

func (project *Project) GetConfig(projectId string) (*Config, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/admin/v2/projects/%s/config", projectId)

	resp, err := project.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, formatError(resp.Body)
	}

	return hydrate(resp.Body)
}

func (project *Project) UpdateConfig(projectId string, newConfig *Config) (*Config, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/admin/v2/projects/%s/config", projectId)

	payload, err := json.Marshal(newConfig)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := project.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, formatError(resp.Body)
	}

	return hydrate(resp.Body)
}

func hydrate(r io.Reader) (*Config, error) {
	var config Config
	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

func formatError(r io.Reader) error {
	var apiErrorResponse ApiErrorResponse
	err := json.NewDecoder(r).Decode(&apiErrorResponse)
	if err != nil {
		return err
	}

	return errors.New(apiErrorResponse.Error.Message)
}
