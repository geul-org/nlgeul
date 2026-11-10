package token

// Irregular predicates (Hangul orthography, clause 18). A word table
// writes a predicate under its stem as the dictionary has it; in a
// sentence the stem's last syllable changes before some tails. A
// changed form is never a stem of its own: it is read only together
// with a tail it changes before, and the word is printed, and looked
// up, under the stem of the table (들었다 → 듣+었+다). A changed form
// that is also another word stands next to that word, as a contest.
//
// Kinds of change, as bits. irrL is not irregular: every stem that
// ends in ㄹ drops it before ㄴ, ㅂ, ㅅ and a fused ㄹ, whatever the
// table says of it.
const (
	irrL   = 1 << iota // ㄹ 탈락: 살 → 사 (사는, 삽니다, 사시고, 살까 = 사+ㄹ까)
	irrD               // ㄷ: 듣 → 들 (들어, 들으니)
	irrS               // ㅅ: 짓 → 지 (지어, 지으니)
	irrB               // ㅂ: 돕 → 도우 (도운, 도우니, 더워)
	irrBO              // ㅂ, 돕 and 곱 only: 도오 (도와)
	irrH               // ㅎ: 파랗 → 파라 (파란, 파라면, 그러면)
	irrHE              // ㅎ with 어 fused: 파래, 그래
	irrReu             // 르: 모르 → 몰르 (몰라 = 몰르+아 with ㅡ dropped)
	irrReo             // 러: 이르 → 이르르 (이르러)
	irrU               // 우: 푸 → 프 (퍼)
)

// typeIrr reads the irregular type a word table gives a predicate.
var typeIrr = map[string]int{"ㄷ": irrD, "ㅅ": irrS, "ㅂ": irrB, "ㅎ": irrH, "르": irrReu, "러": irrReo, "우": irrU}

// What a tail begins with, as bits of its first piece (headOf).
const (
	headFused  = 1 << iota // 어/아 or 었/았 fused into the stem's last syllable
	headWhole              // 어/아/여, 었/았/였 as a syllable of its own
	headEuForm             // the form after a consonant of a tail that takes 으: 으…, 은, 을, 음, 읍
	headEuOpen             // the form after a vowel of such a tail: 면, 니, 시, a fused ㄴ·ㄹ·ㅁ
	headN                  // ㄴ: 는, 니, 네, a fused ㄴ
	headS                  // ㅅ: 시, 세요, 습니다
	headL                  // a fused ㄹ
	headB                  // a fused ㅂ
	headNe                 // 네, before which ㅎ stays or goes
)

// irrHeads says before what each change happens.
var irrHeads = map[int]int{
	irrL:   headN | headS | headL | headB,
	irrD:   headWhole | headEuForm,
	irrS:   headWhole | headEuForm,
	irrB:   headEuOpen | headFused,
	irrBO:  headFused,
	irrH:   headEuOpen | headNe,
	irrHE:  headFused,
	irrReu: headFused,
	irrReo: headFused,
	irrU:   headFused,
}

// irrOrder is the order the kinds are gone through in.
var irrOrder = []int{irrL, irrD, irrS, irrB, irrBO, irrH, irrHE, irrReu, irrReo, irrU}

// roundB are the two ㅂ stems whose 어 is 아 (도와, 고와); every other
// one takes 워. wideH are the ㅎ stems whose fused 어 gives ㅐ whatever
// their vowel (그래, 어때). eoStems are stems the table calls regular
// that fuse 어 the same way (그러다 → 그래).
var roundB = []string{"돕", "곱"}
var wideH = []string{"그렇", "이렇", "저렇", "어떻"}
var eoStems = map[string]string{"그러": "그래", "이러": "이래", "저러": "저래", "어쩌": "어째"}

// linkTakers are the closings of the tables that take the link vowel
// 으 as a piece of its own (읽으니, 먹었으나); the others that take it
// have a form with 으 in the tables (으면, 은, 을).
var linkTakers = []string{"맺음:니", "맺음:니까", "맺음:나", "맺음:냐"}

// Final consonants and vowels the changes turn on.
const (
	jongD   = 7  // ㄷ
	jongLM  = 10 // ㄻ: ㅁ put under a stem's ㄹ (삶, 만듦)
	jongS   = 19 // ㅅ
	jongH   = 27 // ㅎ
	jungYeo = 6  // ㅕ
	jungU   = 13 // ㅜ
	jungEu  = 18 // ㅡ
	choN    = 2  // ㄴ
	choS    = 9  // ㅅ
)
