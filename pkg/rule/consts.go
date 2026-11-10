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

// KeyMark is the context key of the sign under judgment (Mark), for
// graphs that judge one reading of a sign such as the full stop.
const KeyMark = "mark"

// What stands on one side of a sign.
const (
	SideAlnum  = 1 << iota // a run of Latin letters, digits and underscores
	SideDigit              // such a run made of digits only
	SideHangul             // a Hangul syllable
)

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

// What follows a stem in its word (StemTail).
const (
	TailBare   = 1 << iota // nothing: the stem stands alone
	TailAux                // particles, none of them a case particle
	TailCase               // particles with a case particle among them
	TailPred               // a predicate tail: slots and a closing
	TailVerb               // a predicate tail whose closing only verbs take
	TailCopula             // the copula 이다 and its tail (a noun stands before it)
	TailAdj                // a predicate tail whose closing only adjectives take with no slot before it
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

// Kinds of particle by what may follow them in a chain (Links). They
// are bit places, so that a set of kinds is a mask.
const (
	LinkSuffix = iota // 접미: 들, 씩, 쯤, 끼리, a name's 이
	LinkCase          // 주격·보격·목적격: 이/가, 께서, 을/를
	LinkAdn           // 관형격: 의
	LinkAdv           // 부사격: 에, 에서, 에게, 으로 …
	LinkComp          // 견줌의 부사격: 보다, 처럼, 같이, 만큼
	LinkWith          // 함께·이음: 과/와, 이랑
	LinkRange         // 범위·한정의 보조사: 까지, 부터, 만, 뿐, 조차, 마저, 마다, 밖에, 대로
	LinkTopic         // 주제류 보조사: 은/는, 도, 이나, 이라도, 이든, 이야 …
	LinkVoc           // 호격: 아/야, 이여
	LinkYo            // 요
	LinkDir           // 방향·출발의 부사격: 으로/로, 으로부터/로부터 (에게로, 께로부터)
	LinkHonor         // 높임의 주격 께서 (께서는, 께서도)
	LinkEmph          // 부사격 조사 뒤의 강조 보조사 가 (에서가, 에가)
)

// What stands right after a piece (Piece.Right).
const (
	RightDone     = 1 << iota // the unit ends, or the piece is whole as the tables give it
	RightParticle             // a particle
	RightOther                // something else
)
