//ff:type feature=token type=model
//ff:what JSONCandidate — one candidate of a word as nlgeul token --format json writes it: its parts, its verdict, its preference and its lambda (null when the verdict is 0 or less, minus infinity)
package token

type JSONCandidate struct {
	Parts   []JSONPart `json:"parts"`
	Verdict float64    `json:"verdict"`
	Pref    float64    `json:"pref"`
	Lambda  *float64   `json:"lambda"`
}
