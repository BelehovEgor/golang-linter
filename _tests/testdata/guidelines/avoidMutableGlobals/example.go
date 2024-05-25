package avoidMutableGlobals // want "Don't use mutable global variable '_someData', use dependency injection."

var _someData = 123

func signEx(msg string) string {
	if _someData > 10 {
		return msg + "123"
	}

	return msg
}
