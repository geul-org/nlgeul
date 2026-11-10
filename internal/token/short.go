//ff:type feature=token type=model
//ff:what short — a shortened closing read only with a word table: its row as the closing tables write one, the form it is short for, whether the copula does not take it, and the kind of tail it makes (0 for any predicate's)
package token

import "github.com/geul-org/nlgeul/internal/parse"

type short struct {
	m      parse.Marker
	origin string
	bars   bool
	tail   int
}
