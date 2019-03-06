// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type FacetsOption struct {
	value []string
}

func Facets(v ...string) *FacetsOption {
	return &FacetsOption{v}
}

func (o FacetsOption) Get() []string {
	return o.value
}

func (o FacetsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *FacetsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
