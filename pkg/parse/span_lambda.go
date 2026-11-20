//ff:type feature=parse type=model
//ff:what spanLambda — the best lambda of the candidates of one span of the net, with how many states stood at the span when it was worked out (a span with more states is worked out again)
package parse

type spanLambda struct {
	states int
	value  float64
}
