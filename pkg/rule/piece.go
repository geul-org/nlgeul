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
// word of the word table or a fused form (for a suffix read after a
// nominal, that the stem before it and the suffix together are a word of
// the word table), and Lone that the number is a lone Sino-Korean digit
// read with no word table. Adn tells that the
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
// stem's or the marker's letters are a wrong way to write another word (a
// stem the word table gives so, whose form the stem then is; a closing
// of the table of closings written wrong, whose form the closing then is).
// YoOpen tells that the auxiliary particle 요 stands right after the
// piece and that the piece stands right on its stem, the stem ending in a
// Hangul syllable with no final consonant, or right after the suffix 이
// that stands on a stem (에요 after 고양이 and after 고양+이: the letters
// are those of the closing 예요 written wrong). Order tells that the
// number right before the piece begins with the prefix 제, or that a
// joint to a following number (의 or 제 and a digit) stands right after
// the piece: it numbers a part of a provision (제5조, 31조의4, 36조제1항).
// Apart tells that the letters of a new word hold, right after a run
// of Latin letters at its head, a word the spacing articles write
// apart (a title of article 48 or a dependent noun of article 42), and
// that they read there as nothing else (A씨, B군, PD님). Conn tells that
// the piece is the stem of an auxiliary verb with a predicate tail and
// that the word right before it ends in the linking closing 어/아 (an
// auxiliary verb written onto its main verb: 주 in 해줘).
// Change tells that a new word stands over a place where a run of
// Latin letters and digits meets a Hangul syllable (a place where a word
// may end with no marker), and that the two sides cut there each read
// as a known word or as such a run (A씨, K팝).
// Glossed tells that the piece stands in a unit whose letters (after
// its head) are Hangul syllables, right after which a round bracket
// holds nothing but Han characters, as many as those syllables: the
// letters are the sound of the characters (대곤(大棍), 관치(官治)).
// UnitWord tells that the piece is a stem that begins right after a
// number and whose letters are a unit of the unit table and then a word
// written after a unit (5년간: 년 and 간, 6월경: 월 and 경).
// Bound tells that the piece is the stem of a dependent noun that the
// parse layer lets stand on the word it follows (것 and its spoken form
// 거) and that the word right before it ends in a modifying closing (a
// dependent noun written onto its modifier: 거 in 하는거).
// Modifier tells that the letters of a new word split at a place inside
// them into a modifier the parse layer lists and a known word of what it
// modifies, which the spacing rules write apart: a negative adverb before
// a predicate (못와), a numeral determiner before a nominal (두손), an
// adverb before an adverb (좀더), the modifying form of a predicate before
// the dependent noun 것 or 거 (할거).
// Main tells that the piece is a stem that begins its unit, that its word
// ends in the linking closing 어/아, and that the letters right after that
// word begin with the stem of an auxiliary verb (a main verb with an
// auxiliary verb written onto it: 가 in 가봐).
// Quoted tells that the piece stands right after a quotation mark that
// closes (what is quoted ends there, and the piece is written onto the
// mark: 는 in 간다"는). Called tells that the piece ends the letters of
// its unit that are read and that a comma or an exclamation mark stands
// right after it (where one who is called is written: 영숙이, 영숙이!).
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
	YoOpen   bool
	Order    bool
	Apart    bool
	Conn     bool
	Change   bool
	Glossed  bool
	UnitWord bool
	Bound    bool
	Modifier bool
	Main     bool
	Quoted   bool
	Called   bool
}
