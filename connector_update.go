package afas

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func (s *ConnectorService) NewUpdateRequest() ConnectorUpdateRequest {
	return ConnectorUpdateRequest{
		api:          s.api,
		method:       http.MethodPut,
		urlParams:    ConnectorUpdateURLParams{},
		queryParams:  ConnectorUpdateQueryParams{},
		requestBody:  s.NewUpdateRequestBody(),
		responseBody: s.NewUpdateResponseBody(),
	}
}

type ConnectorUpdateRequest struct {
	api          *API
	method       string
	urlParams    ConnectorUpdateURLParams
	queryParams  ConnectorUpdateQueryParams
	requestBody  ConnectorUpdateRequestBody
	responseBody *ConnectorUpdateResponseBody
}

func (r *ConnectorUpdateRequest) Method() string {
	return r.method
}

func (r *ConnectorUpdateRequest) SetMethod(method string) {
	r.method = method
}

func (r *ConnectorUpdateRequest) RequestBody() ConnectorUpdateRequestBody {
	return r.requestBody
}

func (r *ConnectorUpdateRequest) SetRequestBody(body ConnectorUpdateRequestBody) {
	r.requestBody = body
}

func (r *ConnectorUpdateRequest) ResponseBody() *ConnectorUpdateResponseBody {
	return r.responseBody
}

func (r *ConnectorUpdateRequest) URL() url.URL {
	path := "/connectors/{connectorid}[/{subelement}]"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	path = strings.Replace(path, "[/{subelement}]", r.urlParams.SubElement, 1)
	return r.api.GetEndpointURL(path)
}

func (r *ConnectorUpdateRequest) Do() (*ConnectorUpdateResponseBody, error) {
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

func (s *ConnectorService) NewUpdateResponseBody() *ConnectorUpdateResponseBody {
	return &ConnectorUpdateResponseBody{}
}

func (r *ConnectorUpdateRequest) QueryParams() *ConnectorUpdateQueryParams {
	return &r.queryParams
}

type ConnectorUpdateQueryParams struct {
}

func (r *ConnectorUpdateRequest) URLParams() *ConnectorUpdateURLParams {
	return &r.urlParams
}

type ConnectorUpdateURLParams struct {
	ConnectorID string
	SubElement  string
}

func (s *ConnectorService) NewUpdateRequestBody() ConnectorUpdateRequestBody {
	return struct{}{}
}

type ConnectorUpdateRequestBody interface{}

// {"results":{"FiEntryPar":{"UnId":"4","EnNo":"54605"}}}
type ConnectorUpdateResponseBody struct {
	Results json.RawMessage `json:"results"`
}
