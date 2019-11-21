package round2

import "strings"

type Meta struct {
	valueType    string
	valueDefault string
}

func NewMeta(t, d string) Meta {
	return Meta{t, d}
}

type Schema struct {
	typeMap map[string]Meta
}

func NewSchema(text string) Schema {
	schema := Schema{typeMap: make(map[string]Meta)}
	schema.parse(text)
	return schema
}

func (s *Schema) parse(text string) {
	// l:bool:true p:int:0 d:str:""
	labelTypes := strings.Fields(text)
	for _, lt := range labelTypes {
		items := strings.Split(lt, ":")
		s.typeMap[items[0]] = NewMeta(items[1], items[2])
	}
}

func (s *Schema) GetType(label string) string {
	t, _ := s.typeMap[label]
	return t.valueType
}

func (s *Schema) GetLabels() []string {
	labels := []string{}
	for k := range s.typeMap {
		labels = append(labels, k)
	}
	return labels
}

func (s *Schema) GetDefaultValue(label string) string {
	t, _ := s.typeMap[label]
	return t.valueDefault
}
