package afas_test

import (
	"log"
	"os"
	"testing"

	afas "github.com/tim-online/go-afas-profit-rest"
)

// Verkooprelatie (Rapport)
type OMNI_verkooprelatie struct {
	NummerDebiteur         string `json:"Nummer_debiteur"`          // Nummer debiteur
	NummerOrgPers          string `json:"Nummer_org-pers"`          // Nummer org-pers
	OrganisatiePersoonCode string `json:"Organisatie_persoon_code"` // Organisatie/persoon code
}

func TestConnectorList(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.Connector.NewListRequest()
	rows := []OMNI_verkooprelatie{}
	req.ResponseBody().Rows = &rows
	req.URLParams().ConnectorID = "OMNI_verkooprelatie"
	_, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	for _, o := range rows {
		log.Println(o.NummerDebiteur)
		log.Println(o.NummerOrgPers)
		log.Println(o.OrganisatiePersoonCode)
	}
	log.Fatal("TEST")
}

// func TestConnectorInsert(t *testing.T) {
// 	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
// 	token := os.Getenv("AFAS_TOKEN")
// 	api := afas.NewAPI(nil, accountNumber, token)
// 	api.SetDebug(true)

// 	req := api.Connector.NewInsertRequest()
// 	ptProject := struct {
// 		Name string
// 		Test int
// 	}{
// 		Name: "test",
// 		Test: 12,
// 	}
// 	req.SetRequestBody(ptProject)
// 	req.URLParams().ConnectorID = "PtProject"
// 	resp, err := req.Do()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(resp)
// }
