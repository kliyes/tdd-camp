package round2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	schema := NewSchema(`l:bool:true p:int:0 d:str:""`)

	t.Run("test_command_with_all_values", func(t *testing.T) {
		command := NewCommand(schema, "-l true -p 8080 -d /usr/local")
		args := NewArgs(schema, command)
		assert.Equal(t, true, args.GetValue("l"))
		assert.Equal(t, 8080, args.GetValue("p"))
		assert.Equal(t, "/usr/local", args.GetValue("d"))
	})
	t.Run("test_command_with_default_value", func(t *testing.T) {
		command := NewCommand(schema, "-l -p -d")
		args := NewArgs(schema, command)
		assert.Equal(t, true, args.GetValue("l"))
		assert.Equal(t, 0, args.GetValue("p"))
		assert.Equal(t, `""`, args.GetValue("d"))
	})
	t.Run("test_command_with_wrong_type", func(t *testing.T) {
		command := NewCommand(schema, "-l x -p x")
		args := NewArgs(schema, command)
		assert.Equal(t, true, args.GetValue("l"))
		assert.Equal(t, 0, args.GetValue("p"))
	})
}
