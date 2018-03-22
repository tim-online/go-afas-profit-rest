package afas

func NewCRMService(api *API) *CRMService {
	return &CRMService{api: api}
}

type CRMService struct {
	api *API
}
