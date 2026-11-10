//ff:func feature=token type=reader control=iteration dimension=1
//ff:what auxNext — do the letters from place j begin with the stem of an auxiliary verb read after a main verb's 어/아 form
package token

import "strings"

func (sc *scan) auxNext(j int) bool {
	rest := string(sc.s[j:])
	for _, stem := range sc.ix.auxStems() {
		if strings.HasPrefix(rest, stem) {
			return true
		}
	}
	return false
}
