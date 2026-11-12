//ff:type feature=token type=model
//ff:what entry — one reading of one marker in the token index: its graph's name, its surface form and how it attaches, its class and place, whether it is read only with a piece after it, whether it is read against the tables' attachment rule, its observed count, its judgment graph, and for an entry that reads a closing written wrong the letters it is written with and the right closing it stands for
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
// graph rebuts it. wrote is, for an entry that reads a closing written
// wrong (typoClosings), its written form with its lead (ㄹ께, 을께, 읍니다),
// and right the right closing as it is printed ((ㄹ)게, 습니다); both
// are "" for any other entry. Such an entry has the name, the class and
// every way of attaching of the right closing, so that each list and
// exception that asks for the name holds for it as it does for the right one.
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
	wrote string
	right string
	g     *toulmin.Graph
}
