package afas_test

import (
	"log"
	"os"
	"testing"

	afas "github.com/tim-online/go-afas-profit-rest"
)

func TestOMNIVerkooprelatieList(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.OMNIVerkooprelatie().NewListRequest()
	_, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	for _, o := range req.ResponseBody().Rows {
		log.Println(o.NummerDebiteur)
		log.Println(o.NummerOrgPers)
		log.Println(o.OrganisatiePersoonCode)
	}
}
