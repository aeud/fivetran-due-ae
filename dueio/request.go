package dueio

import "fmt"

type URLGenerator func(pageNumber int) string

func GetAnswerSetsURL(pageNumber int) string {
	return fmt.Sprintf("%s/%s/?page[number]=%d&sort=-updated_at", EndpointV1, DUECAnswerSetsLabel, pageNumber)
}
func GetClientsURL(pageNumber int) string {
	return fmt.Sprintf("%s/%s/?page[number]=%d", EndpointV1, DUEClientsLabel, pageNumber)
}
func GetSurveysURL(pageNumber int) string {
	return fmt.Sprintf("%s/%s/?page[number]=%d", EndpointV1, DUECSurveysLabel, pageNumber)
}
func GetFeedbacksURL(pageNumber int) string {
	return fmt.Sprintf("%s/%s/?page[number]=%d&sort=-updated-at", EndpointV3, DUECFeedbacksLabel, pageNumber)
}

func AddPagSizeParameter(url string, pageSize int) string {
	return fmt.Sprintf("%s&page[size]=%d", url, pageSize)
}
