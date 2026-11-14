//ff:type feature=token type=model
//ff:what numeral — a number before a unit that is not a run of digits: its letters and its preference
package token

type numeral struct {
	text string
	pref float64
}
