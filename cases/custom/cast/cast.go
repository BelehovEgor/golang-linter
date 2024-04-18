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
	a, ok := someInterfaceV.(concreteType) // never panics
	print(a)
	print(ok)

}

func castCheckGoodToo() {
	var someInterfaceV someInterface
	a := someInterfaceV.(concreteType) // type assertion, can panic, check
	print(a)
}

func castCheckBad(someInterfaceV someInterface) {
	a := someInterfaceV.(concreteType2) // type assertion, can panic, check
	print(a)
}
