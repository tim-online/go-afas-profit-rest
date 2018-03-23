package afas

func NewOMNIVerkooprelatieService(api *API) *OMNIVerkooprelatieService {
	return &OMNIVerkooprelatieService{api: api}
}

type OMNIVerkooprelatieService struct {
	api *API
}
