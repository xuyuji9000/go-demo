package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Fprint(pipeWriter, "Geeks\n")
		}

		pipeWriter.Close()
	}()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(pipeReader)

	fmt.Print(buffer.String())
}
