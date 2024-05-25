package strToByte

import "os"

func badExampleStringToByte(w os.File) {
	for i := 0; i < 10; i++ {
		_, err := w.Write([]byte("Hello world")) // want "Do not create byte slices from a fixed string repeatedly. Instead, perform the conversion once and capture the result."
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
