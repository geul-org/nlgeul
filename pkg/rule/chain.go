//ff:type feature=rule type=model
//ff:what Chain — what was read as a particle chain at a word's end: its length, its particles, its last particle and that particle's observed count, the case of its outermost case particle, and the word divided at the chain
package rule

// Size is in half syllables (0: no chain was read). Parts is the
// particles in order, joined with "+". Kase is 0 when the chain holds
// no case particle, CaseBare when a dependent noun stands with no
// particle at all.
type Chain struct {
	Size   int
	Parts  string
	Last   string
	Count  int
	Kase   int
	Marked string
}
