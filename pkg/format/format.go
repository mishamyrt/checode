package format

import (
	"fmt"
	"strconv"
)

var esc = "\033"
var reset = esc + "[0m"
var grey = 2
var underline = 4
var red = 31
var yellow = 33
var blue = 34

func printWithCode(s string, c int) string {
	return fmt.Sprintf(esc+"["+strconv.Itoa(c)+"m%s"+reset, s)
}

// Grey formats colored terminal text.
func Grey(s string) string {
	return printWithCode(s, grey)
}

// Underline formats underlined terminal text.
func Underline(s string) string {
	return printWithCode(s, underline)
}

// Red formats colored terminal text.
func Red(s string) string {
	return printWithCode(s, red)
}

// Yellow formats colored terminal text.
func Yellow(s string) string {
	return printWithCode(s, yellow)
}

// Blue formats colored terminal text.
func Blue(s string) string {
	return printWithCode(s, blue)
}
