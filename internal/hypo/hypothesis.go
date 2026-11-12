//ff:type feature=hypo type=model
//ff:what Hypothesis — one reading picked for every bundle of a sentence: its number among the combinations (Picked turns the number back into the readings) and its score
package hypo

// The numbers run in the order of the readings as the token layer gave
// them, the first bundle counting most: number 0 picks the first
// reading of every bundle.
type Hypothesis struct {
	Number int
	Score  Score
}
