package input

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadInput() string {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fileInfo.Mode()&os.ModeCharDevice != 0 {
		log.Fatal("No input data, exiting...")
	}

	reader := bufio.NewReader(os.Stdin)

	var inputRead []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		inputRead = append(inputRead, input)
	}

	if len(inputRead) == 0 {
		log.Fatal("No input data, exiting...")
	}

	return string(inputRead)
}
