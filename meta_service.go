package afas

func NewMetaService(api *API) *MetaService {
	return &MetaService{api: api}
}

type MetaService struct {
	api *API
}
