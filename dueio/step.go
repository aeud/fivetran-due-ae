package dueio

type Step struct {
	Name            string
	UseCursor       bool
	CursorAttribute string
	MinCursorValue  string
	URLGenerator    URLGenerator
	PrimaryKey      string
}

func GetAvailableSteps() map[string]Step {
	return map[string]Step{
		DUECAnswerSetsLabel: {
			Name:            DUECAnswerSetsLabel,
			UseCursor:       true,
			CursorAttribute: "updated_at",
			URLGenerator:    GetAnswerSetsURL,
			PrimaryKey:      "id",
		},
		DUEClientsLabel: {
			Name:         DUEClientsLabel,
			UseCursor:    false,
			URLGenerator: GetClientsURL,
			PrimaryKey:   "id",
		},
		DUECSurveysLabel: {
			Name:         DUECSurveysLabel,
			UseCursor:    false,
			URLGenerator: GetSurveysURL,
			PrimaryKey:   "id",
		},
		DUECFeedbacksLabel: {
			Name:            DUECFeedbacksLabel,
			UseCursor:       true,
			CursorAttribute: "updated-at",
			URLGenerator:    GetFeedbacksURL,
			PrimaryKey:      "id",
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
