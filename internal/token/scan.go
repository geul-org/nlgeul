//ff:type feature=token type=model
//ff:what scan — one unit under reading: the index, the unit's tokens, the places inside an atom (where no word may begin or end), the atoms by where they start, the runs of digits and a unit cut in two by where they start (and where they end), the places where a number ends, the places where a word that ends in a modifying closing ends in some candidate, the places where a determiner ends in some candidate, for each place the places from which the letters up to it read as markers alone, the trace lines when a trace is asked for, whether the scan reads the unit as markers that attach to the unit before (set when such a scan is made, and not changed), whether a sign that ends a sentence (. ? !) stands right after the letters that are read, and how many letters at the unit's head are a sign of a pair or a span of code, with what that is called, the signs between letters that are also read as a word of their own (by their place), and the candidates kept before each of them (by their place)
package token

type scan struct {
	ix       *index
	s        []rune
	inside   []bool
	atoms    map[int]atom
	wide     map[int]int
	numEnds  map[int]bool
	adnEnds  map[int]bool
	detEnds  map[int]bool
	marks    []map[int]bool
	trace    *[]string
	attach   bool
	closed   bool
	head     int
	headName string
	mid      map[int]midSign
	kept     map[int][]*path
}
