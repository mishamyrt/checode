package stdout

import (
	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/colours"
	"github.com/mishamyrt/checode/v1/pkg/config"
)

func colorize(bitmap bit.Map) func(s string) string {
	if bitmap.IsSet(config.ErrFlag) {
		return colours.Red
	} else if bitmap.IsSet(config.WarnFlag) {
		return colours.Yellow
	}
	return colours.Blue
}
