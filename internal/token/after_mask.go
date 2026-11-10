//ff:func feature=token type=reader control=selection
//ff:what afterMask — the ends a marker may attach to, from the table's attachment condition
package token

import "github.com/geul-org/koparse/pkg/rule"

func afterMask(after int) int {
	switch after {
	case afterBatchim:
		return rule.EndL | rule.EndHard | rule.EndOther
	case afterHard:
		return rule.EndHard | rule.EndOther
	case afterOpen:
		return rule.EndVowel | rule.EndL | rule.EndOther
	case afterVowel:
		return rule.EndVowel | rule.EndOther
	}
	return rule.EndVowel | rule.EndL | rule.EndHard | rule.EndOther
}
