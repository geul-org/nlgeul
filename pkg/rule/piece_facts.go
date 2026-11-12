//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what pieceFacts — the yes/no facts of a piece as flags (Piece…, in the order of their constants), for PieceIs
package rule

func pieceFacts(p Piece) int {
	have := 0
	for i, yes := range []bool{p.HasStem, p.Num, p.Lone, p.Whole, p.Adn, p.Turned, p.Marked, p.Next < 0, p.Keeps, p.Sole, p.Det, p.Closed, p.Same, p.Misspelt} {
		if yes {
			have |= 1 << i
		}
	}
	return have
}
