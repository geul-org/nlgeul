//ff:type feature=token type=model
//ff:what markCandidate — one candidate of a sign that stands at one place: the name of the graph that read it and its verdict
package token

type markCandidate struct {
	name string
	v    float64
}
