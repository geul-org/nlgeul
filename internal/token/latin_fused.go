//ff:func feature=token type=rule control=sequence
//ff:what latinFused — is tail t, read from b, a piece taken out of a syllable that stands right after a letter that is no Hangul inside the stem (api과 read as api고+아): a piece is fused only into a Hangul stem's last syllable; 하 and 되, which make a predicate of the word before them, and 이, the copula's stem, are let be (push해, merge된, api인데)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) latinFused(b int, t tail) bool {
	if t.base == 0 || t.start-2 < b || rule.JungOf(sc.s[t.start-2]) >= 0 {
		return false
	}
	return t.base != '하' && t.base != '되' && t.base != '이'
}
