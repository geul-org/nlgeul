//ff:type feature=token type=model
//ff:what path — a candidate of a unit up to one place, as a chain of words: the last word, the candidate before it, the whole printed so far, the least verdict and the summed preference, and for a state of the net (a word with all the ways it is reached) the states before it (preds, in the order they first came, the root among them as the first of the unit) and where its word begins
package token

type path struct {
	prev    *path
	last    step
	sig     string
	verdict float64
	pref    float64
	preds   []*path
	start   int
}
