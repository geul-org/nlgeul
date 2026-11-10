//ff:func feature=token type=reader control=sequence
//ff:what bareBases — the syllables a lone final consonant is read under: ㅁ under 하, 되, 이; ㄴ, ㄹ and ㅂ under 하, 되, 시, 키, 이
package token

func bareBases(lead int) []string {
	if lead == jongM {
		return []string{"하", "되", "이"}
	}
	return []string{"하", "되", "시", "키", "이"}
}
