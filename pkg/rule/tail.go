//ff:type feature=rule type=model
//ff:what Tail — one reading of a word's end as a predicate tail: the closing and its kind, the stem before it, the length read, whether the copula stem was read, the observed count, and the word divided into its pieces
package rule

// Size is in half syllables (0: no tail). Count is 0 when the tables
// give none.
type Tail struct {
	Closing string
	Kind    int
	Stem    string
	Size    int
	Copula  bool
	Count   int
	Marked  string
}
