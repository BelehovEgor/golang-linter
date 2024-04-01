package main

func main() {

}

func wrong_name_style(b bool) int {
	var a int
	var c = 2
	var d = 3
	if b {
		c++
		a = 100
	} else {
		d--
		a = 10
	}

	return a
}
