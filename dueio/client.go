package dueio

import (
	"due/fivetranio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DUEAPIClient struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewDUEAPIClient(key string) *DUEAPIClient {
	httpClient := new(http.Client)
	return &DUEAPIClient{
		APIKey:     key,
		HTTPClient: httpClient,
	}
}
func (client *DUEAPIClient) NewGetRequest(url string) (*http.Request, error) {
	log.Printf("Calling: %s\n", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, "")
	return req, nil
}
func (client *DUEAPIClient) DoRequest(req *http.Request) (*DUEHttpResponse, error) {
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		bs, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("api response error %d (%s)", resp.StatusCode, bs)
	}
	defer resp.Body.Close()
	var v DUEHttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (client *DUEAPIClient) Collect(endpointURL *APIEndpointURL) (*DUEHttpResponse, error) {
	req, err := client.NewGetRequest(endpointURL.String())
	if err != nil {
		return nil, err
	}
	return client.DoRequest(req)
}

func (client *DUEAPIClient) ExecuteState(state *fivetranio.State) (data []map[string]interface{}, nextState *fivetranio.State, hasMore bool, err error) {
	availableSteps := GetAvailableSteps()
	currentStepName := state.CurrentStep
	pageNumber := state.NextPageNumber

	step, exists := availableSteps[currentStepName]
	if !exists {
		log.Printf("missing step `%s`, skipping\n", currentStepName)
		data = make([]map[string]interface{}, 0)
		nextState, hasMore, _ = state.NextStep()
		return
	}

	u := step.APIEnpointURL
	u.AddPageNumberParameter(pageNumber)

	pageSize := step.PageSize
	if pageSize == 0 {
		pageSize = DefaultPageSize
	}
	u.AddPageSizeParameter(pageSize)

	if step.UseCursor {
		if v := state.GetCursorNextValue(step.Name); v != "" {
			u.AddFilterParameter(step.FilterAttribute, v)
		} else if step.MinCursorValue != "" {
			u.AddFilterParameter(step.FilterAttribute, step.MinCursorValue)
		}
		u.AddSortParameter(step.CursorAttribute, true) // Last element to fetch will be the latest updated
	}

	resp, err := client.Collect(u)
	if err != nil {
		return
	}

	// Build the `next state`
	if p := resp.Meta.NextPage; p > 0 {
		nextState, hasMore, _ = state.NextPage(p)
	} else {
		if step.UseCursor && len(resp.Data) > 0 {
			lastElement := resp.Data[len(resp.Data)-1]
			lastCursorVal := lastElement.Attributes[step.CursorAttribute]
			state.SetCursorNextValue(step.Name, fmt.Sprintf("%v", lastCursorVal))
		}
		nextState, hasMore, _ = state.NextStep()
	}

	// Build the `data` payload
	data = make([]map[string]interface{}, len(resp.Data))
	for i, row := range resp.Data {
		data[i] = row.Attributes
		data[i]["id"] = row.ID
		data[i]["relationships"] = row.Relationships
	}

	return
}
