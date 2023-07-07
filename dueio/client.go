package dueio

import (
	"due/fivetranio"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		bs, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("api response error %d (%s)", resp.StatusCode, bs)
	}
	defer resp.Body.Close()
	var v DUEHttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (client *DUEAPIClient) Collect(url string) (*DUEHttpResponse, error) {
	req, err := client.NewGetRequest(url)
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
		err = fmt.Errorf("no matching step")
		return
	}

	url := step.URLGenerator(pageNumber)

	pageSize := state.PageSize
	if pageSize == 0 {
		pageSize = DefaultPageSize
	}
	url = AddPagSizeParameter(url, pageSize)

	resp, err := client.Collect(url)
	if err != nil {
		return
	}

	// Build the `next state`

	cursorTest := false // assumes there is no cursor
	if step.UseCursor {
		if pageNumber == 1 {
			firstElement := resp.Data[0]
			targetCursor := firstElement.Attributes[step.CursorAttribute]
			state.SetTargetCursor(step.Name, fmt.Sprintf("%v", targetCursor))
		}
		lastElement := resp.Data[len(resp.Data)-1]
		lastCursorVal := lastElement.Attributes[step.CursorAttribute]
		if state.IsLowerThanCurrentCursor(step.Name, lastCursorVal) || fmt.Sprintf("%v", lastCursorVal) < step.MinCursorValue {
			cursorTest = true
		}
	}

	if p := resp.Meta.NextPage; p > 0 && !cursorTest {
		nextState, hasMore, _ = state.NextPage(p)
	} else {
		state.CloseCursor(step.Name)
		nextState, hasMore, _ = state.NextStep()
	}

	// Build the `data`` payload
	data = make([]map[string]interface{}, len(resp.Data))
	for i, row := range resp.Data {
		data[i] = row.Attributes
		data[i]["id"] = row.ID
		data[i]["relationships"] = row.Relationships
	}

	return
}
