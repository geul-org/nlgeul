//ff:func feature=token type=helper control=sequence
//ff:what wordKeyOf — the word a candidate ends with by where it begins (where the candidate before it ends), where it ends and how it is printed
package token

func wordKeyOf(p *path) wordKey {
	return wordKey{start: p.prev.last.end, end: p.last.end, sig: p.last.sig}
}
