//ff:func feature=token type=helper control=sequence
//ff:what unread — the letters between two places as a stretch that nothing reads: a stretch with no word, which the stream sends as it is (no nodes, Rank 0); counted
package token

import "math/big"

func (r *netRun) unread(from, to int) unitStretch {
	r.stats.unread++
	cs := cells{from: from, ways: new(big.Int)}
	return unitStretch{contest: Contest{Text: string(r.sc.s[from:to])}, cells: &cs}
}
