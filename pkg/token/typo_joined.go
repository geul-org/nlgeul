//ff:func feature=token type=reader control=iteration dimension=1
//ff:what typoJoined — the misspelt forms of a written form with one more line: the kind joins the form of the same right form, or the right form is added after the others
package token

func typoJoined(have []typo, right string, kind int) []typo {
	for i := range have {
		if have[i].right == right {
			have[i].kinds |= kind
			return have
		}
	}
	return append(have, typo{right, kind})
}
