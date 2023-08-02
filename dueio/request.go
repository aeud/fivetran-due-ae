package dueio

import (
	"fmt"
	"log"
	"net/url"
)

type APIEndpointURL struct {
	url *url.URL
}

func NewClientsAPIEnpointURL() *APIEndpointURL {
	u, err := url.Parse(fmt.Sprintf("%s/%s/", EndpointV1, DUEClientsLabel))
	if err != nil {
		log.Fatalln(err)
	}
	return &APIEndpointURL{
		url: u,
	}
}
func NewSurveysAPIEnpointURL() *APIEndpointURL {
	u, err := url.Parse(fmt.Sprintf("%s/%s/", EndpointV1, DUESurveysLabel))
	if err != nil {
		log.Fatalln(err)
	}
	return &APIEndpointURL{
		url: u,
	}
}
func NewAnswerSetsAPIEnpointURL() *APIEndpointURL {
	u, err := url.Parse(fmt.Sprintf("%s/%s/", EndpointV1, DUEAnswerSetsLabel))
	if err != nil {
		log.Fatalln(err)
	}
	return &APIEndpointURL{
		url: u,
	}
}

func (e *APIEndpointURL) AddPageSizeParameter(pageSize int) {
	p := e.url.Query()
	p.Add("page[size]", fmt.Sprintf("%d", pageSize))
	e.url.RawQuery = p.Encode()
}
func (e *APIEndpointURL) AddPageNumberParameter(pageNumber int) {
	p := e.url.Query()
	p.Add("page[number]", fmt.Sprintf("%d", pageNumber))
	e.url.RawQuery = p.Encode()
}
func (e *APIEndpointURL) AddFilterParameter(filterName, filterValue string) {
	p := e.url.Query()
	p.Add(fmt.Sprintf("filter[%s]", filterName), filterValue)
	e.url.RawQuery = p.Encode()
}
func (e *APIEndpointURL) AddSortParameter(sortParameter string, asc bool) {
	if !asc {
		sortParameter = fmt.Sprintf("-%s", sortParameter)
	}
	p := e.url.Query()
	p.Add("sort", sortParameter)
	e.url.RawQuery = p.Encode()
}
func (e *APIEndpointURL) String() string {
	return e.url.String()
}
