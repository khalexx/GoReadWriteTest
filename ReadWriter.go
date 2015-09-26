package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func main() {

	var (
		text string
		buff bytes.Buffer
	)

	consoleReader := bufio.NewReader(os.Stdin)
	consoleWriter := bufio.NewWriter(os.Stdout)
	temp := bufio.NewReadWriter(bufio.NewReader(&buff), bufio.NewWriter(&buff))

	for !strings.EqualFold(strings.TrimSpace(text), "exit") {
		text, _ = consoleReader.ReadString('\n')
		temp.WriteString(text)
		temp.Writer.Flush()
		io.Copy(consoleWriter, temp)
		consoleWriter.Flush()
	}

}
