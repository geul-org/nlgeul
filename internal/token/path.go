//ff:type feature=token type=model
//ff:what path — a reading of a unit up to one place, as a chain of words: the last word, the reading before it, the whole printed so far, the least verdict and the summed preference
package token

type path struct {
	prev    *path
	last    step
	sig     string
	verdict float64
	pref    float64
}
