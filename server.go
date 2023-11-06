package main

import (
	"due/dueio"
	"due/fivetranio"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fivetranReq, err := fivetranio.NewFivetranCloudFunctionRequestFromReader(r.Body)
	if err != nil {
		fivetranio.NewFivetranCloudFunctionErrorResponse(err.Error()).Send(w)
		return
	}

	// Send a canonical response for Fivetran's tests
	if fivetranReq.SetupTest {
		fivetranio.NewFivetranCloudFunctionResponse().Send(w)
		return
	}

	// Parse the secrets
	secrets := fivetranReq.Secrets
	if secrets == nil {
		fivetranio.NewFivetranCloudFunctionErrorResponse("missing secrets").Send(w)
		return
	}
	workEntities := secrets.Entities
	if len(workEntities) == 0 {
		workEntities = dueio.GetAvailableEntities()
	}
	state := fivetranReq.State
	if state == nil {
		fivetranio.NewFivetranCloudFunctionErrorResponse("missing state").Send(w)
		return
	}
	// Initial state
	if state.Version != fivetranio.CurrentStateVersion || state.CurrentStep == "" {
		if err := state.Reset(workEntities); err != nil {
			fivetranio.NewFivetranCloudFunctionErrorResponse(err.Error()).Send(w)
			return
		}
	}

	DUEClient := dueio.NewDUEAPIClient(fivetranReq.Secrets.APIKey)
	fivetranResp := fivetranio.NewFivetranCloudFunctionResponse()

	availableSteps := dueio.GetAvailableSteps()
	for _, e := range workEntities {
		fivetranResp.AddPrimaryKey(e, availableSteps[e].PrimaryKey)
	}

	if r.URL.Query().Has("debug") || os.Getenv("DEBUG") == "Y" {
		state.Debug = true
	}

	data, nextState, hasMore, err := DUEClient.ExecuteState(state)
	if err != nil {
		fivetranio.NewFivetranCloudFunctionErrorResponse(err.Error()).Send(w)
		return
	}
	fivetranResp.HasMore = hasMore
	fivetranResp.State = nextState
	fivetranResp.Insert[state.CurrentStep] = data
	fivetranResp.Send(w)
}
