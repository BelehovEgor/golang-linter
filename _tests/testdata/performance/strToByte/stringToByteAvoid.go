package strToByte

import "os"

func badExampleStringToByte(w os.File) {
	for i := 0; i < 10; i++ {
		_, err := w.Write([]byte("Hello world"))
		if err != nil {
			return
		}
	}
}

func goodExampleStringToByte(w os.File) {
	data := []byte("Hello world")
	for i := 0; i < 10; i++ {
		_, err := w.Write(data)
		if err != nil {
			return
		}
	}
}
