package maps

import (
	"fmt"
	"github.com/danielcomboni/collections-in-go/utils"
)

type Map[K int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool, V any] struct {
	elements map[K]V
}

func NewMap[K int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool, V any]() *Map[K, V] {
	n := new(Map[K, V])
	m := make(map[K]V)
	n.elements = m
	return n
}

func (m *Map[K, V]) Put(k K, v V) *Map[K, V] {
	m.elements[k] = v
	return m
}

func (m *Map[K, V]) PutAllFromMap(vMap map[K]V) *Map[K, V] {
	m.elements = vMap
	return m
}

func (m *Map[K, V]) PutAllFromSlice(vSlice []V, keyProperty string) *Map[K, V] {
	for _, v := range vSlice {
		key := utils.SafeGetFromInterface(v, fmt.Sprintf("$.%v", keyProperty))
		m.elements[key.(K)] = v
	}
	return m
}

func (m *Map[K, V]) Size() int64 {
	return int64(len(m.elements))
}

func (m *Map[K, V]) IsEmpty() bool {
	return len(m.elements) == 0
}

func (m *Map[K, V]) Contains(k K) bool {
	return !utils.IsNullOrEmpty(m.elements[k])
}

func (m *Map[K, V]) Clear() *Map[K, V] {
	m.elements = make(map[K]V)
	return m
}

func (m *Map[K, V]) ForEach(fn func(element V, index K)) {
	for i, v := range m.elements {
		fn(v, i)
	}
}

func (m *Map[K, V]) First(condition func(v V) bool) V {
	for _, v := range m.elements {
		if condition(v) {
			return v
		}
	}
	return *new(V)
}

func (m *Map[K, V]) Where(condition func(v V) bool) *Map[K, V] {
	n := new(Map[K, V])
	for k, v := range m.elements {
		if condition(v) {
			n.Put(k, v)
		}
	}
	return n
}

func (m *Map[K, V]) GetKeys() []K {
	keys := make([]K, 0, len(m.elements))

	for k := range m.elements {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) GetValues() []V {
	values := make([]V, 0, len(m.elements))
	for _, v := range m.elements {
		values = append(values, v)
	}
	return values
}

func (m *Map[K, V]) Replace(oldKey K, newValue V) V {
	m.elements[oldKey] = newValue
	return m.elements[oldKey]
}
