package avoidMutableGlobals

var _someData = 123

func signEx(msg string) string {
	if _someData > 10 {
		return msg + "123"
	}

	return msg
}
