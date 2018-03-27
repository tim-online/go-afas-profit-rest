package afas

import "encoding/json"

type Elements []Element

func (e Elements) MarshalJSON() ([]byte, error) {
	type alias Element

	st := struct {
		Elements []alias `json:"Element"`
	}{}
	for _, e := range e {
		st.Elements = append(st.Elements, alias(e))
	}
	return json.Marshal(st)
}

type Element struct {
	DBID    string                 `json:"@DbId,omitempty"`
	Fields  map[string]interface{} `json:"Fields,omitempty"`
	Objects map[string]interface{} `json:"Objects,omitempty"`
}

func (e Element) MarshalJSON() ([]byte, error) {
	type alias Element
	a := alias(e)

	st := struct {
		alias `json:"Element"`
	}{a}
	return json.Marshal(st)
}
