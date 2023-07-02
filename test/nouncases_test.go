package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNounCases(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/nouncases/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var nouns map[string]interface{}
	err = json.Unmarshal(content, &nouns)
	require.NoError(t, err)

	cases := []string{"nominative", "genitive", "accusative", "dative", "instrumental", "prepositional"}

	for _, noun := range nouns {
		for _, singularPlural := range []string{"singular", "plural"} {
			assert.Contains(t, noun, singularPlural)
			nounMap, ok := noun.(map[string]interface{})
			require.True(t, ok)

			quantityMap, ok := nounMap[singularPlural].(map[string]interface{})
			require.True(t, ok)

			for _, caseName := range cases {
				assert.Contains(t, quantityMap, caseName)
				caseMap, ok := quantityMap[caseName].(map[string]interface{})
				require.True(t, ok)

				// Must always have text
				assert.Contains(t, caseMap, "text")
				text, ok := caseMap["text"].(string)
				require.True(t, ok)
				assert.NotEmpty(t, text)

				// Singular nominative doesn't have a case rule
				if !(singularPlural == "singular" && caseName == "nominative") {
					assert.Contains(t, caseMap, "caseRule")
					ruleNumber, ok := caseMap["caseRule"].(float64)
					assert.True(t, ok)
					assert.GreaterOrEqual(t, ruleNumber, -1.0)
				}

				if _, ok = caseMap["spellingRule"]; ok {
					ruleNumber, ok := caseMap["spellingRule"].(float64)
					assert.True(t, ok)
					assert.GreaterOrEqual(t, ruleNumber, 0.0)
				}
			}
		}
	}
}
