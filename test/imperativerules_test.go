package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImperativeRules(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/imperativerules/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var rules map[string]interface{}
	err = json.Unmarshal(content, &rules)
	require.NoError(t, err)

	// Validate singular rules
	assert.Contains(t, rules, "singular")
	singularRules, ok := rules["singular"].([]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, singularRules)

	for _, item := range singularRules {
		rule, ok := item.(string)
		assert.True(t, ok)
		assert.NotEmpty(t, rule)
	}

	// Validate plural rules
	assert.Contains(t, rules, "plural")
	pluralRules, ok := rules["plural"].([]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, pluralRules)

	for _, item := range pluralRules {
		rule, ok := item.(string)
		assert.True(t, ok)
		assert.NotEmpty(t, rule)
	}
}