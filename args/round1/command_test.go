package round1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	t.Run("test_has_value", func(t *testing.T) {
		command := NewCommand("-l true -p 8080 -d /usr/local")
		assert.Equal(t, "true", command.GetValue("l"))
		assert.Equal(t, "8080", command.GetValue("p"))
		assert.Equal(t, "/usr/local", command.GetValue("d"))
	})
	t.Run("test_has_no_value", func(t *testing.T) {
		command := NewCommand("-l -p 8080 -d")
		assert.Equal(t, "", command.GetValue("l"))
		assert.Equal(t, "", command.GetValue("d"))
	})
}
