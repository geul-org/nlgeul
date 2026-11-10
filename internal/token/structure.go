//ff:func feature=token type=reader control=iteration dimension=1
//ff:what structure — the cost of a word's structure (third layer of a word table): a boundary inside the unit, each stem and marker, and each piece that leaves no trace
package token

func (ix *index) structure(b int, pieces []piece) float64 {
	if ix.layers < layerCost {
		return 0
	}
	cost := costPiece * float64(1+len(pieces))
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
