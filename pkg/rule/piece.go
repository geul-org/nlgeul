//ff:type feature=rule type=model
//ff:what Piece — the structural facts of one token sequence read as a marker or a stem at one place: whether a stem stands before it, the syllable before it, its place in a chain and the place of the piece after it, the syllable a fused piece sits under, the pieces around it, and the yes/no facts PieceIs tests
package rule

// Prev is 0 when nothing stands before the piece. Slot and Next are
// places in a particle chain or the slots before a closing; Next is -1
// when nothing follows. Base is "" unless the piece is fused into a
// stem's last syllable. Tail is set when the piece is a stem: what
// follows it in its word (Tail…); Count is then its observed count
// (negative when no table row gives one). Before and After are the
// names of the pieces before and after it, when a rule asks for them.
// Right is what stands right after the piece (Right…), and Stem the
// kinds the tables give the stem: the word before a marker, or the
// piece itself when it is a stem. Num tells that a number stands right
// before the piece; Whole that this number and the piece together are a
// word of the word table or a fused form, and Lone that the number is a
// lone Sino-Korean digit read with no word table. Adn tells that the
// word right before the piece ends in a modifying closing, and Turned
// that the piece's own word is a modifying or a naming form. Marked
// tells that the letters of a new word after a boundary read as markers
// of the index. Keeps tells that the word table gives a stem as one
// that keeps its last vowel before 어/아 (a 르 or 러 irregular). Sole
// tells that the piece is a stem whose word is the whole of its unit:
// the stem begins the unit and nothing but its own markers follows it.
// Det tells that the word right before the piece is a determiner.
// Same tells that a restored syllable stands in the letters as it is:
// the piece fused into it left no trace. Closed tells that the piece stands right on its stem and ends its
// unit, and that a full stop, a question mark or an exclamation mark
// stands right after it in the unit's letters. Misspelt tells that the
// stem's letters are what the word table gives as a wrong way to write
// another word, whose form the stem then is.
type Piece struct {
	Form     string
	HasStem  bool
	Prev     rune
	Slot     int
	Next     int
	Base     string
	Tail     int
	Count    int
	Before   string
	After    string
	Right    int
	Stem     int
	Num      bool
	Whole    bool
	Lone     bool
	Adn      bool
	Turned   bool
	Marked   bool
	Keeps    bool
	Sole     bool
	Det      bool
	Closed   bool
	Same     bool
	Misspelt bool
}
