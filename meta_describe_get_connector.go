package afas

import (
	"net/http"
	"net/url"
	"strings"
)

func (s *MetaService) NewDescribeGetConnectorRequest() MetaDescribeGetConnectorRequest {
	return MetaDescribeGetConnectorRequest{
		api:         s.api,
		method:      http.MethodGet,
		urlParams:   MetaDescribeGetConnectorURLParams{},
		queryParams: MetaDescribeGetConnectorQueryParams{},
	}
}

type MetaDescribeGetConnectorRequest struct {
	api         *API
	method      string
	urlParams   MetaDescribeGetConnectorURLParams
	queryParams MetaDescribeGetConnectorQueryParams
}

func (r *MetaDescribeGetConnectorRequest) Method() string {
	return r.method
}

func (r *MetaDescribeGetConnectorRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetaDescribeGetConnectorRequest) RequestBody() *EmptyRequestBody {
	return &EmptyRequestBody{}
}

func (r *MetaDescribeGetConnectorRequest) SetRequestBody(body EmptyRequestBody) {
}

func (r *MetaDescribeGetConnectorRequest) URL() url.URL {
	path := "metainfo/get/{connectorid}"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	return r.api.GetEndpointURL(path)
}

func (r *MetaDescribeGetConnectorRequest) Do() (MetaDescribeGetConnectorResponseBody, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *MetaDescribeGetConnectorRequest) NewResponseBody() *MetaDescribeGetConnectorResponseBody {
	return &MetaDescribeGetConnectorResponseBody{}
}

func (r *MetaDescribeGetConnectorRequest) QueryParams() *MetaDescribeGetConnectorQueryParams {
	return &r.queryParams
}

type MetaDescribeGetConnectorQueryParams struct {
}

func (r *MetaDescribeGetConnectorRequest) URLParams() *MetaDescribeGetConnectorURLParams {
	return &r.urlParams
}

type MetaDescribeGetConnectorURLParams struct {
	ConnectorID string
}

type MetaDescribeGetConnectorResponseBody struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Fields      []GetConnectorField `json:"fields"`
}

type GetConnectorField struct {
	ID              string `json:"id"`
	FieldID         string `json:"fieldId"`
	DataType        string `json:"dataType"`
	Label           string `json:"label"`
	Length          int    `json:"length"`
	ControlType     int    `json:"controlType"`
	Decimals        int    `json:"decimals"`
	DecimalsFieldID string `json:"decimalsFieldId"`
}
