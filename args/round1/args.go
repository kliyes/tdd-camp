package round1

type Args struct {
	schema  Schema
	command Command
}

func NewArgs(schema Schema, command Command) Args {
	return Args{schema, command}
}

func (a Args) GetValue(label string) interface{} {
	return a.schema.GetValue(label, a.command.GetValue(label))
}
