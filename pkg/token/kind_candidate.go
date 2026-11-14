//ff:type feature=token type=model
//ff:what kindCandidate — one way a stem stands with what follows it: the kind it stands as (0 when no kind is told), the name of that candidate and the verdict of that kind's graph
package token

type kindCandidate struct {
	kind int
	name string
	v    float64
}
