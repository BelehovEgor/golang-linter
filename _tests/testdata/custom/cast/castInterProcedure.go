package cast

func callBadCast() {
	var someInterfaceV someInterface
	castCheckBad(someInterfaceV)
}
