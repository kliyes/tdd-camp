package round2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSchema(t *testing.T) {
	schemaText := `l:bool:true p:int:0 d:str:""`
	schema := NewSchema(schemaText)

	t.Run("test_get_type", func(t *testing.T) {
		assert.Equal(t, "bool", schema.GetType("l"))
		assert.Equal(t, "int", schema.GetType("p"))
		assert.Equal(t, "str", schema.GetType("d"))
	})
	t.Run("test_get_default_value", func(t *testing.T) {
		assert.Equal(t, "true", schema.GetDefaultValue("l"))
		assert.Equal(t, "0", schema.GetDefaultValue("p"))
		assert.Equal(t, `""`, schema.GetDefaultValue("d"))
	})
	t.Run("test_get_labels", func(t *testing.T) {
		got := schema.GetLabels()
		assert.Len(t, got, 3)
		assert.Contains(t, got, "l")
		assert.Contains(t, got, "p")
		assert.Contains(t, got, "d")
	})
}
