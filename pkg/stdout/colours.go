package stdout

import (
	"fmt"
	"strconv"

	"github.com/mishamyrt/checode/v1/pkg/config"
)

var esc = "\033"
var reset = esc + "[0m"

func printWithCode(s string, c int) string {
	return fmt.Sprintf(esc+"["+strconv.Itoa(c)+"m%s"+reset, s)
}

func isSet(bitmap uint8, flag uint8) bool {
	return (bitmap & flag) == flag
}

func colorize(bitmap uint8) func(s string) string {
	if isSet(bitmap, config.ErrFlag) {
		return red
	} else if isSet(bitmap, config.WarnFlag) {
		return yellow
	}
	return blue
}

func grey(s string) string {
	return printWithCode(s, 2)
}

func underline(s string) string {
	return printWithCode(s, 4)
}

func red(s string) string {
	return printWithCode(s, 31)
}

func yellow(s string) string {
	return printWithCode(s, 33)
}

func blue(s string) string {
	return printWithCode(s, 34)
}
