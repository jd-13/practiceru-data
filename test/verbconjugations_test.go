package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerbConjugations(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/verbconjugations/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var data []interface{}
	err = json.Unmarshal(content, &data)
	require.NoError(t, err)

	// Validate each item
	for _, item := range data {
		verb, ok := item.(map[string]interface{})
		require.True(t, ok)

		// Validate imperfective
		require.Contains(t, verb, "imperfective")
		imperfective, ok := verb["imperfective"].(map[string]interface{})
		require.True(t, ok)

		assert.Contains(t, imperfective, "infinitive")
		assert.NotEmpty(t, imperfective["infinitive"])

		require.Contains(t, imperfective, "past")
		past, ok := imperfective["past"].(map[string]interface{})
		require.True(t, ok)

		for _, key := range []string{"он", "она", "оно", "они"} {
			assert.Contains(t, past, key)
			assert.NotEmpty(t, past[key])
			_, ok := past[key].(string)
			assert.True(t, ok)
		}

		require.Contains(t, imperfective, "present")
		present, ok := imperfective["present"].(map[string]interface{})
		require.True(t, ok)

		for _, key := range []string{"я", "ты", "он", "она", "оно", "мы", "вы", "они"} {
			assert.Contains(t, present, key)
			assert.NotEmpty(t, present[key])
			_, ok := present[key].(string)
			assert.True(t, ok)
		}

		// Validate perfective
		require.Contains(t, verb, "perfective")
		perfective, ok := verb["perfective"].(map[string]interface{})
		require.True(t, ok)

		assert.Contains(t, perfective, "infinitive")
		assert.NotEmpty(t, perfective["infinitive"])

		require.Contains(t, perfective, "past")
		past, ok = perfective["past"].(map[string]interface{})
		require.True(t, ok)

		for _, key := range []string{"он", "она", "оно", "они"} {
			assert.Contains(t, past, key)
			assert.NotEmpty(t, past[key])
			_, ok := past[key].(string)
			assert.True(t, ok)
		}

		require.Contains(t, perfective, "future")
		future, ok := perfective["future"].(map[string]interface{})
		require.True(t, ok)

		for _, key := range []string{"я", "ты", "он", "она", "оно", "мы", "вы", "они"} {
			assert.Contains(t, future, key)
			assert.NotEmpty(t, future[key])
			_, ok := future[key].(string)
			assert.True(t, ok)
		}
	}
}