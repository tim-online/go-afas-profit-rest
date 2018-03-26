package afas

import (
	"net/http"
	"net/url"
	"strings"
)

func (s *MetaService) NewDescribeUpdateConnectorRequest() MetaDescribeUpdateConnectorRequest {
	return MetaDescribeUpdateConnectorRequest{
		api:         s.api,
		method:      http.MethodGet,
		urlParams:   MetaDescribeUpdateConnectorURLParams{},
		queryParams: MetaDescribeUpdateConnectorQueryParams{},
	}
}

type MetaDescribeUpdateConnectorRequest struct {
	api         *API
	method      string
	urlParams   MetaDescribeUpdateConnectorURLParams
	queryParams MetaDescribeUpdateConnectorQueryParams
}

func (r *MetaDescribeUpdateConnectorRequest) Method() string {
	return r.method
}

func (r *MetaDescribeUpdateConnectorRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetaDescribeUpdateConnectorRequest) URL() url.URL {
	path := "metainfo/update/{connectorid}"
	path = strings.Replace(path, "{connectorid}", r.urlParams.ConnectorID, 1)
	return r.api.GetEndpointURL(path)
}

func (r *MetaDescribeUpdateConnectorRequest) Do() (MetaDescribeUpdateConnectorResponseBody, error) {
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

func (r *MetaDescribeUpdateConnectorRequest) NewResponseBody() *MetaDescribeUpdateConnectorResponseBody {
	return &MetaDescribeUpdateConnectorResponseBody{}
}

func (r *MetaDescribeUpdateConnectorRequest) QueryParams() *MetaDescribeUpdateConnectorQueryParams {
	return &r.queryParams
}

type MetaDescribeUpdateConnectorQueryParams struct {
}

func (r *MetaDescribeUpdateConnectorRequest) URLParams() *MetaDescribeUpdateConnectorURLParams {
	return &r.urlParams
}

type MetaDescribeUpdateConnectorURLParams struct {
	ConnectorID string
}

// {
//   "name": "OmniBoost_ProjectMaintenance_Contacts",
//   "description": "Contacten",
//   "fields": [
//     {
//       "id": "Nummer_debiteur",
//       "fieldId": "U003",
//       "dataType": "string",
//       "label": "Debiteur nummer",
//       "length": 15,
//       "controlType": 5,
//       "decimals": 0,
//       "decimalsFieldId": ""
//     },
//     {
//       "id": "Debiteur_naam",
//       "fieldId": "U004",
//       "dataType": "string",
//       "label": "Debiteur naam",
//       "length": 255,
//       "controlType": 1,
//       "decimals": 0,
//       "decimalsFieldId": ""
//     }
// }

type MetaDescribeUpdateConnectorResponseBody struct {
	ID          string `json:"id"`
	Description string `json:"description"`

	UpdateConnectorObject
}

type UpdateConnectorField struct {
	FieldID        string `json:"fieldId"`
	PrimaryKey     bool   `json:"primaryKey"`
	DataType       string `json:"dataType"`
	Label          string `json:"label"`
	Mandatory      bool   `json:"mandatory"`
	Length         int    `json:"length"`
	Decimals       int    `json:"decimals"`
	DecimalFieldID string `json:"decimalFieldId"`
	Notzero        bool   `json:"notzero"`
	ControlType    int    `json:"controlType"`
	Values         []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"values"`
}

type UpdateConnectorObject struct {
	Name    string                  `json:"name"`
	Fields  []UpdateConnectorField  `json:"fields"`
	Objects []UpdateConnectorObject `json:"objects"`
}
