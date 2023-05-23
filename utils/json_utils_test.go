package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Id   int64
	Name string
	Age  int64
}

func TestSafeGetFromInterface(t *testing.T) {
	u := User{
		Id:   1,
		Name: "Daniel",
		Age:  90,
	}
	d := SafeGetFromInterface(u, "$.id")
	println(fmt.Sprintf("dua: %v", d))
}

type TypeWrapper interface {
	Number | interface{}
}

func Sum[T Number](numbers []T) T {
	var total T
	println(reflect.TypeOf(new(T)).Name())
	for _, x := range numbers {
		total += x
	}
	return total
}

func TestWrapper(t *testing.T) {

	//s := Sum[int64]([]int64{
	//	1, 2, 3, 4,
	//})
	//println(s)
}
