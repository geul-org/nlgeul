//ff:func feature=parse type=helper control=sequence
//ff:what share — the rival's part of the two observed counts, as a rule value: rival / (rival + mine) mapped onto [-1, 1]
package parse

// When the tables give no count for either side the two cannot be
// weighed, and the rule does not apply (-1).
func share(mine, rival int) float64 {
	if mine == 0 || rival == 0 {
		return -1
	}
	return 2*float64(rival)/float64(rival+mine) - 1
}
