//ff:type feature=token type=model
//ff:what tail — a run of markers that ends at one place: where the stem before it ends, its pieces, the stem's last syllable as it was before a piece fused into it, and whether 요 may follow
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// base is 0 when no piece is fused into the stem. close is the closing
// of a predicate tail. only names the one stem the tail attaches to
// ("" for any). again and facts are the graph and the facts of a tail
// whose word is read only right after a word that ends in a modifying
// closing: the graph is asked again on each reading (nil for any other
// tail). head is what the tail begins with (head…, of its first piece),
// and irr the kind of change the stem before it must have gone through
// (0 for none): its first piece was judged by the stem's own last
// syllable, not by the letter that stands there. wrote is the letters of
// a stem that the word table gives as a wrong way to write another
// word (shown is that word, and kinds the kinds the table gives the
// wrong form); "" for any other tail. boundTypo tells that a stem of words
// bound by a slash or a hyphen is read with the words of it that are a
// wrong way to write another word as the right form (the tail otherwise
// reads the stem as it stands).
type tail struct {
	start     int
	pieces    []piece
	base      rune
	noYo      bool
	close     *entry
	only      string
	shown     string
	label     string
	as        string
	weight    float64
	gate      *entry
	again     *entry
	facts     *rule.Piece
	head      int
	irr       int
	wrote     string
	kinds     int
	boundTypo bool
}
