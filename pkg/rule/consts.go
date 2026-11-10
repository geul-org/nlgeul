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

// What the syllable before a piece must look like for the piece to
// attach (Fits).
const (
	AfterAny     = iota // no condition
	AfterBatchim        // ends in a consonant
	AfterHard           // ends in a consonant other than ㄹ (으로)
	AfterOpen           // ends in a vowel or ㄹ
	AfterVowel          // ends in a vowel
)

// Final-consonant indexes of a Hangul syllable, and the lead of a
// piece that begins with 어/아/여.
const (
	JongN     = 4  // ㄴ
	JongL     = 8  // ㄹ
	JongM     = 16 // ㅁ
	JongB     = 17 // ㅂ
	JongSS    = 20 // ㅆ
	LeadVowel = -1 // 어/아/여, whole or fused into the stem's last vowel
)

// The slots before a closing, in their order from the stem outwards
// (SlotOf, SlotBefore); a closing stands after them all.
const (
	SlotHonor   = 1 // 시
	SlotPast    = 2 // 었/았/였, which may come twice
	SlotGuess   = 3 // 겠
	SlotRecall  = 4 // 더
	SlotClosing = 9
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
