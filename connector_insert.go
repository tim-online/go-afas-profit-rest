package afas

import (
	"net/http"
	"net/url"
	"strings"
)

func (s *ConnectorService) NewInsertRequest() ConnectorInsertRequest {
	return ConnectorInsertRequest{
		api:         s.api,
		method:      http.MethodPost,
		urlParams:   ConnectorInsertURLParams{},
		queryParams: ConnectorInsertQueryParams{},
		requestBody: s.NewInsertRequestBody(),
	}
}

type ConnectorInsertRequest struct {
	api         *API
	method      string
	urlParams   ConnectorInsertURLParams
	queryParams ConnectorInsertQueryParams
	requestBody ConnectorInsertRequestBody
}

func (r *ConnectorInsertRequest) Method() string {
	return r.method
}

func (r *ConnectorInsertRequest) SetMethod(method string) {
	r.method = method
}

func (r *ConnectorInsertRequest) RequestBody() *ConnectorInsertRequestBody {
	return &r.requestBody
}

func (r *ConnectorInsertRequest) SetRequestBody(body ConnectorInsertRequestBody) {
	r.requestBody = body
}

func (r *ConnectorInsertRequest) URL() url.URL {
	path := "/connectors/{connectorid}[/{subelement}]"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	path = strings.Replace(path, "[/{subelement}]", r.urlParams.SubElement, 1)
	return r.api.GetEndpointURL(path)
}

func (r *ConnectorInsertRequest) Do() (ConnectorInsertResponseBody, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return responseBody, err
}

func (r *ConnectorInsertRequest) NewResponseBody() ConnectorInsertResponseBody {
	return struct{}{}
}

func (r *ConnectorInsertRequest) QueryParams() *ConnectorInsertQueryParams {
	return &r.queryParams
}

type ConnectorInsertQueryParams struct {
}

func (r *ConnectorInsertRequest) URLParams() *ConnectorInsertURLParams {
	return &r.urlParams
}

type ConnectorInsertURLParams struct {
	ConnectorID string
	SubElement  string
}

func (s *ConnectorService) NewInsertRequestBody() ConnectorInsertRequestBody {
	return struct{}{}
}

type ConnectorInsertRequestBody interface{}

type ConnectorInsertResponseBody interface{}
