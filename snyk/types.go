package snyk

type ApiResponse[T any] struct {
	JsonApi struct {
		Version string `json:"version,omitempty"`
	} `json:"jsonapi,omitempty"`
	Links struct {
		Self  string `json:"self,omitempty"`
		First string `json:"first,omitempty"`
		Last  string `json:"last,omitempty"`
		Prev  string `json:"prev,omitempty"`
		Next  string `json:"next,omitempty"`
	} `json:"links,omitempty"`
	Data []ApiListItem[T] `json:"data,omitempty"`
}

type ApiListItem[T any] struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Attributes T      `json:"attributes,omitempty"`
}
