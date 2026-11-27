//ff:func feature=data type=helper control=iteration dimension=1
//ff:what decompValue — the characters a decomposition is made of, without its tag
package main

import (
	"strconv"
	"strings"
)

func decompValue(decomp string) string {
	var b strings.Builder
	for _, h := range strings.Fields(decomp) {
		if n, err := strconv.ParseInt(h, 16, 32); err == nil {
			b.WriteRune(rune(n))
		}
	}
	return b.String()
}
