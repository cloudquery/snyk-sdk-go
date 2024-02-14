package snyk

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProject_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/long-uuid/projects", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, _ = fmt.Fprint(w, `
{
  "jsonapi": {
		"version": "1.0"
	},
  "data": [
    {
      "id": "e8feca4a-4ebc-494f-80d9-f8b0532188da",
      "type": "project",
      "attributes": {
				"name": "test-org/test-project",
				"origin": "github"
			}
    }
  ]
}
`)
	})
	expectedProjects := []ApiListItem[Project]{
		{
			ID:   "e8feca4a-4ebc-494f-80d9-f8b0532188da",
			Type: "project",
			Attributes: Project{
				Name:   "test-org/test-project",
				Origin: "github",
			},
		},
	}

	actualProjects, _, err := client.Projects.List(ctx, "long-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedProjects, actualProjects)
}

func TestProject_List_emptyOrganizationID(t *testing.T) {
	setup()
	defer teardown()

	_, _, err := client.Projects.List(ctx, "")

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyArgument, err)
}
