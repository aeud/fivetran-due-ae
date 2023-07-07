package fivetranio

import (
	"encoding/json"
	"io"
	"net/http"
)

type FivetranCloudFunctionState map[string]interface{}

type FivetranCloudFunctionRequest struct {
	Agent     string  `json:"agent"`
	State     *State  `json:"state"`
	Secrets   *Secret `json:"secrets"`
	SetupTest bool    `json:"setup_test,omitempty"`
}

func NewFivetranCloudFunctionRequestFromReader(r io.Reader) (*FivetranCloudFunctionRequest, error) {
	var v FivetranCloudFunctionRequest
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

type FivetranCloudFunctionResponse struct {
	State        *State                              `json:"state"`
	Schema       map[string]map[string][]string      `json:"schema"`
	Insert       map[string][]map[string]interface{} `json:"insert"`
	Delete       map[string][]interface{}            `json:"delete"`
	HasMore      bool                                `json:"hasMore"`
	ErrorMessage string                              `json:"errorMessage,omitempty"`
}

func NewFivetranCloudFunctionResponse() *FivetranCloudFunctionResponse {
	return &FivetranCloudFunctionResponse{
		Schema:  make(map[string]map[string][]string),
		Insert:  make(map[string][]map[string]interface{}),
		Delete:  make(map[string][]interface{}),
		HasMore: false,
	}
}
func NewFivetranCloudFunctionErrorResponse(errorMessage string) *FivetranCloudFunctionResponse {
	return &FivetranCloudFunctionResponse{
		ErrorMessage: errorMessage,
	}
}

func (r *FivetranCloudFunctionResponse) AddPrimaryKey(entity, attribute string) {
	r.Schema[entity] = map[string][]string{
		"primary_key": {attribute},
	}
}

func (r *FivetranCloudFunctionResponse) SetNextState(state *State) {
	r.State = state
}

func (r *FivetranCloudFunctionResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
func (r *FivetranCloudFunctionResponse) MarshalForce() []byte {
	bs, _ := r.Marshal()
	return bs
}

func (r *FivetranCloudFunctionResponse) Send(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	if r.ErrorMessage != "" {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err := w.Write(r.MarshalForce())
	if err != nil {
		return err
	}
	return nil
}
