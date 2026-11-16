//ff:type feature=rule type=model
//ff:what Tail — one reading of a word's end as a predicate tail: the closing and its kind, the length read and the observed count
package rule

// Size is in half syllables (0: no tail). Count is 0 when the tables
// give none.
type Tail struct {
	Closing string
	Kind    int
	Size    int
	Count   int
}
