package wanakana

// TODO: Make a default CustomMapping implementation that is a string -> string map.
type CustomMapping interface {
	Apply(TreeMapper)
}

type TreeMapper interface {
	Map(key, value string)
}

type CustomMappingKeyValue struct {
	m map[string]string
}

func NewCustomMappingKeyValue(m map[string]string) *CustomMappingKeyValue {
	return &CustomMappingKeyValue{m: m}
}

func (m *CustomMappingKeyValue) Apply(tm TreeMapper) {
	for k, v := range m.m {
		tm.Map(k, v)
	}
}
