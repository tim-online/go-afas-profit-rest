package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	afas "github.com/tim-online/go-afas-profit-rest"
)

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
