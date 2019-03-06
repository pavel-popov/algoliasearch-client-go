// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestResponseFields(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ResponseFieldsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.ResponseFields([]string{"*"}...),
		},
		{
			opts:     []interface{}{opt.ResponseFields("value1")},
			expected: opt.ResponseFields("value1"),
		},
		{
			opts:     []interface{}{opt.ResponseFields("value1", "value2", "value3")},
			expected: opt.ResponseFields("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractResponseFields(c.opts...)
			out opt.ResponseFieldsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
