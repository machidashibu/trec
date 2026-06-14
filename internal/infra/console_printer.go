package infra

import (
	"fmt"
)

type ConsolePrinter struct{}

func NewConsolePrinter() *ConsolePrinter {
	return &ConsolePrinter{}
}

func (cp ConsolePrinter) Clear() {
	fmt.Print("\033[H\033[2J")
}

func (cp ConsolePrinter) PrintLine(text string) {
	fmt.Printf("\r\033[2K%s", text)
}

func (cp ConsolePrinter) Print(text string) {
	fmt.Println(text)
}
