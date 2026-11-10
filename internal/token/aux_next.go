//ff:func feature=token type=reader control=iteration dimension=1
//ff:what auxNext — do the letters from place j begin with the stem of an auxiliary verb read after a main verb's 어/아 form, as it stands or with a piece fused into its last syllable (줘 = 주+어, 볼 = 보+ㄹ, 봤 = 보+았)
package token

func (sc *scan) auxNext(j int) bool {
	rest := sc.s[j:]
	for _, stem := range sc.ix.auxStems() {
		if auxHead(rest, []rune(stem)) {
			return true
		}
	}
	return false
}
