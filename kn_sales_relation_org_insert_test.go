package afas_test

import (
	"os"
	"testing"

	afas "github.com/tim-online/go-afas-profit-rest"
)

func TestKnSalesRelationOrgInsert(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.KnSalesRelationOrg().NewInsertRequest()
	req.RequestBody().NummerDebiteur = "12345test"
	req.RequestBody().KnOrganisation.Naam = "Leon Bogaert"
	req.RequestBody().KnOrganisation.OrganisatieVergelijkenOp = "3"
	_, err := req.Do()
	if err != nil {
		t.Error(err)
	}
}
