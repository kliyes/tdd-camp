package round1

import "strings"

type Command struct {
	cmd  string
	args map[string]string
}

func NewCommand(cmd string) Command {
	command := Command{cmd: cmd, args: make(map[string]string)}
	command.parse()
	return command
}

// -l true -p 8080 -d /usr/local
func (c *Command) parse() {
	items := strings.Fields(c.cmd)
	for n, i := range items {
		if strings.HasPrefix(i, "-") {
			label := strings.TrimPrefix(i, "-")
			if n == len(items)-1 {
				c.args[label] = ""
			} else {
				nextItem := items[n+1]
				if strings.HasPrefix(nextItem, "-") {
					nextItem = ""
				}
				c.args[label] = nextItem
			}
		}
	}
}

func (c *Command) GetValue(label string) string {
	return c.args[label]
}
