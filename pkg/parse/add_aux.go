//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what addAux — the stems of the auxiliary verbs read right after a main verb's 어/아 form: those of the closed set of the grammar tables but the ones always written onto the main verb (helpers: 지), and with a word table those this layer adds (auxMore: 가, 내, 드리, 오, 달) too
package parse

func (ix *index) addAux(closed []string) {
	for _, w := range closed {
		if _, onto := helpers[atOf([]rune(w), 0)]; !onto || len([]rune(w)) > 1 {
			ix.aux = append(ix.aux, w)
		}
	}
	if ix.lexed {
		ix.aux = append(ix.aux, auxMore...)
	}
}
