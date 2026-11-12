//ff:type feature=token type=model
//ff:what atom — a stretch of a unit that is not cut inside: a run of Latin letters and digits, or runs joined by dots that a joining candidate stands for; where it ends and the candidates that join it (one name each)
package token

// names is empty for a single run that is no number. num tells that the
// atom is a number.
type atom struct {
	end   int
	names []string
	num   bool
}
