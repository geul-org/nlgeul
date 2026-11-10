//ff:type feature=mini type=model
//ff:what chainFact — what was read as a particle chain at the word's end: its length, its particles, its last particle and that particle's count, and the case of its outermost case particle
package mini

// size is in half syllables (0: no chain was read). parts is the
// particles in order, joined with "+" (에서+는). kase is 0 when the
// chain holds no case particle (only auxiliaries, connectives,
// suffixes).
type chainFact struct {
	size  int
	parts string
	last  string
	count int
	kase  int
}
