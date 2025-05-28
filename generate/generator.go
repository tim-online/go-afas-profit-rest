package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"unicode"

	"golang.org/x/tools/imports"

	afas "github.com/tim-online/go-afas-profit-rest"
)

type Generator struct {
}

func (g Generator) NewAPI() *afas.API {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	if accountNumber == "" {
		log.Fatal("AFAS_ACCOUNTNUMBER can't be empty")
	}

	token := os.Getenv("AFAS_TOKEN")
	if token == "" {
		log.Fatal("AFAS_TOKEN can't be empty")
	}

	api := afas.NewAPI(nil, accountNumber, token)

	debug := os.Getenv("AFAS_DEBUG")
	if debug != "" {
        api.SetDebug(true)
	}

	baseURL := os.Getenv("AFAS_BASE_URL")
	if baseURL != "" {
		u, err := url.Parse(baseURL)
		if err !=  nil {
			log.Fatal(err)
		}
		api.SetBaseURL(*u)
	}
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

	getGenerator := GetGenerator{api: api}
	getFiles, err := getGenerator.Generate(resp.GetConnectors)
	if err != nil {
		return err
	}
	for k, v := range getFiles {
		files[k] = v
	}

	updateGenerator := UpdateGenerator{api: api}
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

		// remove Vim modeline
		b = bytes.Replace(b, []byte("// vim: ft=gotexttmpl noet"), []byte{}, -1)

		formatted, err := imports.Process(f, b, nil)
		if err != nil {
			log.Println(string(b))
			return err
		}

		err = ioutil.WriteFile(f, formatted, 0644)
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
