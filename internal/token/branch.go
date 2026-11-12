//ff:type feature=token type=model
//ff:what Branch — one way of reading a stretch of a unit: its words (more than one when the reading puts a boundary inside the stretch), the least verdict and the summed preference
package token

type Branch struct {
	Words   []Eojeol
	Verdict float64
	Pref    float64
}
