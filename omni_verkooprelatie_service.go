package afas

func (api *API) OMNIVerkooprelatie() *OMNIVerkooprelatieService {
	return &OMNIVerkooprelatieService{api: api}

}

type OMNIVerkooprelatieService struct {
	api *API
}
