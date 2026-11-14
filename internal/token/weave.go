//ff:type feature=token type=model
//ff:what weave — a portion of the net laid out for counting: the index (to put words in order), where it begins and ends, its states (by their place in the portion), the spans where words stand (sorted by start and then end) and the place of each, the states that stand at each span, for each state the states before it that are in the portion (preds), the span it stands at, for each span the spans that follow it (some state there has a state before it at the first), a stamp for telling which states a set holds, and how many (span, set) keys the ways counted
package token

type weave struct {
	ix     *index
	from   int
	to     int
	states []*path
	spans  []Span
	at     map[Span]int
	of     [][]int
	preds  [][]int
	node   []int
	succ   [][]int
	mark   []int
	tick   int
	sets   int
}
