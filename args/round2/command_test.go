package round2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	schema := NewSchema(`l:bool:true p:int:0 d:str:""`)

	t.Run("test_with_all_values_shown", func(t *testing.T) {
		command := NewCommand(schema, "-l true -p 8080 -d /usr/local")
		assert.Equal(t, "true", command.GetValue("l"))
		assert.Equal(t, "8080", command.GetValue("p"))
		assert.Equal(t, "/usr/local", command.GetValue("d"))
	})
	t.Run("test_with_default_value", func(t *testing.T) {
		command := NewCommand(schema, "-l -p -d")
		assert.Equal(t, "true", command.GetValue("l"))
		assert.Equal(t, "0", command.GetValue("p"))
		assert.Equal(t, `""`, command.GetValue("d"))
	})
}
