//ff:type feature=rule type=model
//ff:what Near — one word of a sentence as one of its candidates reads it, as far as the rules of a sentence ask: the stem as it is printed, the stem's name (보조 용언:있, 열쇠:형용사; "" for a new word), the name of its last piece ("" when the stem is all there is), whether it carries a closing (it is a predicate), and the names of its markers (the parts after the stem) in order, joined with "+" (MarksOf gives them apart; kept as one string so that two of these can be compared)
package rule

type Near struct {
	Stem  string
	Kind  string
	Last  string
	Pred  bool
	Marks string
}
