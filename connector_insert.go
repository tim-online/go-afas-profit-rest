package afas

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func (s *ConnectorService) NewInsertRequest() ConnectorInsertRequest {
	return ConnectorInsertRequest{
		api:          s.api,
		method:       http.MethodPost,
		urlParams:    ConnectorInsertURLParams{},
		queryParams:  ConnectorInsertQueryParams{},
		requestBody:  s.NewInsertRequestBody(),
		responseBody: s.NewInsertResponseBody(),
	}
}

type ConnectorInsertRequest struct {
	api          *API
	method       string
	urlParams    ConnectorInsertURLParams
	queryParams  ConnectorInsertQueryParams
	requestBody  ConnectorInsertRequestBody
	responseBody *ConnectorInsertResponseBody
}

func (r *ConnectorInsertRequest) Method() string {
	return r.method
}

func (r *ConnectorInsertRequest) SetMethod(method string) {
	r.method = method
}

func (r *ConnectorInsertRequest) RequestBody() ConnectorInsertRequestBody {
	return r.requestBody
}

func (r *ConnectorInsertRequest) SetRequestBody(body ConnectorInsertRequestBody) {
	r.requestBody = body
}

func (r *ConnectorInsertRequest) ResponseBody() *ConnectorInsertResponseBody {
	return r.responseBody
}

func (r *ConnectorInsertRequest) URL() url.URL {
	path := "/connectors/{connectorid}[/{subelement}]"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	path = strings.Replace(path, "[/{subelement}]", r.urlParams.SubElement, 1)
	return r.api.GetEndpointURL(path)
}

func (r *ConnectorInsertRequest) Do() (*ConnectorInsertResponseBody, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return r.ResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return r.ResponseBody(), err
	}

	responseBody := r.ResponseBody()
	_, err = r.api.Do(req, responseBody)
	return responseBody, err
}

func (s *ConnectorService) NewInsertResponseBody() *ConnectorInsertResponseBody {
	return &ConnectorInsertResponseBody{}
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

// {"results":{"FiEntryPar":{"UnId":"4","EnNo":"54605"}}}
type ConnectorInsertResponseBody struct {
	Results json.RawMessage `json:"results"`
}
