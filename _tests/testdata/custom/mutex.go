package custom

import (
	"fmt"
	"reflect"
	"sync"
)

type A struct {
	a   int `protectedby:"mtx"`
	mtx sync.Mutex
}

func TestMutexGood() {
	var s = A{a: 1}
	s.mtx.Lock()
	s.a += 1
	s.mtx.Unlock()
	fmt.Println(getFieldTag(s, "a", "protectedby")) //name_field
}

func TestMutexBad() {
	var s = A{a: 1}
	s.a += 1
	fmt.Println(getFieldTag(s, "a", "protectedby")) //name_field
}

func getFieldTag(s A, fieldName string, tagName string) string {
	field, ok := reflect.TypeOf(&s).Elem().FieldByName(fieldName)
	if !ok {
		panic("Field not found")
	}
	return field.Tag.Get(tagName)
}
