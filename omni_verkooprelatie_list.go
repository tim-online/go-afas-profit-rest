package afas

type OMNIVerkooprelatieListRequest struct {
	ConnectorListRequest
}

type OMNIVerkooprelatieListResponseBody struct {
	ConnectorListResponseBody

	Rows []OMNIVerkooprelatie `json:"rows"`
}

func (r *OMNIVerkooprelatieListRequest) ResponseBody() *OMNIVerkooprelatieListResponseBody {
	rb := r.ConnectorListRequest.ResponseBody()
	rows := rb.Rows.(*[]OMNIVerkooprelatie)
	return &OMNIVerkooprelatieListResponseBody{
		ConnectorListResponseBody: *rb,
		Rows: *rows,
	}
}

func (s *OMNIVerkooprelatieService) NewListRequest() OMNIVerkooprelatieListRequest {
	r := s.api.Connector.NewListRequest()

	rows := []OMNIVerkooprelatie{}
	r.ResponseBody().Rows = &rows

	r.URLParams().ConnectorID = "OMNI_verkooprelatie"
	return OMNIVerkooprelatieListRequest{ConnectorListRequest: r}
}
