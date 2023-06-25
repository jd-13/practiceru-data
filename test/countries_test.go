package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountries(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/countries/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var countries []interface{}
	err = json.Unmarshal(content, &countries)
	require.NoError(t, err)

	// Validate each item
	for _, item := range countries {
		country, ok := item.(map[string]interface{})
		require.True(t, ok)

		assert.Contains(t, country, "flag")
		_, ok = country["flag"].(string)
		assert.True(t, ok)

		assert.Contains(t, country, "country")
		_, ok = country["country"].(string)
		assert.True(t, ok)

		assert.Contains(t, country, "genitive")
		_, ok = country["genitive"].(string)
		assert.True(t, ok)

		assert.Contains(t, country, "languages")
		languages, ok := country["languages"].([]interface{})
		assert.True(t, ok)

		for _, language := range languages {
			_, ok = language.(string)
			assert.True(t, ok)
		}

		assert.Contains(t, country, "nationality")
		nationality, ok := country["nationality"].(map[string]interface{})
		require.True(t, ok)

		assert.Contains(t, nationality, "masculine")
		_, ok = nationality["masculine"].(string)
		assert.True(t, ok)

		assert.Contains(t, nationality, "feminine")
		_, ok = nationality["feminine"].(string)
		assert.True(t, ok)

		assert.Contains(t, nationality, "plural")
		_, ok = nationality["plural"].(string)
		assert.True(t, ok)
	}
}