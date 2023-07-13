package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImperative(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/imperative/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var verbs []interface{}
	err = json.Unmarshal(content, &verbs)
	require.NoError(t, err)

	// Validate each item
	for _, item := range verbs {
		verb, ok := item.(map[string]interface{})
		require.True(t, ok)

		assert.Contains(t, verb, "infinitive")
		infinitive, ok := verb["infinitive"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, infinitive)

		assert.Contains(t, verb, "singular")
		singular, ok := verb["singular"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, singular)

		assert.Contains(t, verb, "plural")
		plural, ok := verb["plural"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, plural)

		if _, ok = verb["rule"]; ok {
			ruleNumber, ok := verb["rule"].(float64)
			assert.True(t, ok)
			assert.GreaterOrEqual(t, ruleNumber, 0.0)
		}
	}
}