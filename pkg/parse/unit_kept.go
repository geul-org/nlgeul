//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what unitKept — where a unit's core begins when the signs right before it are a unit of the unit table that a suffix of bareSuffixes is written onto (㎡당): the longest such run of signs is kept; else where it began
package parse

import (
	"slices"
	"strings"
)

func unitKept(all []rune, from int) int {
	onto := false
	for _, s := range bareSuffixes {
		onto = onto || strings.HasPrefix(string(all[from:]), s)
	}
	if !onto {
		return from
	}
	for k := 0; k < from; k++ {
		if signRune(all[from-1]) && slices.Contains(unitWords, string(all[k:from])) {
			return k
		}
	}
	return from
}
