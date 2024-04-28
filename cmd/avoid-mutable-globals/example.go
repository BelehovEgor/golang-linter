package main

var _someData = 123

func sign(msg string) string {
	if _someData > 10 {
		return msg + "123"
	}

	return msg
}
