//ff:type feature=token type=model
//ff:what reach — the best way to each state of the net from the last cut: its value (the best lambdas of the spans along the way summed), the state before it on that way (a seed or the root when it is the first), and the states at each span (the states of a span give the span's best lambda together)
package token

type reach struct {
	value map[*path]float64
	back  map[*path]*path
	group map[Span][]*path
}
