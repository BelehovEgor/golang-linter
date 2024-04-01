package main

func main() {

}

func wrong_name_style(b bool) int {
	var a int
	var c = 2
	var d = 3
	area := Area(2.1)
	vol := Volume(3.1)
	if b {
		area += area
		vol += vol
		c++
		a = 100
	} else {
		d--
		a = 10
	}
	return a
}

type Area float64
type Volume float64
