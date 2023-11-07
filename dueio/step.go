package dueio

import "time"

type Step struct {
	Name                string
	UseCursor           bool
	CursorAttribute     string
	FilterAttribute     string
	InitFilterAttribute string
	SortAttribute       string
	MinCursorValue      string
	APIEnpointURL       *APIEndpointURL
	PrimaryKey          string
	PageSize            int
}

func GetAvailableSteps() map[string]Step {
	return map[string]Step{
		DUEAnswerSetsLabel: {
			Name:                DUEAnswerSetsLabel,
			UseCursor:           true,
			CursorAttribute:     "updated_at",
			FilterAttribute:     "start_updated_date",
			InitFilterAttribute: "start_updated_date",
			SortAttribute:       "updated_at",
			APIEnpointURL:       NewAnswerSetsAPIEnpointURL(),
			PrimaryKey:          "id",
			// MinCursorValue:      time.Now().Add(-2 * 365 * 24 * time.Hour).Truncate(time.Hour).Format("2006-01-02T00:00:00.000Z"), // 2 years from now
			MinCursorValue: time.Now().Add(-2 * 93 * 24 * time.Hour).Truncate(time.Hour).Format("2006-01-02T00:00:00.000Z"), // 93 days from now
			PageSize:       100,
		},
		// DUEClientsLabel: {
		// 	Name:          DUEClientsLabel,
		// 	UseCursor:     false,
		// 	APIEnpointURL: NewClientsAPIEnpointURL(),
		// 	PrimaryKey:    "id",
		// 	PageSize:      10000,
		// },
		DUESurveysLabel: {
			Name:          DUESurveysLabel,
			UseCursor:     false,
			APIEnpointURL: NewSurveysAPIEnpointURL(),
			PrimaryKey:    "id",
			PageSize:      100,
		},
	}
}

func GetAvailableEntities() []string {
	allSteps := GetAvailableSteps()
	availableEntities := make([]string, len(allSteps))
	i := 0
	for e := range allSteps {
		availableEntities[i] = e
		i++
	}
	return availableEntities
}
