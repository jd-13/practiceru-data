package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNumbers(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/numbers/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var numbers map[string]interface{}
	err = json.Unmarshal(content, &numbers)
	require.NoError(t, err)

	// Validate cardinal
	assert.Contains(t, numbers, "cardinal")
	for k, v := range numbers["cardinal"].(map[string]interface{}) {
		assert.IsType(t, "string", k)

		if value, ok := v.(string); ok {
			assert.IsType(t, "string", value)
		} else if values, ok := v.([]interface{}); ok {
			assert.Len(t, values, 3)
			for _, value := range values {
				assert.IsType(t, "string", value)
			}
		} else {
			t.Errorf("Invalid type for key %s", k)
		}
	}

	// Validate ordinal
	assert.Contains(t, numbers, "ordinal")
	for k, v := range numbers["ordinal"].(map[string]interface{}) {
		assert.IsType(t, "string", k)

		if values, ok := v.([]interface{}); ok {
			assert.Len(t, values, 4)
			for _, value := range values {
				assert.IsType(t, "string", value)
			}
		} else {
			t.Errorf("Invalid type for key %s", k)
		}
	}
}