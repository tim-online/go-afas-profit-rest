package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	afas "github.com/tim-online/go-afas-profit-rest"
)

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
	return generateUpdateConnectorObject(d.UpdateConnectorObject)
}

func generateUpdateConnectorStructFields(d afas.UpdateConnectorObject) ([]jen.Code, error) {
	fields := []jen.Code{}
	for _, f := range d.Fields {
		sf, err := generateUpdateConnectorStructFieldFromField(f)
		if err != nil {
			return fields, err
		}
		if sf != nil {
			fields = append(fields, sf)
		}
	}

	for _, o := range d.Objects {
		sf, err := generateUpdateConnectorStructFieldFromObject(o)
		if err != nil {
			return fields, err
		}
		if sf != nil {
			fields = append(fields, sf)
		}
	}

	return fields, nil
}

func generateUpdateConnectorStructFieldFromField(f afas.UpdateConnectorField) (*jen.Statement, error) {
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

func generateUpdateConnectorStructFieldFromObject(o afas.UpdateConnectorObject) (*jen.Statement, error) {
	fID := normalizeIdentifier(o.Name)
	sf := jen.Id(fID).Id(fID)

	// json tags
	sf.Tag(map[string]string{"json": o.Name})

	// comment behind struct field
	sf.Comment(o.Name)

	return sf, nil
}

func generateUpdateConnectorObject(o afas.UpdateConnectorObject) (*jen.Statement, error) {
	g := jen.Custom(jen.Options{
		Open:      "",
		Close:     "",
		Separator: "",
		Multi:     true,
	})

	fields, err := generateUpdateConnectorStructFields(o)
	if err != nil {
		return g, err
	}

	// generate struct
	id := normalizeIdentifier(o.Name)
	g.Type().Id(id).Struct(fields...).Line().Line()

	for _, o2 := range o.Objects {
		st, err := generateUpdateConnectorObject(o2)
		if err != nil {
			return g, err
		}
		g.Add(st)
	}

	return g, nil
}

func generateUpdateConnectorStructMethods(o afas.UpdateConnectorObject) (io.Reader, error) {
	id := normalizeIdentifier(o.Name)
	first := string([]rune(id)[0])

	fields := []string{}
	dbIDField := ""
	for _, f := range o.Fields {
		fieldID := normalizeIdentifier(f.FieldID)
		fieldName := normalizeIdentifier(f.Label)
		if fieldID == "DbId" {
			dbIDField = fieldName
			continue
		}
		fields = append(fields, fieldID)
	}

	objects := []string{}
	for _, o := range o.Objects {
		id := normalizeIdentifier(o.Name)
		objects = append(objects, o.Name)
	}

	data := struct {
		TypeVariable string
		Type         string
		DBIDField    string
		Fields       []string
		Objects      []string
	}{
		TypeVariable: first,
		Type:         id,
		DBIDField:    dbIDField,
		Fields:       fields,
		Objects:      objects,
	}

	tmpl, err := template.ParseFiles("generate/update_connector_struct_methods.tmpl")
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})
	r, w := io.Pipe()
	err = tmpl.Execute(w, data)
	return r, nil
}
