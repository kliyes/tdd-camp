package round1

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	t.Run("test_bool", func(t *testing.T) {
		schema := Schema{configMap{"l": reflect.Bool}}
		assert.Equal(t, false, schema.GetValue("l", "false"))
		assert.Equal(t, true, schema.GetValue("l", "true"))
		assert.Equal(t, true, schema.GetValue("l", ""))
	})
	t.Run("test_int", func(t *testing.T) {
		schema := Schema{configMap{"p": reflect.Int}}
		assert.Equal(t, 8080, schema.GetValue("p", "8080"))
		assert.Equal(t, 0, schema.GetValue("p", ""))
	})
	t.Run("test_string", func(t *testing.T) {
		schema := Schema{configMap{"d": reflect.String}}
		assert.Equal(t, "/usr/local", schema.GetValue("d", "/usr/local"))
		assert.Equal(t, "", schema.GetValue("d", ""))
	})
}
