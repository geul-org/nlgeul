//ff:func feature=token type=reader control=iteration dimension=1
//ff:what structure — the cost of a word's structure (with no word table, or with its third layer): a boundary inside the unit, the stem, and each piece that leaves no trace; the markers of the index cost nothing
package token

func (ix *index) structure(b int, pieces []piece) float64 {
	if ix.layers != layerNone && ix.layers < layerCost {
		return 0
	}
	cost := costPiece
	if b > 0 {
		cost += costBoundary
	}
	for _, p := range pieces {
		if p.unseen {
			cost += costUnseen
		}
	}
	return cost
}
