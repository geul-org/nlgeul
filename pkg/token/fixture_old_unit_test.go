//ff:func feature=token type=helper control=sequence
//ff:what oldUnit — one unit read the old way, as the tests compare the stream with: readUnit, but a sign of a pair with letters written onto it read in two (readOnto); the stream reads such a unit by ontoReader, which gives readUnit only the sign
package token

func oldUnit(ix *index, u unit, trace bool) ([]Contest, []string) {
	if u.name == pairName && u.head < len([]rune(u.text)) {
		return readOnto(ix, u, trace)
	}
	return readUnit(ix, u, trace)
}
