package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"

	"github.com/dave/jennifer/jen"
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
	var st *jen.Statement
	file := jen.NewFile("main")

	api := g.NewAPI()
	req := api.Meta.NewListConnectorsRequest()
	resp, err := req.Do()
	if err != nil {
		return err
	}

	st, err = g.generateGetConnectors(resp.GetConnectors)
	if err != nil {
		return err
	}
	if st != nil {
		file.Add(st)
	}

	st, err = g.generateUpdateConnectors(resp.UpdateConnectors)
	if err != nil {
		return err
	}
	if st != nil {
		file.Add(st)
	}

	// render package
	buf := &bytes.Buffer{}
	err = file.Render(buf)
	if err != nil {
		return err
	}
	fmt.Println(buf)

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
