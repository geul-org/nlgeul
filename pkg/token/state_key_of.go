//ff:func feature=token type=helper control=sequence
//ff:what stateKeyOf — the state a candidate is: its last word by everything that is asked of it when a word is grown after it
package token

func stateKeyOf(p *path) stateKey {
	e := p.last
	return stateKey{word: wordKeyOf(p), verdict: e.verdict, pref: e.pref, back: e.back, adn: e.adn, det: e.det, conn: e.conn, kind: e.kind}
}
