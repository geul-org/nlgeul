//ff:func feature=token type=helper control=iteration dimension=1
//ff:what kindReadsOf — the readings of a sign as the readings of a stem that stands as no kind of the tables (kind 0)
package token

func kindReadsOf(marks []markRead) []kindRead {
	out := make([]kindRead, len(marks))
	for i, m := range marks {
		out[i] = kindRead{0, m.name, m.v}
	}
	return out
}
