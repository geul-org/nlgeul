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

// KeyAround is the context key of the candidate judged with the words
// of its sentence around it (Around), for the graphs of a sentence.
const KeyAround = "around"

// KeyMark is the context key of the sign under judgment (Mark), for
// graphs that judge one candidate of a sign such as the full stop.
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
	AfterLiquid         // ends in ㄹ
	AfterHangul         // is a Hangul syllable
	AfterNotN           // ends in a consonant other than ㄴ (률)
	AfterOpenN          // ends in a vowel or ㄴ (율)
	AfterHa             // is the syllable 하 (하야, the old spelling of 하여)
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
	SlotPolite  = 5 // 오/옵, 사오/사옵, 자오/자옵 (old polite forms), read right before a closing
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
	TailExist              // a predicate tail whose closing verbs take, and of the adjectives the stems of 있다, 없다 and 계시다 (느냐)
	TailPolite             // the auxiliary particle 요 and nothing else
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
	LinkHonor         // 높임의 주격 께서와 단체 뒤 주격 에서: 보조사를 받는다 (께서는, 께서도, 정부에서는)
	LinkEmph          // 부사격 조사 뒤의 강조 보조사 가 (에서가, 에가)
	LinkCopula        // 서술격 이다의 줄기: 접미 뒤에 선다(꽃들이다)
	LinkLast          // 조사가 뒤에 오지 않는 접미: 체언 뒤의 당 (팀당)
)

// What stands right after a piece (Piece.Right).
const (
	RightDone     = 1 << iota // the unit ends, or the piece is whole as the tables give it
	RightParticle             // a particle
	RightOther                // something else
)

// Yes/no facts of a piece that PieceIs tests.
const (
	PieceStem     = 1 << iota // a stem stands before it (for a stem's restored syllable: letters of the stem stand before it)
	PieceNum                  // a number stands right before it
	PieceLone                 // that number is a lone Sino-Korean digit read with no word table
	PieceWhole                // that number and the piece together are a word (for a suffix read after a nominal: the stem before it and the suffix together are a word of the word table)
	PieceAdn                  // the word right before it ends in a modifying closing
	PieceTurned               // its own word is a modifying or a naming form
	PieceMarked               // it is a new word after a boundary whose letters read as markers
	PieceLast                 // no piece follows it
	PieceKeeps                // it is a stem that keeps its last vowel before 어/아 (a 르 or 러 irregular)
	PieceSole                 // it is a stem whose word is the whole of its unit
	PieceDet                  // the word right before it is a determiner
	PieceClosed               // it stands right on its stem and ends its unit, and a sign that ends a sentence (. ? !) stands right after it
	PieceSame                 // it is a restored syllable that stands in the letters as it is (the piece fused into it left no trace)
	PieceMisspelt             // it is a stem or a marker whose letters are a wrong way to write another word (the word table gives the stem so; the marker is a closing of the table of closings written wrong)
	PieceYoOpen               // the particle 요 stands right after it, and it stands right on a stem that ends in a Hangul syllable with no final consonant or right after the suffix 이 on a stem
	PieceOrder                // the number right before it begins with the prefix 제, or a joint to a following number (의 or 제 and a digit) stands right after it
	PieceApart                // its letters hold, right after a run of Latin letters, a word the spacing articles write apart
	PieceConn                 // the word right before it ends in the linking 어/아 and it is an auxiliary's stem with a predicate tail
	PieceChange               // it is a new word over a place where a run of Latin letters and digits meets a Hangul syllable, and the two sides cut there each read as a known word or such a run
	PieceGlossed              // its unit's letters are Hangul syllables, and right after them stands a round bracket that holds Han characters alone, as many as those syllables (the letters are their sound: 대곤(大棍))
	PieceUnitWord             // it is a stem that begins right after a number, and its letters are a unit of the unit table and then a word written after a unit (5년간)
	PieceBound                // the word right before it ends in a modifying closing and it is the stem of a dependent noun that stands on its modifier (것, 거)
	PieceModifier             // it is a new word whose letters are a modifier and a known word it modifies written together (못와, 두손, 좀더, 할거): a rebuttal of the new word, which stands behind the two words read apart
	PieceMain                 // it is a main verb's stem that begins its unit, and its word ends in the linking 어/아 right before an auxiliary verb's stem written onto it (가 in 가봐)
	PieceQuoted               // it stands right after a quotation mark that closes (는 in 간다"는)
	PieceCalled               // it ends the letters of its unit that are read, and a comma or an exclamation mark stands right after it (where one who is called is written: 영숙이, 영숙이!)
)

// Yes/no facts of a sign that MarkIs tests.
const (
	MarkRepeated   = 1 << iota // two of the sign or more stand in a row
	MarkAfterWord              // a letter stands right before the sign
	MarkNumeric                // every run of the dotted name is digits
	MarkAfterDigit             // a digit stands right before the sign
	MarkAlone                  // the sign (with the same sign in a row) is the whole of its unit
	MarkEnds                   // it stands after the letters of its unit that are read (at a word's end)
	MarkLeads                  // the letters of its unit that are read begin right after it
	MarkHeads                  // it stands at the very start of an input line
	MarkSpaced                 // a space stands right after it
	MarkOpened                 // a sign of its pair's kind may stand open before it
	MarkApart                  // a space, or the sentence's edge, stands on both sides of it
	MarkArrow                  // a hyphen or an equals sign stands right next to it (->, =>, <=)
	MarkBySign                 // a punctuation mark or symbol other than itself stands right next to it
	MarkOpenedAny              // a sign of its pair's kind may stand open before it in at least one of the ways the signs before it may stand open (MarkOpened tells it of one way)
	MarkStopped                // a full stop stands right after the last of the runs the sign joins (1.28. — a date: the orthography's 마침표 (3))
	MarkAfterStop              // a mark that ends a sentence (. ? ! …) stands right before it: the end of the sentence a closing sign closes on (다.』)
	MarkMated                  // a space, or the sentence's edge, stands on both sides of it, and a sign that pairs with it stands so too on the other side of what it encloses: after an opening sign a closing sign of its kind, before a closing sign an opening one (< 제목 >)
	MarkQuoted                 // a marker that takes a quotation (라고, 고, 는, 라는, 란 …) stands right after it and no letter or digit after that marker: a closing quotation mark's place ("라고, '는)
	MarkOpenSide               // a space or the sentence's start stands right before it and a letter or digit right after it, and it is not MarkQuoted: where an opening quotation mark is written
	MarkShutSide               // it is MarkQuoted, or a letter stands right before it and a space or the sentence's end right after it: where a closing quotation mark is written
	MarkBare                   // its input line holds no letter and no digit: a line of signs alone (a rule ---, a table's line |---|---|, a frame -----------)
	MarkDrawn                  // a sign that is no punctuation mark of a sentence stands in the run of signs it is in (-->, <!--, |---|, ├──┼): the signs draw something together
	MarkSum                    // it stands alone where a sum is written: inside a bracket still open, or between two numbers with a space on each side
	MarkTouched                // a sign other than itself that is no punctuation mark of a sentence stands right next to it (=>, |-, ★☆; a sign of the unit table read with the number before it is none: 5%↓)
)

// The facts a list is asked for (List.Of).
const (
	ListForm      = iota + 1 // the piece's form
	ListBefore               // the name of the piece before it
	ListAfter                // the name of the piece after it
	ListBase                 // the syllable a fused piece sits under
	ListExt                  // the run right of a dot
	ListLast                 // the last run of a dotted name
	ListWord                 // the word under judgment
	ListClosing              // the closing of the tail under judgment
	ListSelfStem             // the stem of the candidate judged in its sentence
	ListSelfKind             // the name of that stem
	ListSelfLast             // the name of that candidate's last piece
	ListPrevLast             // the name of the last piece of the word before it
	ListNextStem             // the stem of the word after it
	ListSelfMarks            // the names of that candidate's markers, each (SelfHas reads them; Listed does not)
	ListNextLast             // the name of the last piece of the word after it
	ListNext2Stem            // the stem of the word after that
	ListPrevKind             // the kind of the name of the stem of the word before it (닫기 of 닫기:1)
	ListPrevStem             // the stem of the word before it
	ListNextKind             // the kind of the name of the stem of the word after it (닫기 of 닫기:1)
	ListSelfText             // the letters of the candidate's word (한 of 하다 with ㄴ)
	ListPrevName             // the name of the stem of the word before it, whole (열쇠:동사)
	ListNext2Last            // the name of the last piece of the word after that
	ListNextName             // the name of the stem of the word after it, whole (원:수)
)

// Yes/no facts of a candidate judged in its sentence that AroundIs
// tests.
const (
	AroundGlued              = 1 << iota // no space stands between the candidate's word and the word before it
	AroundAfterNominal                   // the word right before it ends in a nominal (Near.Nominal)
	AroundBeforeModifiable               // a word ending in a modifying closing may stand right before the word right after it (Near.Modifiable)
	AroundBeforeGlued                    // the word right after it is written onto it: no space stands between them (Near.Glued of that word)
	AroundAfterSigns                     // the word right before it is a run of signs alone with no name, which stands for a name kept back (Near.Signs)
	AroundAfterAdnominal                 // each candidate of the word right before it is an adnominal: it ends in a modifying closing or is a determiner standing alone (Near.AdnominalAll)
	AroundBeforeUnmodifiable             // no adnominal may stand right before the word right after it, by any of its candidates (Near.Unmodifiable)
)
