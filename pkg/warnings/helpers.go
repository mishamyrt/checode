package warnings

import "github.com/mishamyrt/checode/v1/pkg/types"

func findKeyword(a *types.Config, x string) (keyword string, flag uint8) {
	for i, m := range *a {
		if x == i {
			return i, m
		}
	}
	return "", 0
}
