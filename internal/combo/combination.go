//ff:type feature=combo type=model
//ff:what Combination — one candidate picked for every contest of a sentence: its number among the combinations (Picked turns the number back into the candidates) and its score
package combo

// The numbers run in the order of the candidates as the token layer gave
// them, the first contest counting most: number 0 picks the first
// candidate of every contest.
type Combination struct {
	Number int
	Score  Score
}
