//ff:type feature=token type=model
//ff:what pairRow — one row of the table of the signs of a pair: the opening sign of its kind, whether the sign opens, whether it closes, whether it is unsure (read as its letter beside its candidates as a sign of a pair), and what keeps it from being such a sign (bar…)
package token

type pairRow struct {
	kind   rune
	opens  bool
	shuts  bool
	unsure bool
	bars   int
}
