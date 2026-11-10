//ff:func feature=token type=reader control=selection
//ff:what vowelRestores — the stems a syllable may come from when 어/아 is fused into its vowel (endings.md section 4): 해 from 하+여, ㅘ from ㅗ+아, ㅝ from ㅜ+어, ㅕ from ㅣ+어, ㅙ from ㅚ+어; ㅏ and ㅓ from the same vowel or from a dropped ㅡ; ㅐ and ㅔ from the same vowel; 놔 also from 놓+아 (clause 35 of the orthography)
package token

import "github.com/geul-org/koparse/pkg/rule"

func vowelRestores(r rune) []restore {
	v := rule.JungOf(r)
	to := func(jung int) rune { return r + rune((jung-v)*28) }
	switch {
	case r == '해':
		return []restore{{'하', "여"}}
	case r == '놔':
		return []restore{{'노', "아"}, {'놓', "아"}}
	case v == 9:
		return []restore{{to(8), "아"}}
	case v == 14:
		return []restore{{to(13), "어"}}
	case v == 6:
		return []restore{{to(20), "어"}}
	case v == 10:
		return []restore{{to(11), "어"}}
	case v == 0:
		return []restore{{r, "아"}, {to(18), "아"}}
	case v == 4:
		return []restore{{r, "어"}, {to(18), "어"}}
	}
	return []restore{{r, "어"}}
}
