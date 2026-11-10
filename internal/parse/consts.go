package parse

// Sentence constituents kopas parse prints (the seven of school
// grammar) and the name for a word no rule speaks for.
const (
	labelSubj = "주어"
	labelObj  = "목적어"
	labelComp = "보어"
	labelPred = "서술어"
	labelAdn  = "관형어"
	labelAdv  = "부사어"
	labelVoc  = "독립어"
	labelNone = "미정"
)

// Cases a particle marks, as particles.md names its function. A
// particle with none of them is an auxiliary (보조사), a connective or
// a suffix.
const (
	caseSubj = 1 << iota // 주격
	caseComp             // 보격 (이/가 before 되다, 아니다)
	caseObj              // 목적격
	caseAdn              // 관형격
	caseAdv              // 부사격
	caseVoc              // 호격
	caseBare             // no particle at all: a bound noun written onto a modifying closing (아는거)
)

// Kinds of closing, as endings.md section 3 marks them.
const (
	kindFinal = 1 << iota // 종: 종결
	kindLink              // 연: 연결
	kindAdn               // 꾸: 꾸밈(관형사형)
	kindNoun              // 이: 이름(명사형)
)

// What the syllable before a piece must look like.
const (
	afterAny     = iota // no condition
	afterBatchim        // ends in a consonant
	afterHard           // ends in a consonant other than ㄹ (으로)
	afterOpen           // ends in a vowel or ㄹ
	afterVowel          // ends in a vowel
)

// Final-consonant indexes of a Hangul syllable, and the lead of an
// ending that fuses into the syllable before it.
const (
	leadNone = 0  // the ending is whole syllables
	leadA    = -1 // 어/아/여, whole or fused into the stem's last vowel
	jongN    = 4  // ㄴ
	jongL    = 8  // ㄹ
	jongM    = 16 // ㅁ
	jongB    = 17 // ㅂ
	jongSS   = 20 // ㅆ
)

// Ending flags: how a closing attaches. None of them names a constituent.
const (
	fEu      = 1 << iota // takes 으 after a consonant, attaches bare after a vowel or ㄹ
	fNoIda               // does not attach to the copula stem 이 (는, 은, 을, 습니다, and the closings of verbs)
	fNoYo                // 요 does not follow it
	fNeedIda             // attaches only to the copula stem 이 or to 아니 (에요, 야)
	fSelfIda             // holds the copula stem itself (예요 = 이 + 에요)
)

// idaCount is the observed count of the copula 이다 in morphemes.md.
const idaCount = 4520000

// keyWord and keyTail are the context keys the rules read.
const (
	keyWord = "word"
	keyTail = "tail"
)
