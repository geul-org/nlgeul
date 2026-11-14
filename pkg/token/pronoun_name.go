//ff:func feature=token type=reader control=sequence
//ff:what pronounName — a closed-set word's name with what follows it, when there is no word table: a determiner that is a pronoun too (이, 그, 저) is named a pronoun before particles or the copula; alone, and every other word, it keeps its name
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func pronounName(name, stem string, tail int) string {
	if tail == rule.TailBare || !slices.Contains(pronouns, stem) {
		return name
	}
	return strings.ReplaceAll(name, determinerLabel+":"+stem, pronounLabel+":"+stem)
}
