package afas

import (
	"net/http"
	"net/url"
	"strings"
)

func (s *ConnectorService) NewListRequest() ConnectorListRequest {
	return ConnectorListRequest{
		api:       s.api,
		method:    http.MethodGet,
		urlParams: ConnectorListURLParams{},
		queryParams: ConnectorListQueryParams{
			Skip:            0,
			Take:            0,
			OrderByFieldIDs: "",
			FilterFieldIDs:  "",
			FilterValues:    "",
			OperatorTypes:   "",
		},
		requestBody:  s.NewListRequestBody(),
		responseBody: s.NewListResponseBody(),
	}
}

type ConnectorListRequest struct {
	api          *API
	method       string
	urlParams    ConnectorListURLParams
	queryParams  ConnectorListQueryParams
	requestBody  EmptyRequestBody
	responseBody *ConnectorListResponseBody
}

func (r *ConnectorListRequest) Method() string {
	return r.method
}

func (r *ConnectorListRequest) SetMethod(method string) {
	r.method = method
}

func (r *ConnectorListRequest) RequestBody() *EmptyRequestBody {
	return &r.requestBody
}

func (r *ConnectorListRequest) SetRequestBody(body EmptyRequestBody) {
	r.requestBody = body
}

func (r *ConnectorListRequest) ResponseBody() *ConnectorListResponseBody {
	return r.responseBody
}

func (r *ConnectorListRequest) SetResponseBody(body *ConnectorListResponseBody) {
	r.responseBody = body
}

func (r *ConnectorListRequest) URL() url.URL {
	path := "/connectors/{connectorid}"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	return r.api.GetEndpointURL(path)
}

func (r *ConnectorListRequest) Do() (*http.Response, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return nil, err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return nil, err
	}

	resp, err := r.api.Do(req, r.responseBody)
	return resp, err
}

func (s *ConnectorService) NewListResponseBody() *ConnectorListResponseBody {
	return &ConnectorListResponseBody{}
}

func (r *ConnectorListRequest) QueryParams() *ConnectorListQueryParams {
	return &r.queryParams
}

type ConnectorListQueryParams struct {
	Skip            int    `schema:"skip,omitempty"`
	Take            int    `schema:"take,omitempty"`
	OrderByFieldIDs string `schema:"orderbyfieldids,omitempty"`
	FilterFieldIDs  string `schema:"filterfieldids,omitempty"`
	FilterValues    string `schema:"filtervalues,omitempty"`
	OperatorTypes   string `schema:"operatortypes,omitempty"`
}

func (r *ConnectorListRequest) URLParams() *ConnectorListURLParams {
	return &r.urlParams
}

type ConnectorListURLParams struct {
	ConnectorID string
	SubElement  string
}

func (s *ConnectorService) NewListRequestBody() EmptyRequestBody {
	return EmptyRequestBody{}
}

type ConnectorListResponseBody struct {
	Skip int         `json:"skip"`
	Take int         `json:"take"`
	Rows interface{} `json:"rows"`
}
