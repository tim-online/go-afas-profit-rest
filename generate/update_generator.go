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

type UpdateGenerator struct {
	api *afas.API
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
	api := g.api

	for _, c := range connectors {
		req := api.Meta.NewDescribeUpdateConnectorRequest()
		req.URLParams().ConnectorID = c.ID
		resp, err := req.Do()
		if err != nil {
			return files, err
		}

		st, err := generateUpdateConnectorStruct(resp)
		if err != nil {
			return files, err
		}
		filenameBase := SnakeCase(st.ID)

		r, err := g.GenerateTypesCode(st)
		if err != nil {
			return files, err
		}
		filename := fmt.Sprintf("%s_types.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateInsertCode(st)
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_insert.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateTestInsertCode(st)
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_insert_test.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateUpdateCode(st)
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_update.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateTestUpdateCode(st)
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_update_test.go", filenameBase)
		files[filename] = r

		r, err = g.GenerateServiceCode(st)
		if err != nil {
			return files, err
		}
		filename = fmt.Sprintf("%s_service.go", filenameBase)
		files[filename] = r
	}

	return files, nil
}

func (g UpdateGenerator) GenerateTypesCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_types.go.tmpl")
	if err != nil {
		return buf, err
	}
	err = tmpl.Execute(buf, st)
	return buf, err
}

func (g UpdateGenerator) GenerateInsertCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_insert.go.tmpl")
	if err != nil {
		return buf, err
	}

	data := struct {
		ID string
		UpdateConnectorObjectStruct
	}{
		ID: st.ID,
		UpdateConnectorObjectStruct: st.Objects[0],
	}

	err = tmpl.Execute(buf, data)
	return buf, err
}

func (g UpdateGenerator) GenerateTestInsertCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_insert_test.go.tmpl")
	if err != nil {
		return buf, err
	}

	data := struct {
		ID string
		UpdateConnectorObjectStruct
	}{
		ID: st.ID,
		UpdateConnectorObjectStruct: st.Objects[0],
	}

	err = tmpl.Execute(buf, data)
	return buf, err
}

func (g UpdateGenerator) GenerateUpdateCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_update.go.tmpl")
	if err != nil {
		return buf, err
	}

	data := struct {
		ID string
		UpdateConnectorObjectStruct
	}{
		ID: st.ID,
		UpdateConnectorObjectStruct: st.Objects[0],
	}

	err = tmpl.Execute(buf, data)
	return buf, err
}

func (g UpdateGenerator) GenerateTestUpdateCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_update_test.go.tmpl")
	if err != nil {
		return buf, err
	}

	data := struct {
		ID string
		UpdateConnectorObjectStruct
	}{
		ID: st.ID,
		UpdateConnectorObjectStruct: st.Objects[0],
	}

	err = tmpl.Execute(buf, data)
	return buf, err
}

func (g UpdateGenerator) GenerateServiceCode(st UpdateConnectorStruct) (io.Reader, error) {
	buf := bytes.NewBuffer([]byte{})
	tmpl, err := template.ParseFiles("generate/update_connector_service.go.tmpl")
	if err != nil {
		return buf, err
	}

	data := struct {
		ID string
		UpdateConnectorObjectStruct
	}{
		ID: st.ID,
		UpdateConnectorObjectStruct: st.Objects[0],
	}

	err = tmpl.Execute(buf, data)
	return buf, err
}

func generateUpdateConnectorStruct(d afas.MetaDescribeUpdateConnectorResponseBody) (UpdateConnectorStruct, error) {
	objects, err := generateUpdateConnectorObjects(d.UpdateConnectorObject)
	if err != nil {
		return UpdateConnectorStruct{}, err
	}

	return UpdateConnectorStruct{
		ID:      d.ID,
		Objects: objects,
	}, nil
}

func generateUpdateConnectorObjectStructFields(d afas.UpdateConnectorObject) (UpdateConnectorObjectStructFields, error) {
	fields := UpdateConnectorObjectStructFields{}
	for _, f := range d.Fields {
		sf, err := generateUpdateConnectorObjectStructFieldFromField(f)
		if err != nil {
			return fields, err
		}
		fields = append(fields, sf)
	}

	for _, o := range d.Objects {
		sf, err := generateUpdateConnectorObjectStructFieldFromObject(o)
		if err != nil {
			return fields, err
		}
		fields = append(fields, sf)
	}

	return fields, nil
}

func generateUpdateConnectorObjectStructFieldFromField(f afas.UpdateConnectorField) (UpdateConnectorObjectStructField, error) {
	// do type
	typ := ""
	switch f.DataType {
	case "string":
		typ = "string"
	case "int":
		typ = "int"
	case "boolean":
		if f.Mandatory {
			typ = "bool"
		} else {
			typ = "*bool"
		}
	case "date":
		if f.Mandatory {
			typ = "time.Time"
		} else {
			typ = "*time.Time"
		}
	case "decimal":
		typ = "float64"
	case "blob":
		typ = "[]byte"
	default:
		return UpdateConnectorObjectStructField{}, errors.Errorf("Unkown datatype: %s", f.DataType)
	}

	name := normalizeIdentifier(f.Label)

	if len(f.Values) > 0 {
		// @TODO: soortement van enum maken?
	}

	jsonName := f.FieldID
	if f.PrimaryKey {
		// when a key is primary it gets serialized into {"Element":{"@DbId":""}}
		jsonName = ""
	}

	// json tags
	tags := ""
	if f.Mandatory {
		tags = ""
	} else {
		tags = ",omitempty"
	}

	vr := ValidationRules{
		Required:    f.Mandatory,
		NotZero:     f.NotZero,
		MaxLength:   f.Length,
		NumDecimals: f.Decimals,
	}

	// comment behind struct field
	comment := f.Label

	return UpdateConnectorObjectStructField{
		Comment:         comment,
		Name:            name,
		Tags:            tags,
		Type:            typ,
		JSONName:        jsonName,
		IsObject:        false,
		ValidationRules: vr,
	}, nil
}

func generateUpdateConnectorObjectStructFieldFromObject(o afas.UpdateConnectorObject) (UpdateConnectorObjectStructField, error) {
	name := normalizeIdentifier(o.Name)

	typ := name

	// json tags
	tags := ""

	// comment behind struct field
	comment := o.Name

	return UpdateConnectorObjectStructField{
		Comment:  comment,
		Name:     name,
		Tags:     tags,
		Type:     typ,
		JSONName: name,
		IsObject: true,
	}, nil
}

type UpdateConnectorStruct struct {
	ID      string
	Objects []UpdateConnectorObjectStruct
}

type UpdateConnectorObjectStruct struct {
	Comment   string
	ID        string
	Name      string
	Variable  string
	Fields    UpdateConnectorObjectStructFields
	Objects   []string
	DBIDField string
	IsSlice   bool
	Plural    string
}

type UpdateConnectorObjectStructFields []UpdateConnectorObjectStructField

type UpdateConnectorObjectStructField struct {
	Name            string
	Type            string
	Tags            string
	Comment         string
	JSONName        string
	IsObject        bool
	ValidationRules ValidationRules
}

func generateUpdateConnectorObjects(o afas.UpdateConnectorObject) ([]UpdateConnectorObjectStruct, error) {
	structs := []UpdateConnectorObjectStruct{}

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

		for _, substruct := range substructs {
			structs = append(structs, substruct)
		}
	}

	// remove duplicates
	for i, st := range structs {
		j := 0
		for _, st2 := range structs {
			if st.Name == st2.Name {
				j = j + 1
			}

			if j > 1 {
				break
			}
		}

		if j > 1 {
			structs = append(structs[:i], structs[i+1:]...)
		}
	}

	return structs, nil
}

func generateUpdateConnectorObject(o afas.UpdateConnectorObject) (UpdateConnectorObjectStruct, error) {
	fields, err := generateUpdateConnectorObjectStructFields(o)
	if err != nil {
		return UpdateConnectorObjectStruct{}, err
	}

	objects := []string{}
	for _, o := range o.Objects {
		objects = append(objects, o.Name)
	}

	dbIDField := ""
	for _, f := range o.Fields {
		if f.PrimaryKey {
			dbIDField = normalizeIdentifier(f.Label)
		}
	}

	name := normalizeIdentifier(o.Name)
	variable := strings.ToLower(string([]rune(name)[0]))

	isSlice := false
	plural := ""
	if IsPlural(name) {
		isSlice = true
		plural = name
		name = GetSingular(name)
	}

	return UpdateConnectorObjectStruct{
		Comment:   name,
		Name:      name,
		Variable:  variable,
		Fields:    fields,
		DBIDField: dbIDField,
		Objects:   objects,
		IsSlice:   isSlice,
		Plural:    plural,
	}, nil
}

type ValidationRules struct {
	Required    bool
	NotZero     bool
	MaxLength   int
	NumDecimals int
}
