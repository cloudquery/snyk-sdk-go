package snyk

import (
	"context"
	"net/http"
)

const orgBasePath = "orgs"

// OrgsService handles communication with the organization related methods of the Snyk API.
type OrgsService service

// Organization represents a Snyk organization.
type Organization struct {
	Name       string `json:"name,omitempty"`
	Slug       string `json:"slug,omitempty"`
	IsPersonal bool   `json:"is_personal,omitempty"`
}

// List provides a list of all organizations a user belongs to.
func (s *OrgsService) List(ctx context.Context) ([]ApiListItem[Organization], *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, orgBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ApiResponse[Organization])
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Data, resp, nil
}
