package handleTypeAssertionFailures

func someHandlingGood(i any) {
	t, ok := i.(string)
	if !ok {
	} else {
		print(t)
	}
}
