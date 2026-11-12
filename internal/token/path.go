//ff:type feature=token type=model
//ff:what path — a reading of a unit up to one place, as a chain of words: the last word, the reading before it, the whole printed so far, the least verdict and the summed preference, and whether any word of it is read from letters written wrong
package token

type path struct {
	prev    *path
	last    step
	sig     string
	verdict float64
	pref    float64
	typo    bool
}
