//ff:type feature=token type=model
//ff:what Branch — one way of reading a stretch of a unit: its words (more than one when the candidate puts a boundary inside the stretch), the least verdict and the summed preference
package token

// scores is kept beside Words (not sent up): the verdict and the
// preference of each word, so that the words of a branch can be split
// apart. A branch made without it has one word, which has the branch's
// own values.
type Branch struct {
	Words   []Eojeol
	Verdict float64
	Pref    float64
	scores  []wordScore
}
