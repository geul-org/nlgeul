//ff:type feature=token type=model
//ff:what scan — one unit under reading: the index, the unit's tokens, the places inside an atom (where no word may begin or end), the atoms by where they start, the runs of digits and a unit cut in two by where they start (and where they end), the places where a number ends, the places where a word that ends in a modifying closing ends in some candidate, the places where a determiner ends in some candidate, for each place the places from which the letters up to it read as markers alone, the trace lines when a trace is asked for, whether the scan reads the unit as markers that attach to the unit before (set when such a scan is made, and not changed), whether a sign that ends a sentence (. ? !) stands right after the letters that are read, and how many letters at the unit's head are a sign of a pair or a span of code, with what that is called, the signs between letters that are also read as a word of their own (by their place), and the candidates kept before each of them (by their place), and for the spacing rules that are on: the words that may stand as a word of no markers inside the unit (by where they end and begin) and the places after which a word of the spacing words that stands there is shown (by where it begins and ends), the registry of the words made (reg), and whether the scan builds the net (merging the candidates that end in the same word into one state) instead of keeping each candidate, for each place how many tokens lie up to it (count, so that the length of a stem is two numbers), from which places a word may begin that ends at a place (begins: the places that are within newStemMax tokens, and the farther ones where a known word may begin), the unit's letters as one string with where each begins (text, offsets), and when only some places may begin a word, which (starts), and when the unit holds a slash or a hyphen, for each place where the run of letters with no such sign that ends there begins (parts), and the place before which a window has cut the net (floor: no state stands before it)
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
	stands   []map[int]bool
	shows    map[int]map[int]bool
	reg      *registry
	netted   bool
	count    []int
	begins   [][]int
	text     string
	offsets  []int
	starts   map[int]bool
	parts    []int
	floor    int
}
