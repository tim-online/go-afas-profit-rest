package afas_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/aodin/date"
	afas "github.com/tim-online/go-afas-profit-rest"
)

func TestFiEntryParInsert(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.FiEntryPar().NewInsertRequest()
	rb := req.RequestBody()
	rb.Boekjaar = 2018
	rb.Periode = 2
	rb.NummerAdministratie = 4
	rb.Dagboek = "80"

	entry := afas.FiEntries{
		KenmerkRekening: "1",
		Rekeningnummer:  "10025",
		Boekstuknummer:  "1",
		DatumBoeking:    date.FromTime(time.Now()),
		Boekstukdatum:   date.FromTime(time.Now()),
	}
	rb.FiEntries = append(rb.FiEntries, entry)

	entry = afas.FiEntries{
		KenmerkRekening: "1",
		Rekeningnummer:  "10025",
		Boekstuknummer:  "1",
		DatumBoeking:    date.FromTime(time.Now()),
		Boekstukdatum:   date.FromTime(time.Now()),
	}
	rb.FiEntries = append(rb.FiEntries, entry)

	// b, _ := json.MarshalIndent(rb.FiEntries, "", "  ")
	// log.Fatal(string(b))

	// b, _ := json.MarshalIndent(rb, "", "  ")
	// log.Fatal(string(b))

	_, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v", req.ResponseBody().UniID)
	log.Printf("%+v", req.ResponseBody().EnNo)
	log.Fatal("TEST")
}
