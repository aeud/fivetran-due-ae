package dueio

import "encoding/json"

type DUEHttpResponse struct {
	Data    []*DUEHttpResponseItem   `json:"data"`
	Meta    *DUEHttpResponseMetadata `json:"meta"`
	JSONAPI map[string]interface{}   `json:"jsonapi"`
}

func (r *DUEHttpResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
func (r *DUEHttpResponse) MarshalForce() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

type DUEHttpResponseItem struct {
	ID            string                                     `json:"id"`
	Type          string                                     `json:"type"`
	Attributes    map[string]interface{}                     `json:"attributes"`
	Relationships map[string]DUEHttpResponseItemRelationship `json:"relationships"`
}

type DUEHttpResponseItemRelationship struct {
	Data interface{} `json:"data"`
}

type DUEHttpResponseMetadata struct {
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PageCount   int `json:"page_count"`
	PrevPage    int `json:"prev_page"`
	TotalCount  int `json:"total_count"`
}

func (r *DUEHttpResponseMetadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
func (r *DUEHttpResponseMetadata) MarshalForce() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
