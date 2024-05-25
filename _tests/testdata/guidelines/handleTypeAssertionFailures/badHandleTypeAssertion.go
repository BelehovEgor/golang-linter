package handleTypeAssertionFailures

func someHandlingBad(i any) {
	t := i.(string) // want "The single return value form of a type assertion will panic on an incorrect type. Therefore, always use the \"comma ok\" idiom."
	print(t)
}
