//ff:type feature=token type=model
//ff:what index — the token index: from a form's last token (or the key of a fused piece) to the entries that end there, the closed-set words, the graph of a new word and its verdict, and the total observed count
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

type index struct {
	byKey   map[rune][]*entry
	words   map[string]*entry
	stem    *toulmin.Graph
	unknown float64
	total   float64
	graphs  int
	edges   int
}
