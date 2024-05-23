package handleTypeAssertionFailures

func someHandlingBad(i any) {
	t := i.(string)
	print(t)
}
