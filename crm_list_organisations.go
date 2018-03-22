package afas

import (
	"net/http"
	"net/url"
)

func (s *CRMService) NewListOrganisationsRequest() CRMListOrganisationsRequest {
	return CRMListOrganisationsRequest{
		api:       s.api,
		method:    http.MethodGet,
		urlParams: CRMListOrganisationsURLParams{},
		queryParams: CRMListOrganisationsQueryParams{
			Skip:            0,
			Take:            0,
			OrderByFieldIDs: "",
			FilterFieldIDs:  "",
		},
	}
}

type CRMListOrganisationsRequest struct {
	api         *API
	method      string
	urlParams   CRMListOrganisationsURLParams
	queryParams CRMListOrganisationsQueryParams
}

func (r *CRMListOrganisationsRequest) Method() string {
	return r.method
}

func (r *CRMListOrganisationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *CRMListOrganisationsRequest) URL() url.URL {
	path := "clients.json"
	return r.api.GetEndpointURL(path)
}

func (r *CRMListOrganisationsRequest) Do() (CRMListOrganisationsResponseBody, error) {
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

func (r *CRMListOrganisationsRequest) NewResponseBody() *CRMListOrganisationsResponseBody {
	return &CRMListOrganisationsResponseBody{}
}

type CRMListOrganisationsURLParams struct {
}

type CRMListOrganisationsQueryParams struct {
	Skip            int    `schema:"skip,omitempty"`
	Take            int    `schema:"take,omitempty"`
	OrderByFieldIDs string `schema:"orderbyfieldids,omitempty"`
	FilterFieldIDs  string `schema:"filterfieldsids,omitempty"`
}

type CRMListOrganisationsResponseBody struct {
}
