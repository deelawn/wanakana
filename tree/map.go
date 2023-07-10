package tree

// Map is a tree structure that maps sequences of runes to strings.
type Map struct {
	Value    *string
	Branches map[rune]*Map
}

// Map is an alias for PutValue and exists to satisfy an interface requirement in the public API.
func (m *Map) Map(key string, value string) {
	m.PutValue([]rune(key), value)
}

// PutValue inserts a key/value pair into the tree, each of the key's runes
// representing a branch in the tree.
func (m *Map) PutValue(key []rune, value string) {

	if m == nil || len(key) == 0 {
		return
	}

	if m.Branches == nil {
		m.Branches = make(map[rune]*Map)
	}

	targetMap := new(Map)
	if existingMap, ok := m.Branches[key[0]]; ok {
		targetMap = existingMap
	}

	if len(key) == 1 {
		targetMap.Value = &value
	} else {
		targetMap.PutValue(key[1:], value)
	}

	m.Branches[key[0]] = targetMap
}

// PutMap adds a new branch to the current map witht the given key/value pair.
func (m *Map) PutMap(key []rune, newMap *Map) {

	if m == nil || newMap == nil || len(key) == 0 {
		return
	}

	if m.Branches == nil {
		m.Branches = make(map[rune]*Map)
	}

	if len(key) == 1 {
		m.Branches[key[0]] = newMap
		return
	}

	m.Branches[key[0]].PutMap(key[1:], newMap)
}

// GetValue returns the value associated with the given key, or an empty string if the key is not found.
func (m *Map) GetValue(key []rune) string {

	if m == nil || len(key) == 0 || m.Branches == nil {
		return ""
	}

	var (
		targetMap *Map
		ok        bool
	)
	if targetMap, ok = m.Branches[key[0]]; !ok {
		return ""
	}

	if len(key) == 1 {
		if targetMap.Value != nil {
			return *targetMap.Value
		}

		return ""
	}

	return targetMap.GetValue(key[1:])
}

// GetMap returns the map associated with the given key, or nil if the key is not found.
func (m *Map) GetMap(key []rune) *Map {

	if m == nil || len(key) == 0 || m.Branches == nil {
		return nil
	}

	var (
		targetMap *Map
		ok        bool
	)
	if targetMap, ok = m.Branches[key[0]]; !ok {
		return nil
	}

	if len(key) == 1 {
		return targetMap
	}

	return targetMap.GetMap(key[1:])
}

// Copy returns a deep copy of the current map.
func (n *Map) Copy() *Map {

	if n == nil {
		return nil
	}

	branches := make(map[rune]*Map)
	for k, v := range n.Branches {
		branches[k] = v.Copy()
	}

	return &Map{
		Value:    n.Value,
		Branches: branches,
	}
}
