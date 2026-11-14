//ff:type feature=parse type=model
//ff:what chainFact — what was read as a particle chain at the word's end: its length, its particles, its last particle and that particle's count, the case of its outermost case particle, and the word divided at the chain
package parse

// size is in half syllables (0: no chain was read). parts is the
// particles in order, joined with "+" (에서+는). kase is 0 when the
// chain holds no case particle (only auxiliaries, connectives,
// suffixes), caseBare when a dependent noun stands with no particle at
// all. marked is the whole word with "+" before each piece.
type chainFact struct {
	size   int
	parts  string
	last   string
	count  int
	kase   int
	marked string
}
