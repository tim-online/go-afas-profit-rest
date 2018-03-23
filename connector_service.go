package afas

func NewConnectorService(api *API) *ConnectorService {
	return &ConnectorService{api: api}
}

type ConnectorService struct {
	api *API
}
