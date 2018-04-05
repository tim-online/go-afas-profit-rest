package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	afas "github.com/tim-online/go-afas-profit-rest"
)

type GetGenerator struct {
}

func (g GetGenerator) NewAPI() *afas.API {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(false)
	return api
}

func (g GetGenerator) Generate(connectors afas.GetConnectors) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}
	api := g.NewAPI()

	for _, c := range connectors {
		req := api.Meta.NewDescribeGetConnectorRequest()
		req.URLParams().ConnectorID = c.ID
		resp, err := req.Do()
		if err != nil {
			return files, err
		}

		filenameBase := SnakeCase(resp.Name)
		structs, err := generateGetConnectorResponseStructs(resp)
		if err != nil {
			return files, err
		}

		r, err := g.GenerateListCode(structs[0])
		if err != nil {
			return files, err
		}
		filename := fmt.Sprintf("%s_list.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateTestListCode(structs[0])
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_list_test.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateServiceCode(structs[0])
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_service.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateTypesCode(structs)
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_types.go", filenameBase)
		files[filename] = r
	}

	return files, nil
}

func (g GetGenerator) GenerateTypesCode(structs []GetConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/get_connector_types.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, structs)
	return buf, err
}

func (g GetGenerator) GenerateListCode(st GetConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/get_connector_list.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, st)
	return buf, err
}

func (g GetGenerator) GenerateTestListCode(st GetConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/get_connector_list_test.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, st)
	return buf, err
}

func (g GetGenerator) GenerateServiceCode(st GetConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/get_connector_service.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, st)
	return buf, err
}

func generateGetConnectorResponseStructs(d afas.MetaDescribeGetConnectorResponseBody) ([]GetConnectorStruct, error) {
	st, err := generateGetConnectorResponseStruct(d)
	return []GetConnectorStruct{st}, err
}

func generateGetConnectorResponseStruct(d afas.MetaDescribeGetConnectorResponseBody) (GetConnectorStruct, error) {
	fields, err := generateGetConnectorStructFields(d)
	if err != nil {
		return GetConnectorStruct{}, err
	}

	// struct comment
	comment := d.Description

	// struct with fields
	name := normalizeIdentifier(d.Name)
	variable := strings.ToLower(string([]rune(name)[0]))

	isSlice := false
	plural := ""
	if IsPlural(name) {
		isSlice = true
		plural = name
		name = GetSingular(name)
	}

	return GetConnectorStruct{
		Comment:  comment,
		Name:     name,
		ID:       d.Name,
		Variable: variable,
		Fields:   fields,
		IsSlice:  isSlice,
		Plural:   plural,
	}, nil
}

func generateGetConnectorStructFields(d afas.MetaDescribeGetConnectorResponseBody) (GetConnectorStructFields, error) {
	fields := GetConnectorStructFields{}
	for _, f := range d.Fields {
		name := normalizeIdentifier(f.ID)

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
			typ = "float64"
		case "blob":
			typ = "[]byte"
		default:
			return GetConnectorStructFields{}, errors.Errorf("Unkown datatype: %s", f.DataType)
		}

		jsonName := f.ID

		// json tags
		tags := fmt.Sprintf(`json:"%s"`, jsonName)

		// comment behind struct field
		comment := f.Label

		field := GetConnectorStructField{
			Name:     name,
			Type:     typ,
			Tags:     tags,
			Comment:  comment,
			JSONName: jsonName,
		}
		fields = append(fields, field)
	}

	return fields, nil
}

type GetConnectorStruct struct {
	Comment  string
	Name     string
	ID       string
	Variable string
	Fields   GetConnectorStructFields
	IsSlice  bool
	Plural   string
}

type GetConnectorStructFields []GetConnectorStructField

type GetConnectorStructField struct {
	Name     string
	Type     string
	Tags     string
	Comment  string
	JSONName string
}
