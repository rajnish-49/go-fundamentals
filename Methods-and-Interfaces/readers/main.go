package main

//Imagine you are reading a book. You don't read the whole book at once — you read page by page.
// io.Reader is the same idea. It reads data chunk by chunk, not all at once.

import (
	"fmt"
	"io"
	"strings"
)

// the io.Reader interface has a single method, Read 
// which takes a byte slice and returns the number of bytes read and an error if any. 
// The Read method reads data into the provided byte slice and returns the number of bytes read and any error encountered. 
// When the end of the data is reached, it returns io.EOF.

type Reader interface {
    Read(b []byte) (int, error)
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) // the bucket , the size of the chunk we want to read 
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
