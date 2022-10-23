package input

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReturnListHangmanStates() []string {
	file, err := os.Open("../doc/hangman.txt")

	if err != nil {
		log.Fatal(err)
	}

	// declare chunk size
	const maxSz = 79
	var listHangState []string

	// create buffer
	b := make([]byte, maxSz)

	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		listHangState = append(listHangState, string(b[:readTotal])) // print content from buffer
	}

	file.Close()
	return listHangState
}
