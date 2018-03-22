package afas

import (
	"net/http"
	"net/url"
)

func (s *MetaService) NewListConnectorsRequest() MetaListConnectorsRequest {
	return MetaListConnectorsRequest{
		api:         s.api,
		method:      http.MethodGet,
		urlParams:   MetaListConnectorsURLParams{},
		queryParams: MetaListConnectorsQueryParams{},
	}
}

type MetaListConnectorsRequest struct {
	api         *API
	method      string
	urlParams   MetaListConnectorsURLParams
	queryParams MetaListConnectorsQueryParams
}

func (r *MetaListConnectorsRequest) Method() string {
	return r.method
}

func (r *MetaListConnectorsRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetaListConnectorsRequest) URL() url.URL {
	path := "metainfo"
	return r.api.GetEndpointURL(path)
}

func (r *MetaListConnectorsRequest) Do() (MetaListConnectorsResponseBody, error) {
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

func (r *MetaListConnectorsRequest) NewResponseBody() *MetaListConnectorsResponseBody {
	return &MetaListConnectorsResponseBody{}
}

type MetaListConnectorsURLParams struct {
}

type MetaListConnectorsQueryParams struct {
}

// {
//   "updateConnectors": [
//     {
//       "id": "PtProject",
//       "description": "Project"
//     }
//   ],
//   "getConnectors": [
//     {
//       "id": "OmniBoost_ProjectMaintenance_Contacts",
//       "description": "OmniBoost_ProjectMaintenance_Contacts"
//     },
//     {
//       "id": "OmniBoost_ProjectMaintenance_ProjectGroups",
//       "description": "OmniBoost_ProjectMaintenance_ProjectGroups"
//     },
//     {
//       "id": "OmniBoost_ProjectMaintenance_Projects",
//       "description": "OmniBoost_ProjectMaintenance_Projects"
//     }
//   ]
// }

type MetaListConnectorsResponseBody struct {
	UpdateConnectors UpdateConnectors `json:"updateConnectors"`
	GetConnectors    GetConnectors    `json:"getConnectors"`
}
