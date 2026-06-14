package infra

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trec/internal/core/logger"
)

type ConsoleReader struct {
	r *bufio.Reader
}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{
		r: bufio.NewReader(os.Stdin),
	}
}

func (cr ConsoleReader) Get(prompt string) (string, error) {
	fmt.Print(prompt)
	line, err := cr.r.ReadString('\n')
	if err != nil {
		return "", logger.Error("ConsoleReader", "read string error", err)
	}
	return strings.Trim(line, "\r\n"), nil
}
