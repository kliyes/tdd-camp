package round1

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	schema := NewSchema(configMap{
		"l": reflect.Bool, "d": reflect.String, "p": reflect.Int,
	})
	t.Run("test", func(t *testing.T) {
		command := NewCommand("-l true -d /usr/local/ -p 8080")
		args := NewArgs(schema, command)
		assert.Equal(t, true, args.GetValue("l"))
		assert.Equal(t, 8080, args.GetValue("p"))
		assert.Equal(t, "/usr/local/", args.GetValue("d"))
	})
	t.Run("test_has_no_value", func(t *testing.T) {
		command := NewCommand("-l -d /usr/local/ -p")
		args := NewArgs(schema, command)
		assert.Equal(t, true, args.GetValue("l"))
		assert.Equal(t, 0, args.GetValue("p"))
	})
}
