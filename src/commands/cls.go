package commands

import (
	"bytes"
	"fmt"
)

func Cls() {
	var buffer bytes.Buffer

	for i := 0; i < 25; i++ {
		buffer.WriteString("\n")
	}

	fmt.Print(buffer.String())
}
