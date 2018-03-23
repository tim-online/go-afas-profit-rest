package afas_test

import (
	"log"
	"os"
	"testing"

	"github.com/tim-online/go-afas-profit-rest"
)

func TestMetaListConnectors(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.Meta.NewListConnectorsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}

func TestMetaDescribeGetConnector(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.Meta.NewDescribeGetConnectorRequest()
	// req.URLParams().ConnectorID = "OmniBoost_ProjectMaintenance_Contacts"
	req.URLParams().ConnectorID = "OMNI_verkooprelatie"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}

// func TestMetaDescribeUpdateConnector(t *testing.T) {
// 	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
// 	token := os.Getenv("AFAS_TOKEN")
// 	api := afas.NewAPI(nil, accountNumber, token)
// 	api.SetDebug(true)

// 	req := api.Meta.NewDescribeUpdateConnectorRequest()
// 	req.URLParams().ConnectorID = "PtProject"
// 	resp, err := req.Do()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(resp)
// }
