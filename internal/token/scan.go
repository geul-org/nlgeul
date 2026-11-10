//ff:type feature=token type=model
//ff:what scan — one unit under reading: the index, the unit's tokens, and the trace lines when a trace is asked for
package token

type scan struct {
	ix    *index
	s     []rune
	trace *[]string
}
