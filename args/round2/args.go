package round2

import "strconv"

type Args struct {
	schema   Schema
	command  Command
	valueMap map[string]interface{}
}

func NewArgs(schema Schema, command Command) Args {
	args := Args{schema, command, make(map[string]interface{})}
	args.parse()
	return args
}

func (a *Args) parse() {
	labels := a.schema.GetLabels()
	for _, label := range labels {
		rawValue := a.command.GetValue(label)
		var (
			value interface{}
			err   error
		)
		t := a.schema.GetType(label)
		if t == "bool" {
			value, err = strconv.ParseBool(rawValue)
			if err != nil {
				value, _ = strconv.ParseBool(a.schema.GetDefaultValue(label))
			}
		} else if t == "int" {
			value, err = strconv.Atoi(rawValue)
			if err != nil {
				value, _ = strconv.Atoi(a.schema.GetDefaultValue(label))
			}
		} else if t == "str" {
			value = rawValue
		}
		a.valueMap[label] = value
	}
}

func (a *Args) GetValue(label string) interface{} {
	return a.valueMap[label]
}
