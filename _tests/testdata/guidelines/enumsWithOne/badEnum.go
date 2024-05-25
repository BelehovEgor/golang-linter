package enumsWithOne

type Operation int

const ( // want "In a constant 'Operation', the starting value 'Add' is equal 0."
	Add Operation = iota
	Subtract
	Multiply
)

func printBadEnum() {
	print(Add, Subtract, Multiply)
}
