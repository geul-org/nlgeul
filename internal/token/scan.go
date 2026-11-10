//ff:type feature=token type=model
//ff:what scan — one unit under reading: the index, the unit's tokens, the places inside an atom (where no word may begin or end), the atoms by where they start, the places where a number ends, and the trace lines when a trace is asked for
package token

type scan struct {
	ix      *index
	s       []rune
	inside  []bool
	atoms   map[int]atom
	numEnds map[int]bool
	trace   *[]string
}
