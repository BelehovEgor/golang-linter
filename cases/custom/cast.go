package custom

type someInterface interface {
}

type concreteType struct {
}

func castCheckGood() {
	var someInterfaceV someInterface
	a, ok := someInterfaceV.(concreteType) // never panics
	print(a)
	print(ok)

}

func castCheckBad() {
	var someInterfaceV someInterface
	a := someInterfaceV.(concreteType) // type assertion, can panic, check
	print(a)
}
