package dueio

import "time"

type Step struct {
	Name            string
	UseCursor       bool
	CursorAttribute string
	MinCursorValue  string
	URLGenerator    URLGenerator
	PrimaryKey      string
	PageSize        int
}

func GetAvailableSteps() map[string]Step {
	return map[string]Step{
		DUECAnswerSetsLabel: {
			Name:            DUECAnswerSetsLabel,
			UseCursor:       true,
			CursorAttribute: "updated_at",
			URLGenerator:    GetAnswerSetsURL,
			PrimaryKey:      "id",
			MinCursorValue:  time.Now().Add(-2 * 365 * 24 * time.Hour).Format("2006-01"),
			PageSize:        100,
		},
		DUEClientsLabel: {
			Name:         DUEClientsLabel,
			UseCursor:    false,
			URLGenerator: GetClientsURL,
			PrimaryKey:   "id",
			PageSize:     10000,
		},
		DUECSurveysLabel: {
			Name:         DUECSurveysLabel,
			UseCursor:    false,
			URLGenerator: GetSurveysURL,
			PrimaryKey:   "id",
			PageSize:     100,
		},
		// DUECFeedbacksLabel: {
		// 	Name:            DUECFeedbacksLabel,
		// 	UseCursor:       true,
		// 	CursorAttribute: "updated-at",
		// 	URLGenerator:    GetFeedbacksURL,
		// 	PrimaryKey:      "id",
		// 	MinCursorValue:  time.Now().Add(-2 * 365 * 24 * time.Hour).Format("2006-01"),
		// 	PageSize:        100,
		// },
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
