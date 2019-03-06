// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestUnretrievableAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.UnretrievableAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.UnretrievableAttributes([]string{}...),
		},
		{
			opts:     []interface{}{opt.UnretrievableAttributes("value1")},
			expected: opt.UnretrievableAttributes("value1"),
		},
		{
			opts:     []interface{}{opt.UnretrievableAttributes("value1", "value2", "value3")},
			expected: opt.UnretrievableAttributes("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractUnretrievableAttributes(c.opts...)
			out opt.UnretrievableAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
