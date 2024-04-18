package cast

type someInterface interface {
	someInterfaceMethod()
}

type concreteType struct {
}

type concreteType2 struct {
}

func (receiver concreteType) someInterfaceMethod() {
	print("ya method")
}

func castCheckGood() {
	var someInterfaceV someInterface
	_, ok := someInterfaceV.(concreteType) // never panics
	print(ok)
}

func castCheckGoodToo() {
	var someInterfaceV someInterface
	c := someInterfaceV.(concreteType) // type assertion, can panic, check
	c.someInterfaceMethod()
}

func castCheckBad(someInterfaceV someInterface) {
	//does not compile
	//a := someInterfaceV.(concreteType2) // type assertion, can panic, check
	//print(a)
}
