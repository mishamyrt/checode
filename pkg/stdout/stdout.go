package stdout

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/parser"
)

type Printer struct {
	result string
}

func (r *Printer) Flush() {
	fmt.Print(r.result)
}

func (r *Printer) AddMatch(m parser.FileMatches) {
	r.result += printPath(m.Path)
	for _, match := range m.Matches {
		r.result += printLine(match.Line)
		r.result += printKeywords(match.Keywords, match.Flags)
		r.result += printMessage(match.Message)
		r.result += "\n"
	}
	r.result += "\n"
}

func printLine(n int) string {
	return grey(fmt.Sprintf("  %-5s", strconv.Itoa(n)))
}

func printPath(p string) string {
	return underline(fmt.Sprintln(p))
}

func printKeywords(k []string, bitmap bit.Map) (kw string) {
	kw = strings.Join(k, ": ") + ":"
	kw = colorize(bitmap)(kw)
	return
}

func printMessage(m string) string {
	return fmt.Sprintf(" %s", m)
}

func Print(res *parser.Parsing) {
	var p Printer
	for _, match := range res.Matches {
		p.AddMatch(match)
	}
	p.Flush()
}
