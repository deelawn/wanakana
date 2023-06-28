package wanakana

// TODO: Make a default CustomMapping implementation that is a string -> string map.
type CustomMapping interface {
	Apply(TreeMapper)
}

type TreeMapper interface {
	Map(key, value string)
}
