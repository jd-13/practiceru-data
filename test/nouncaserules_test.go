package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNounCaseRules(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/nouncaserules/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var rules map[string]interface{}
	err = json.Unmarshal(content, &rules)
	require.NoError(t, err)

	cases := []string{"genitive", "accusative", "dative", "instrumental", "prepositional"}

	// Validate singular rules
	assert.Contains(t, rules, "singular rules")
	singularRules, ok := rules["singular rules"].(map[string]interface{})
	require.True(t, ok)

	for _, caseName := range cases {
		assert.Contains(t, singularRules, caseName)
		caseRules, ok := singularRules[caseName].([]interface{})
		require.True(t, ok)

		for _, rule := range caseRules {
			ruleString, ok := rule.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, ruleString)
		}
	}

	// Validate plural rules
	assert.Contains(t, rules, "plural rules")
	pluralRules, ok := rules["plural rules"].(map[string]interface{})
	require.True(t, ok)

	for _, caseName := range cases {
		assert.Contains(t, pluralRules, caseName)
		caseRules, ok := pluralRules[caseName].([]interface{})
		require.True(t, ok)

		for _, rule := range caseRules {
			ruleString, ok := rule.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, ruleString)
		}
	}

	// Validate spelling rules
	assert.Contains(t, rules, "spelling rules")
	spellingRules, ok := rules["spelling rules"].([]interface{})
	require.True(t, ok)

	for _, rule := range spellingRules {
		ruleString, ok := rule.(string)
		assert.True(t, ok)
		assert.NotEmpty(t, ruleString)
	}
}
