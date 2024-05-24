package enumsWithOne

type Operation int

const (
	Add Operation = iota
	Subtract
	Multiply
)

func printBadEnum() {
	print(Add, Subtract, Multiply)
}
