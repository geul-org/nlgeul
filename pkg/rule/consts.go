package rule

// KeyWord and KeyTail are the context keys the rules read: the word
// under judgment (*Word) and the tail reading under judgment (Tail).
const (
	KeyWord = "word"
	KeyTail = "tail"
)

// KeyPiece is the context key of the token sequence under judgment
// (Piece), for graphs that judge one marker at one place.
const KeyPiece = "piece"

// How the syllable before a piece ends.
const (
	EndVowel = 1 << iota // a vowel
	EndL                 // ㄹ
	EndHard              // a consonant other than ㄹ
	EndOther             // not a Hangul syllable: meets every condition
)

// Cases a particle marks. A particle with none of them is an auxiliary
// (보조사), a connective or a suffix.
const (
	CaseSubj = 1 << iota // 주격
	CaseComp             // 보격 (이/가 before 되다, 아니다)
	CaseObj              // 목적격
	CaseAdn              // 관형격
	CaseAdv              // 부사격
	CaseVoc              // 호격
	CaseBare             // no particle at all: a bound noun written onto a modifying closing
)

// Kinds of closing.
const (
	KindFinal = 1 << iota // 종결
	KindLink              // 연결
	KindAdn               // 꾸밈(관형사형)
	KindNoun              // 이름(명사형)
)

// Facts of a word that WordIs tests.
const (
	WordLast  = 1 << iota // the sentence's last word
	WordBound             // begins with a dependent noun
)

// Lengths that Longer compares.
const (
	SideChain   = "chain"   // the particle chain
	SideTail    = "tail"    // the tail reading under judgment
	SideLongest = "longest" // the word's longest tail reading
)
