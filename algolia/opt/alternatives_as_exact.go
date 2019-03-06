// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AlternativesAsExactOption struct {
	value []string
}

func AlternativesAsExact(v ...string) *AlternativesAsExactOption {
	return &AlternativesAsExactOption{v}
}

func (o AlternativesAsExactOption) Get() []string {
	return o.value
}

func (o AlternativesAsExactOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AlternativesAsExactOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{"ignorePlurals", "singleWordSynonym"}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
