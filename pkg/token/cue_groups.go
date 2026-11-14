//ff:func feature=token type=reader control=iteration dimension=1
//ff:what cueGroups — a Hangul numeral that begins at place p (not right after a Sino-Korean numeral letter) read as groups: a small number (digits and 십, 백, 천, not empty) and the greater number 만, 억 or 조 after it, again and again, the last group a small number with none after it; when there are two groups or more each stands as a word of no markers (stands)
package token

import "strings"

func (sc *scan) cueGroups(p int) {
	if p > 0 && strings.ContainsRune(sinoDigits+sinoBigs, sc.s[p-1]) {
		return
	}
	var ends []int
	for i := p; ; {
		j := sc.smallEnd(i)
		if j == i || sinoTokens(string(sc.s[i:j])) == 0 {
			break
		}
		if j < len(sc.s) && strings.ContainsRune("만억조", sc.s[j]) {
			ends, i = append(ends, j+1), j+1
			continue
		}
		ends = append(ends, j)
		break
	}
	if len(ends) < 2 {
		return
	}
	for from, k := p, 0; k < len(ends); k++ {
		sc.stand(ends[k], from)
		from = ends[k]
	}
}
