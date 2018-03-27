// DO NOT EDIT: generated by github.com/tim-online/go-afas-profit-rest/generate

package afas

import "encoding/json"

type FiEntryParInsertRequest struct {
	FiEntryParService
	ConnectorInsertRequest
}

type FiEntryParInsertResponseBody struct {
	*ConnectorInsertResponseBody

	UniID string `json:"Unid"`
	EnNo  string `json:"EnNo"`
}

type FiEntryParInsertRequestBody struct {
	FiEntryPar `json:"FiEntryPar"`
}

func (f *FiEntryParInsertRequest) ResponseBody() *FiEntryParInsertResponseBody {
	type alias FiEntryParInsertResponseBody

	rb := FiEntryParInsertResponseBody{
		ConnectorInsertResponseBody: f.ConnectorInsertRequest.ResponseBody(),
	}

	wrapper := struct {
		FiEntryParInsertResponseBody alias `json:"FiEntryPar"`
	}{
		FiEntryParInsertResponseBody: alias(rb),
	}

	json.Unmarshal(f.ConnectorInsertRequest.ResponseBody().Results, &wrapper)
	rb = FiEntryParInsertResponseBody(wrapper.FiEntryParInsertResponseBody)
	return &rb
}

// Wrap extra object around root type:
// {
//   "root": {}
// }
func (f FiEntryParInsertRequestBody) MarshalJSON() ([]byte, error) {
	type alias FiEntryParInsertRequestBody
	wrapper := struct {
		FiEntryParInsertRequestBody alias `json:"FiEntryPar"`
	}{FiEntryParInsertRequestBody: alias(f)}
	return json.Marshal(wrapper)
}

func (r *FiEntryParInsertRequest) RequestBody() *FiEntryParInsertRequestBody {
	if body, ok := r.ConnectorInsertRequest.RequestBody().(*FiEntryParInsertRequestBody); ok {
		return body
	}

	body := &FiEntryParInsertRequestBody{}
	r.ConnectorInsertRequest.SetRequestBody(body)
	return body
}

func (s *FiEntryParService) NewInsertRequest() FiEntryParInsertRequest {
	r := s.api.Connector.NewInsertRequest()
	r.URLParams().ConnectorID = "FiEntries"
	return FiEntryParInsertRequest{ConnectorInsertRequest: r}
}