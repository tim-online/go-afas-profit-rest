package afas

type KnSalesRelationOrgInsertRequest struct {
	KnSalesRelationOrgService
	ConnectorInsertRequest
}

type KnSalesRelationOrgInsertResponseBody struct {
	ConnectorInsertResponseBody

	Rows []KnSalesRelationOrg `json:"rows"`
}

func (r *KnSalesRelationOrgInsertRequest) RequestBody() *KnSalesRelationOrg {
	if body, ok := r.ConnectorInsertRequest.RequestBody().(*KnSalesRelationOrg); ok {
		return body
	}

	body := &KnSalesRelationOrg{}
	r.ConnectorInsertRequest.SetRequestBody(body)
	return body
}

// func (r *KnSalesRelationOrgInsertRequest) ResponseBody() *KnSalesRelationOrgInsertResponseBody {
// 	rb := r.ConnectorInsertRequest.ResponseBody()
// 	rows := rb.Rows.(*[]KnSalesRelationOrg)
// 	return &KnSalesRelationOrgInsertResponseBody{
// 		ConnectorInsertResponseBody: *rb,
// 		Rows: *rows,
// 	}
// }

func (s *KnSalesRelationOrgService) NewInsertRequest() KnSalesRelationOrgInsertRequest {
	r := s.api.Connector.NewInsertRequest()
	r.URLParams().ConnectorID = "KnSalesRelationOrg"
	return KnSalesRelationOrgInsertRequest{ConnectorInsertRequest: r}
}
