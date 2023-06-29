package tree

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

func (m *Map) GetValue(key string) string {

	if m == nil || len(key) == 0 || m.Branches == nil {
		return ""
	}

	keyRunes := []rune(key)

	var (
		targetMap *Map
		ok        bool
	)
	if targetMap, ok = m.Branches[keyRunes[0]]; !ok {
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

func (m *Map) PrependToLeaves(s string) {

	if len(m.Branches) == 0 {
		value := new(string)
		if m.Value != nil {
			value = m.Value
		}

		*value = s + *value
		return
	}

	for _, mm := range m.Branches {
		mm.PrependToLeaves(s)
	}
}
