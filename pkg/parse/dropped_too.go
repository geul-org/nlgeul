//ff:func feature=parse type=reader control=sequence
//ff:what droppedToo — is a tail whose copula stem is read as the syllable 이 right before it read once more with the stem dropped after that syllable: with a word table, when no piece is fused into the syllable and a letter stands before it; the syllable may be the last of a nominal that ends in 이 (아이, 차이, 돈벌이), after which the stem is dropped as after any vowel; a syllable 이 that begins the unit is the stem itself, said after the unit before (+이+다), not the nominal 이 with the stem dropped
package parse

// e is the copula candidate copulaOf gave the tail.
func (sc *scan) droppedToo(t tail, e *entry) bool {
	return sc.ix.lexed && e == sc.ix.copula && t.base == 0 && t.start >= 2
}
