//ff:type feature=token type=model
//ff:what entry — one reading of one marker in the token index: its graph's name, its surface form and how it attaches, its class and place, whether it is read only with a piece after it, whether it is read against the tables' attachment rule, its observed count, and its judgment graph
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// wide marks what is read only with a word table; gate is then the
// graph the word before must pass, and tail the kind of tail the piece
// makes when it is not the usual one. slot is, for a particle, its kind in chains (rule.Link…); for a slot
// before a closing, its order. bars, needs and holds are a closing's
// ways with the copula stem 이: it does not attach to it, attaches only
// to it (or to 아니), or holds it. lead is 0 for a form of whole syllables, rule.LeadVowel, or the final
// consonant the piece puts under the syllable before form. eu tells
// that a closing takes the link vowel 으 after a consonant, and head
// what a slot or closing begins with (head…). loose marks a form read
// where the tables do not let it attach (은 after a stem's ㄹ): its
// graph rebuts it.
type entry struct {
	name  string
	form  string
	class int
	lead  int
	after int
	slot  int
	kase  int
	noYo  bool
	inner bool
	bars  bool
	needs bool
	holds bool
	wide  bool
	loose bool
	eu    bool
	head  int
	tail  int
	gate  *entry
	count int
	g     *toulmin.Graph
}
