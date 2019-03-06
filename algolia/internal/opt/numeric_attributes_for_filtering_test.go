// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestNumericAttributesForFiltering(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.NumericAttributesForFilteringOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.NumericAttributesForFiltering([]string{}...),
		},
		{
			opts:     []interface{}{opt.NumericAttributesForFiltering("value1")},
			expected: opt.NumericAttributesForFiltering("value1"),
		},
		{
			opts:     []interface{}{opt.NumericAttributesForFiltering("value1", "value2", "value3")},
			expected: opt.NumericAttributesForFiltering("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractNumericAttributesForFiltering(c.opts...)
			out opt.NumericAttributesForFilteringOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
