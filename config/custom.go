package config

// TODO: Make a default CustomMapping implementation that is a string -> string map.
type CustomMapping interface {
	Apply(TreeMapper)
}

// TreeMapper is an interface that allows CustomMapping implementations to apply their mappings to a tree map.
type TreeMapper interface {
	Map(key, value string)
}

// CustomMappingKeyValue is a CustomMapping implementation that maps a string -> string map to a tree map.
type CustomMappingKeyValue struct {
	m map[string]string
}

// NewCustomMappingKeyValue returns a new CustomMappingKeyValue instance initialized with the given map.
func NewCustomMappingKeyValue(m map[string]string) *CustomMappingKeyValue {
	return &CustomMappingKeyValue{m: m}
}

// Apply applies the CustomMappingKeyValue's map to the given TreeMapper.
func (m *CustomMappingKeyValue) Apply(tm TreeMapper) {
	for k, v := range m.m {
		tm.Map(k, v)
	}
}
