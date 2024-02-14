package snyk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestOrgs_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, _ = fmt.Fprint(w, `
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": [
    {
      "id": "long-uuid-first",
      "type": "org",
      "attributes": {
        "is_personal": true,
        "name": "Test Org First",
        "slug": "test-org-first"
      }
    },
    {
      "id": "long-uuid-second",
      "type": "org",
      "attributes": {
        "is_personal": false,
        "name": "Test Org Second",
        "slug": "test-org-second"
      }
    }
  ]
}`)
	})
	expectedOrgs := []ApiListItem[Organization]{
		{
			ID:   "long-uuid-first",
			Type: "org",
			Attributes: Organization{
				IsPersonal: true,
				Name:       "Test Org First",
				Slug:       "test-org-first",
			},
		},
		{
			ID:   "long-uuid-second",
			Type: "org",
			Attributes: Organization{
				IsPersonal: false,
				Name:       "Test Org Second",
				Slug:       "test-org-second",
			},
		},
	}

	actualOrgs, _, err := client.Orgs.List(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrgs, actualOrgs)
}
