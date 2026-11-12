//ff:type feature=token type=model
//ff:what cutter — a sentence being cut into units from its first letter on: the index, the sentence's letters, the spans of code (each backtick's place with its partner's) and the places inside them, the ways the signs of a pair may stand open at the place reached (each the opening signs still open, outermost first), where the letters of the unit being cut begin (a special stretch such as an address is looked for from there), and whether a trace is kept
package token

// open begins as one way with nothing open. A sign that may be read in
// more ways than one leaves more ways than one; the depth of a candidate
// is how many signs stand open in the way it is read in.
type cutter struct {
	ix     *index
	r      []rune
	code   map[int]int
	inside []bool
	open   []string
	trace  bool
	at     int
}
