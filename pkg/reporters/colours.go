package reporters

import (
	"fmt"
	"strconv"
)

var esc = "\033"
var reset = esc + "[0m"

func printWithCode(s string, c int) string {
	return fmt.Sprintf(esc+"["+strconv.Itoa(c)+"m%s"+reset, s)
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
