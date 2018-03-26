package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"go/format"

	afas "github.com/tim-online/go-afas-profit-rest"
)

type Generator struct {
}

func (g Generator) NewAPI() *afas.API {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(false)
	return api
}

func (g Generator) All() error {
	api := g.NewAPI()
	req := api.Meta.NewListConnectorsRequest()
	resp, err := req.Do()
	if err != nil {
		return err
	}

	files := map[string]io.Reader{}

	getGenerator := GetGenerator{}
	getFiles, err := getGenerator.Generate(resp.GetConnectors)
	if err != nil {
		return err
	}
	for k, v := range getFiles {
		files[k] = v
	}

	updateGenerator := UpdateGenerator{}
	updateFiles, err := updateGenerator.Generate(resp.UpdateConnectors)
	if err != nil {
		return err
	}
	for k, v := range updateFiles {
		files[k] = v
	}

	for f, r := range files {
		// format code
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		formatted, err := format.Source(b)
		if err != nil {
			log.Println(string(b))
			return err
		}

		err = ioutil.WriteFile(f, r)
		if err != nil {
			return err
		}
	}

	return nil
}

// % G-rekening -> GRekening
// Loonsom (%) -> Loonsom
// Organisatie/Persoon -> OrganisatiePersoon
// CVEtotaal -> CVETotaal
//
func normalizeIdentifier(id string) string {
	id = UpperCamelCase(id)
	return removeNonAlphanumeric(id)
}

func removeNonAlphanumeric(s string) string {
	buffer := make([]rune, 0, len(s))
	for _, curr := range s {
		if !unicode.IsNumber(curr) && !unicode.IsLetter(curr) {
			continue
		}
		buffer = append(buffer, curr)
	}

	return string(buffer)
}
