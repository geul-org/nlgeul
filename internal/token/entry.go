//ff:type feature=token type=model
//ff:what entry — one reading of one marker in the token index: its graph's name, its surface form and how it attaches, its class and place, its observed count, and its judgment graph
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// lead is 0 for a form of whole syllables, leadVowel, or the final
// consonant the piece puts under the syllable before form.
type entry struct {
	name  string
	form  string
	class int
	lead  int
	after int
	slot  int
	noYo  bool
	count int
	g     *toulmin.Graph
}
