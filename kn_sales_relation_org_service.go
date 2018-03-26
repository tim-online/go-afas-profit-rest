package afas

func (api *API) KnSalesRelationOrg() *KnSalesRelationOrgService {
	return &KnSalesRelationOrgService{api: api}
}

type KnSalesRelationOrgService struct {
	api *API
}
