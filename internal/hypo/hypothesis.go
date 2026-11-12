//ff:type feature=hypo type=model
//ff:what Hypothesis — one candidate picked for every bundle of a sentence: its number among the combinations (Picked turns the number back into the candidates) and its score
package hypo

// The numbers run in the order of the candidates as the token layer gave
// them, the first bundle counting most: number 0 picks the first
// candidate of every bundle.
type Hypothesis struct {
	Number int
	Score  Score
}
