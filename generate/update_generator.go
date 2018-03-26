package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	afas "github.com/tim-online/go-afas-profit-rest"
)

type UpdateGenerator struct {
}

func (g UpdateGenerator) NewAPI() *afas.API {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(false)
	return api
}

func (g UpdateGenerator) Generate(connectors afas.UpdateConnectors) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}
	api := g.NewAPI()

	for _, c := range connectors {
		req := api.Meta.NewDescribeUpdateConnectorRequest()
		req.URLParams().ConnectorID = c.ID
		resp, err := req.Do()
		if err != nil {
			return files, err
		}

		filenameBase := SnakeCase(resp.Name)
		structs, err := generateUpdateConnectorResponseStructs(resp)
		if err != nil {
			return files, err
		}

		r, err := g.GenerateTypesCode(structs)
		if err != nil {
			return files, err
		}
		filename := fmt.Sprintf("%s_types.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateInsertCode(structs[0])
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_insert.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateServiceCode(structs[0])
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_service.go", filenameBase)
		files[filename] = r
	}

	return files, nil
}

func (g UpdateGenerator) GenerateTypesCode(structs []UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_types.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, structs)
	if err != nil {
		return buf, err
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return buf, err
	}

	return bytes.NewBuffer(b), nil
}

func (g UpdateGenerator) GenerateInsertCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_insert.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, st)
	if err != nil {
		return buf, err
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return buf, err
	}

	return bytes.NewBuffer(b), nil
}

func (g UpdateGenerator) GenerateServiceCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_service.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, st)
	if err != nil {
		return buf, err
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return buf, err
	}

	return bytes.NewBuffer(b), nil
}

func generateUpdateConnectorResponseStructs(d afas.MetaDescribeUpdateConnectorResponseBody) ([]UpdateConnectorStruct, error) {
	return generateUpdateConnectorObjects(d.UpdateConnectorObject)
}

func generateUpdateConnectorStructFields(d afas.UpdateConnectorObject) (UpdateConnectorStructFields, error) {
	fields := UpdateConnectorStructFields{}
	for _, f := range d.Fields {
		sf, err := generateUpdateConnectorStructFieldFromField(f)
		if err != nil {
			return fields, err
		}
		fields = append(fields, sf)
	}

	for _, o := range d.Objects {
		sf, err := generateUpdateConnectorStructFieldFromObject(o)
		if err != nil {
			return fields, err
		}
		fields = append(fields, sf)
	}

	return fields, nil
}

func generateUpdateConnectorStructFieldFromField(f afas.UpdateConnectorField) (UpdateConnectorStructField, error) {
	// do type
	typ := ""
	switch f.DataType {
	case "string":
		typ = "string"
	case "int":
		typ = "int"
	case "boolean":
		typ = "bool"
	case "date":
		typ = "time.Time"
	case "decimal":
		typ = "*apd.Decimal"
	case "blob":
		typ = "[]byte"
	default:
		return UpdateConnectorStructField{}, errors.Errorf("Unkown datatype: %s", f.DataType)
	}

	name := normalizeIdentifier(f.Label)

	if len(f.Values) > 0 {
		// @TODO: soortement van enum maken?
	}

	jsonName := f.FieldID

	// json tags
	tags := ""
	if f.Notzero {
		tags = fmt.Sprintf(`json:"%s,omitempty"`, jsonName)
	} else {
		tags = fmt.Sprintf(`json:"%s"`, jsonName)
	}

	// comment behind struct field
	comment := f.Label

	return UpdateConnectorStructField{
		Comment:  comment,
		Name:     name,
		Tags:     tags,
		Type:     typ,
		JSONName: jsonName,
	}, nil
}

func generateUpdateConnectorStructFieldFromObject(o afas.UpdateConnectorObject) (UpdateConnectorStructField, error) {
	name := normalizeIdentifier(o.Name)
	typ := name

	// json tags
	tags := fmt.Sprintf(`json:"%s"`, o.Name)

	// comment behind struct field
	comment := o.Name

	return UpdateConnectorStructField{
		Comment: comment,
		Name:    name,
		Tags:    tags,
		Type:    typ,
	}, nil
}

type UpdateConnectorStruct struct {
	Comment   string
	Name      string
	Variable  string
	Fields    UpdateConnectorStructFields
	Objects   []string
	DBIDField string
}

type UpdateConnectorStructFields []UpdateConnectorStructField

type UpdateConnectorStructField struct {
	Name     string
	Type     string
	Tags     string
	Comment  string
	JSONName string
}

func generateUpdateConnectorObjects(o afas.UpdateConnectorObject) ([]UpdateConnectorStruct, error) {
	structs := []UpdateConnectorStruct{}
	st, err := generateUpdateConnectorObject(o)
	if err != nil {
		return structs, err
	}
	structs = append(structs, st)

	for _, o := range o.Objects {
		substructs, err := generateUpdateConnectorObjects(o)
		if err != nil {
			return structs, err
		}
		structs = append(structs, substructs...)
	}

	return structs, nil
}

func generateUpdateConnectorObject(o afas.UpdateConnectorObject) (UpdateConnectorStruct, error) {
	fields, err := generateUpdateConnectorStructFields(o)
	if err != nil {
		return UpdateConnectorStruct{}, err
	}

	objects := []string{}
	for _, o := range o.Objects {
		objects = append(objects, o.Name)
	}

	dbIDField := ""
	for _, f := range fields {
		if f.JSONName == "dbId" {
			dbIDField = f.Name
		}
	}

	name := normalizeIdentifier(o.Name)
	variable := strings.ToLower(string([]rune(name)[0]))

	return UpdateConnectorStruct{
		Comment:   o.Name,
		Name:      name,
		Variable:  variable,
		Fields:    fields,
		DBIDField: dbIDField,
		Objects:   objects,
	}, nil
}
