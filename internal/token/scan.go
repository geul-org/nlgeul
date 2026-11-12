//ff:type feature=token type=model
//ff:what scan — one unit under reading: the index, the unit's tokens, the places inside an atom (where no word may begin or end), the atoms by where they start, the places where a number ends, the places where a word that ends in a modifying closing ends in some reading, for each place the places from which the letters up to it read as markers alone, the trace lines when a trace is asked for, whether the unit is being read as markers that attach to the unit before, whether the tails being looked up have a particle read right after them, and how many letters at the unit's head are a sign of a pair or a span of code, with what that is called
package token

type scan struct {
	ix       *index
	s        []rune
	inside   []bool
	atoms    map[int]atom
	numEnds  map[int]bool
	adnEnds  map[int]bool
	marks    []map[int]bool
	trace    *[]string
	attach   bool
	chained  bool
	head     int
	headName string
}
