//ff:type feature=token type=model
//ff:what JSONSegment — one stretch of a sentence as nlgeul token --format json writes it: its place, its letters, the rank and the sum of lambdas of the way sent (the best), the number of ways (a number, a string when it does not fit in 2^53), and the words of that way (none for a stretch nothing is known of)
package token

type JSONSegment struct {
	Region [2]int     `json:"region"`
	Text   string     `json:"text"`
	Rank   int        `json:"rank"`
	Lambda *float64   `json:"lambda"`
	Ways   any        `json:"ways"`
	Words  []JSONWord `json:"words"`
}
