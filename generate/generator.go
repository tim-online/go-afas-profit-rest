package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
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

func (g Generator) generateGetConnectors(connectors afas.GetConnectors) (*jen.Statement, error) {
	api := g.NewAPI()

	file := jen.Custom(jen.Options{
		Open:      "",
		Close:     "",
		Separator: "",
		Multi:     true,
	})

	for _, c := range connectors {
		req := api.Meta.NewDescribeGetConnectorRequest()
		req.URLParams().ConnectorID = c.ID
		resp, err := req.Do()
		if err != nil {
			return file, err
		}

		st, err := generateGetConnectorResponseStruct(resp)
		if err != nil {
			return file, err
		}

		file.Add(st)
	}

	return file, nil
}

func generateGetConnectorResponseStruct(d afas.MetaDescribeGetConnectorResponseBody) (*jen.Statement, error) {
	g := jen.Custom(jen.Options{
		Open:      "",
		Close:     "",
		Separator: "",
		Multi:     true,
	})

	fields := []jen.Code{}
	for _, f := range d.Fields {
		fID := normalizeIdentifier(f.ID)
		sf := jen.Id(fID)

		// do type
		switch f.DataType {
		case "string":
			sf = sf.String()
		case "int":
			sf = sf.Int()
		case "boolean":
			sf = sf.Bool()
		case "date":
			sf = sf.Qual("time", "Time")
		default:
			return g, errors.Errorf("Unkown datatype: %s", f.DataType)
		}

		// json tags
		sf = sf.Tag(map[string]string{"json": f.ID})

		// comment behind struct field
		sf.Comment(f.Label)

		fields = append(fields, sf)
	}

	// struct comment
	g.Comment(d.Description).Line()

	// struct with fields
	id := normalizeIdentifier(d.Name)
	g.Type().Id(id).Struct(fields...).Line()

	return g, nil
}

func (g Generator) generateUpdateConnectors(connectors afas.UpdateConnectors) (*jen.Statement, error) {
	api := g.NewAPI()

	file := jen.Custom(jen.Options{
		Open:      "",
		Close:     "",
		Separator: "",
		Multi:     true,
	})

	for _, c := range connectors {
		req := api.Meta.NewDescribeUpdateConnectorRequest()
		req.URLParams().ConnectorID = c.ID
		resp, err := req.Do()
		if err != nil {
			return file, err
		}

		st, err := generateUpdateConnectorResponseStruct(resp)
		if err != nil {
			return file, err
		}

		file.Add(st)
	}

	return file, nil
}

func generateUpdateConnectorResponseStruct(d afas.MetaDescribeUpdateConnectorResponseBody) (*jen.Statement, error) {
	g := jen.Custom(jen.Options{
		Open:      "",
		Close:     "",
		Separator: "",
		Multi:     true,
	})

	fields := []jen.Code{}
	for _, f := range d.Fields {
		sf, err := generateUpdateConnectorField(f)
		if err != nil {
			return g, err
		}
		if sf != nil {
			fields = append(fields, sf)
		}
	}

	for _, o := range d.Objects {
		sf, err := generateUpdateConnectorObject(o)
		if err != nil {
			return g, err
		}
		if sf != nil {
			fields = append(fields, sf)
		}
	}

	// struct comment
	g.Comment(d.Description).Line()

	// struct with fields
	id := normalizeIdentifier(d.Name)
	g.Type().Id(id).Struct(fields...).Line()

	return g, nil
}

func generateUpdateConnectorField(f afas.UpdateConnectorField) (*jen.Statement, error) {
	fID := normalizeIdentifier(f.Label)
	sf := jen.Id(fID)

	// do type
	switch f.DataType {
	case "string":
		sf = sf.String()
	case "int":
		sf = sf.Int()
	case "boolean":
		sf = sf.Bool()
	case "date":
		sf = sf.Qual("time", "Time")
	case "decimal":
		sf = sf.Op("*").Qual("github.com/cockroachdb/apd", "Decimal")
	case "blob":
		sf = sf.Index().Byte()
	default:
		return sf, errors.Errorf("Unkown datatype: %s", f.DataType)
	}

	if len(f.Values) > 0 {
		// @TODO: soortement van enum maken?
	}

	// json tags
	tags := map[string]string{}
	if f.Notzero {
		tags["json"] = fmt.Sprintf("%s,omitempty", f.FieldID)
	} else {
		tags["json"] = f.FieldID
	}
	sf = sf.Tag(tags)

	// comment behind struct field
	sf.Comment(f.Label)
	return sf, nil
}

func generateUpdateConnectorObject(o afas.UpdateConnectorObject) (*jen.Statement, error) {
	g := jen.Custom(jen.Options{
		Open:      "",
		Close:     "",
		Separator: "",
		Multi:     true,
	})

	fields := []jen.Code{}
	for _, f := range o.Fields {
		sf, err := generateUpdateConnectorField(f)
		if err != nil {
			return g, err
		}
		if sf != nil {
			fields = append(fields, sf)
		}
	}

	// inline struct with fields
	id := normalizeIdentifier(o.Name)
	g.Id(id).Struct(fields...)

	// json tags
	g.Tag(map[string]string{"json": o.Name})

	return g, nil
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
