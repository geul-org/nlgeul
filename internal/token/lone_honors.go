//ff:func feature=token type=reader control=sequence
//ff:what loneHonors — the honorific candidates of 심 (시+ㅁ, the noun form) and of 세 (시+어) that end at j: the piece under the letter is one the tables' own graphs do not read there, so it is set here and the letter's graph judges
package token

// vowel asks for 세, which stands only before 요; else for 심.
func (sc *scan) loneHonors(j int, vowel bool) []tail {
	letter, key, name, text := nounHonor, 'ㅁ', nounName, "(ㅁ)"
	if vowel {
		letter, key, name, text = vowelHonor, keyVowel, vowelName, "(어)"
	}
	if j < 1 || sc.s[j-1] != letter {
		return nil
	}
	e := sc.ix.named(key, name)
	out := sc.honors(piece{text, name, 1, sc.ix.fusedPref(e, "어", sc.ix.pref(e.count, "")), tailOf(e), false, ""}, j, !vowel)
	for i := range out {
		out[i].close = e
	}
	return out
}
