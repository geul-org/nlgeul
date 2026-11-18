//ff:type feature=rule type=model
//ff:what Near — one word of a sentence as one of its candidates reads it, as far as the rules of a sentence ask: the stem as it is printed, the stem's name (보조 용언:있, 열쇠:형용사; "" for a new word), the name of its last piece ("" when the stem is all there is), whether it carries a closing (it is a predicate), and the names of its markers (the parts after the stem) in order, joined with "+" (MarksOf gives them apart; kept as one string so that two of these can be compared), whether no space stands right before the word (Glued: it is written onto the word before it), and whether it is a run of signs alone with no name (Signs: ○○ for a name kept back)
package rule

type Near struct {
	Stem  string
	Kind  string
	Last  string
	Pred  bool
	Marks string
	// Glued tells that no space stands between the word and the word
	// before it, Nominal that the word ends in a nominal (its last part
	// is a nominal stem, an order or a unit after a number), and Text is
	// the word's letters. The reader fills them.
	Glued   bool
	Nominal bool
	Text    string
	// Modifiable tells that a word ending in a modifying closing may
	// stand right before the word: its stem is a nominal or a
	// determiner, or it ends in a modifying or a naming form. The
	// reader fills it.
	Modifiable bool
	// Signs tells that the word is a run of signs alone with no name
	// (no letter, digit or name of a sign): in a text such a run stands
	// for a name kept back (○○, □□□). The reader fills it.
	Signs bool
}
