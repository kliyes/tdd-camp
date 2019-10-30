package round1

import (
	"reflect"
	"strconv"
)

type configMap map[string]reflect.Kind

type Schema struct {
	configs configMap
}

func NewSchema(cfg configMap) Schema {
	return Schema{cfg}
}

func (s Schema) GetValue(label, strValue string) interface{} {
	kind, exist := s.configs[label]
	if !exist {
		return nil
	}

	switch kind {
	case reflect.Bool:
		if strValue == "" {
			return true
		}
		value, err := strconv.ParseBool(strValue)
		if err != nil {
			return nil
		}
		return value
	case reflect.Int:
		if strValue == "" {
			return 0
		}
		value, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return nil
		}
		return int(value)
	default:
		return strValue
	}
}
