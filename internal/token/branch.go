//ff:type feature=token type=model
//ff:what Branch — one way of reading a stretch of a unit: its words (more than one when the reading puts a boundary inside the stretch), the least verdict and the summed preference, and whether it is asleep (kept, not sent up as a reading to choose from)
package token

type Branch struct {
	Words   []Eojeol
	Verdict float64
	Pref    float64
	Asleep  bool
}
