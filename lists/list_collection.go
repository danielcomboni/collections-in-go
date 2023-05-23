package lists

import "reflect"

type List[V any] struct {
	elements []V
}

func NewList[V any]() *List[V] {
	n := new(List[V])
	s := make([]V, 0)
	n.elements = s
	return n
}

func (l *List[V]) Add(v V) *List[V] {
	l.elements = append(l.elements, v)
	return l
}

func (l *List[V]) IsEmpty() bool {
	return len(l.elements) == 0
}

func (l *List[V]) AddAll(vList []V) *List[V] {
	if !l.IsEmpty() {
		for _, v := range vList {
			l.Add(v)
		}
		return l
	}
	l.elements = vList
	return l
}

func (l *List[V]) ForEach(fn func(element V, index int)) {
	for i, v := range l.elements {
		fn(v, i)
	}
}

func (l *List[V]) Size() int64 {
	return int64(len(l.elements))
}

func (l *List[V]) Clear() *List[V] {
	l.elements = NewList[V]().elements
	return l
}

func (l *List[V]) First(condition func(v V) bool) V {
	for _, v := range l.elements {
		if condition(v) {
			return v
		}
	}
	return *new(V)
}

func (l *List[V]) Where(condition func(v V) bool) *List[V] {
	n := new(List[V])
	for _, v := range l.elements {
		if condition(v) {
			n.Add(v)
		}
	}
	return n
}

func (l *List[V]) IndexOf(element V) int {
	for i, v := range l.elements {
		if reflect.DeepEqual(element, v) {
			return i
		}
	}
	return -1
}

func (l *List[V]) Replace(oldValue V, newValue V) *List[V] {
	index := l.IndexOf(oldValue)
	if index <= 0 {
		l.elements[index] = newValue
	}
	return l
}

func (l *List[V]) Remove(index int) *List[V] {
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	return l
}
