package tree

type Map struct {
	value    *string
	branches map[rune]*Map
}

// PutValue inserts a key/value pair into the tree, each of the key's runes
// representing a branch in the tree.
func (m *Map) PutValue(key []rune, value string) {

	if m == nil || len(key) == 0 {
		return
	}

	if m.branches == nil {
		m.branches = make(map[rune]*Map)
	}

	targetMap := new(Map)
	if existingMap, ok := m.branches[key[0]]; ok {
		targetMap = existingMap
	}

	if len(key) == 1 {
		targetMap.value = &value
	} else {
		targetMap.PutValue(key[1:], value)
	}

	m.branches[key[0]] = targetMap
}

func (m *Map) PutMap(key []rune, newMap *Map) {

	if m == nil || newMap == nil || len(key) == 0 {
		return
	}

	if m.branches == nil {
		m.branches = make(map[rune]*Map)
	}

	if len(key) == 1 {
		m.branches[key[0]] = newMap
		return
	}

	m.branches[key[0]].PutMap(key[1:], newMap)
}

func (m *Map) GetValue(key string) string {

	if m == nil || len(key) == 0 || m.branches == nil {
		return ""
	}

	keyRunes := []rune(key)

	var (
		targetMap *Map
		ok        bool
	)
	if targetMap, ok = m.branches[keyRunes[0]]; !ok {
		return ""
	}

	if len(key) == 1 {
		if targetMap.value != nil {
			return *targetMap.value
		}

		return ""
	}

	return targetMap.GetValue(key[1:])
}

func (m *Map) GetMap(key []rune) *Map {

	if m == nil || len(key) == 0 || m.branches == nil {
		return nil
	}

	var (
		targetMap *Map
		ok        bool
	)
	if targetMap, ok = m.branches[key[0]]; !ok {
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
	for k, v := range n.branches {
		branches[k] = v.Copy()
	}

	return &Map{
		value:    n.value,
		branches: branches,
	}
}
