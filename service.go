package identity_platform

import (
	"context"
	"golang.org/x/oauth2/google"
)

type ServiceScope string

type Service struct {
	Project *Project
}

// New creates a new instance of the Identity Platform Service.
// Valid scopes can be found at https://cloud.google.com/identity-platform/docs/reference/rest/v2/projects/updateConfig
func New(ctx context.Context, scope string) (*Service, error) {
	httpClient, err := google.DefaultClient(ctx, scope)
	if err != nil {
		return nil, err
	}

	return &Service{
		Project: NewProject(httpClient),
	}, nil
}
