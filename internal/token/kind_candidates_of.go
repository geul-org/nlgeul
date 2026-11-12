//ff:func feature=token type=helper control=iteration dimension=1
//ff:what kindCandidatesOf — the candidates of a sign as the candidates of a stem that stands as no kind of the tables (kind 0)
package token

func kindCandidatesOf(marks []markCandidate) []kindCandidate {
	out := make([]kindCandidate, len(marks))
	for i, m := range marks {
		out[i] = kindCandidate{0, m.name, m.v}
	}
	return out
}
