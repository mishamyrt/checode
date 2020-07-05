package colours

import (
	"fmt"
	"strconv"
)

var esc = "\033"
var reset = esc + "[0m"

func printWithCode(s string, c int) string {
	return fmt.Sprintf(esc+"["+strconv.Itoa(c)+"m%s"+reset, s)
}

// Grey text will be returned
func Grey(s string) string {
	return printWithCode(s, 2)
}

// Underline text will be returned
func Underline(s string) string {
	return printWithCode(s, 4)
}

// Red text will be returned
func Red(s string) string {
	return printWithCode(s, 31)
}

// Yellow text will be returned
func Yellow(s string) string {
	return printWithCode(s, 33)
}

// Blue text will be returned
func Blue(s string) string {
	return printWithCode(s, 34)
}
