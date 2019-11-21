package round2

import "strings"

type Command struct {
	schema   Schema
	valueMap map[string]string
}

func NewCommand(schema Schema, cmd string) Command {
	command := Command{schema, make(map[string]string)}
	command.parse(cmd)
	return command
}

// -l true -p 8080 -d /usr/local
func (c *Command) parse(cmd string) {
	labels := c.schema.GetLabels()
	for _, label := range labels {
		var labelValues []string
		values := strings.SplitN(cmd, "-"+label, 2)
		if len(values) > 1 {
			labelValues = strings.Fields(values[1])
		}

		var value string
		if len(labelValues) > 0 {
			value = labelValues[0]
		}
		if value == "" || strings.HasPrefix(value, "-") {
			value = c.schema.GetDefaultValue(label)
		}
		c.valueMap[label] = value
	}
}

func (c *Command) GetValue(label string) string {
	return c.valueMap[label]
}
