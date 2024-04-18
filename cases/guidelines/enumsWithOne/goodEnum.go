package enumsWithOne

const (
	Add1 Operation = iota + 1
	Subtract1
	Multiply1
)

func printGoodEnum() {
	print(Add1, Subtract1, Multiply1)
}
