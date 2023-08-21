package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColours(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/colours/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	require.NoError(t, err)

	// Validate the colours
	colours, ok := data["colours"].(map[string]interface{})
	require.True(t, ok)

	for key, value := range colours {
		assert.Equal(t, 6, len(key))

		declensions, ok := value.(map[string]interface{})
		require.True(t, ok)

		for _, gender := range []string{"masculine", "feminine", "neuter", "plural"} {
			assert.Contains(t, value, gender)
			declension, ok := declensions[gender].(string)
			assert.True(t, ok)
			assert.NotEmpty(t, declension)
		}
	}

	// Validate the nouns
	nouns, ok := data["nouns"].(map[string]interface{})
	require.True(t, ok)

	for key, value := range nouns {
		assert.NotEmpty(t, key)

		valueStr, ok := value.(string)
		assert.True(t, ok)
		assert.NotEmpty(t, valueStr)
	}
}