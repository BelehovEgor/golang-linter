package custom

import "fmt"

type NullInterface interface {
	Method()
}

func TestNullInterfaceBad() {
	var a NullInterface
	a.Method() // oops, panic
}

func TestNullInterfaceGood() {
	var a NullInterface
	a = S{}
	a.Method() // niiice
}

// S will satisfy interfaces NullInterface
type S struct{}

func (s S) Method() {
	fmt.Println("A")
}
