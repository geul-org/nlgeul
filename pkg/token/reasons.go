package token

// Why an exception defeats what it attacks, as a trace prints it. Each
// exception carries its reason on its own node (except), so that one
// rule function says what it means where it is put.
const (
	exceptLabel = "예외"
	preferLabel = "선호"
	rebutLabel  = "반박"

	whyNoStem       = "앞에 줄기가 없다"
	whyAfter        = "앞 토큰의 끝이 맞지 않는다"
	whyNoHangul     = "앞 글자가 한글 음절이 아니다"
	whyChain        = "조사 사슬에 맞지 않는다"
	whySlot         = "선어말 칸의 순서가 어긋난다"
	whyBase         = "받침만 붙은 꼴을 읽는 음절이 아니다"
	whyWhole        = "통째로 닫힌 집합의 낱말이다"
	whyCase         = "예외 표의 줄이다"
	whyRight        = "바로 뒤에 선 것이 이 조각 뒤에 오지 못한다"
	whyKind         = "앞말이 낱말 표에서 그 갈래의 줄기가 아니다"
	whyUnlisted     = "앞말이 어느 표에도 없는 새 낱말이다"
	whyJoinedWord   = "줄기와 접미를 합친 글자열이 낱말 표의 낱말이다"
	whyShape        = "기호가 잇는 숫자 마디들이 이 풀이의 꼴이 아니다"
	whyNotStopped   = "둘째 수 바로 뒤에 마침표가 없다"
	whyNoNumber     = "바로 앞이 수가 아니다"
	whyLone         = "낱말 표 없이 홀로 선 한자어 한 음절 수사다"
	whyNoDigits     = "기호의 오른쪽이 수가 아니다"
	whyAfterWord    = "기호의 왼쪽에 글자가 붙어 있다"
	whyOrderRight   = "제 바로 뒤가 숫자가 아니다"
	whyOrderLeft    = "제 바로 앞에 수의 끝이 아닌 글자가 붙어 있다"
	whyBarred       = "뒤의 맺음은 서술격 이에 붙지 않는다"
	whyKept         = "뒤의 맺음 앞에서는 이가 빠지지 않는다"
	whyRunTakes     = "한글 아닌 글자 뒤의 빠진 이가 붙는 맺음이다"
	whyNaming       = "뒤가 이름 꼴 맺음이다"
	whyParticle     = "이름 꼴 맺음 바로 뒤에 조사가 붙는다"
	whyNotBefore    = "뒤의 조각이 이 뒤에 오는 것이 아니다"
	whyNotAfter     = "이 조각 뒤에 붙는 것이 아니다"
	whyNotTaken     = "앞의 맺음이 받는 조사가 아니다"
	whyMarked       = "경계 뒤의 통짜 새 낱말이 표지로 읽힌다"
	whyWrittenApart = "영문 글자 바로 뒤에 띄어 쓸 호칭어·의존 명사가 붙은 통짜 새 낱말이다"
	whyChange       = "영숫자 열과 한글이 맞닿은 자리를 넘는 통짜 새 낱말이고, 그 자리에서 가른 두 쪽이 모두 아는 낱말이나 영숫자 열로 읽힌다"
	whyLast         = "뒤에 조사가 없다"
	whyNoAdn        = "앞 어절이 꾸밈 맺음으로 끝나지 않는다"
	whyExistStem    = "있다·없다·계시다의 줄기다"
	whyNotExist     = "뒤의 맺음이 있다·없다·계시다가 받는 것이 아니다"
	whyAdn          = "앞 어절이 꾸밈 맺음으로 끝난다"
	whyDet          = "앞 어절이 관형사다"
	whyClosed       = "문장을 닫는 부호가 바로 뒤에서 단위를 닫는다"
	whyGlossed      = "바로 뒤의 소괄호에 같은 음절 수의 한자를 병기한 한글이다"
	whyBound        = "앞말이 의존명사다"
	whyBoundAux     = "의존 명사에 하다·싶다가 붙은 보조 용언의 줄기다"
	whyCoda         = "새 낱말의 줄기가 ㅆ 받침으로 끝난다"
	whyKnown        = "표가 아는 낱말이다"
	whyPredicate    = "표에서 용언 줄기이기도 하다"
	whyNoPredTail   = "뒤가 서술어 꼬리가 아니다"
	whyUntold       = "낱말 표가 없으면 표지가 붙은 줄기에는 묻지 않는다"
	whyStemHead     = "줄기의 첫 음절이다"
	whyMakes        = "하·되·이는 한글 아닌 글자 뒤에서도 뗀다"
	whyKeeps        = "표가 르·러 불규칙으로 적은 줄기다"
	whyClash        = "어/아가 앞 음절의 모음과 어긋난다"
	whyYeo          = "하 뒤의 어/아는 여이고 글자에 드러난다(해, 하여)"
	whyTrace        = "합쳐진 조각이 글자에 흔적을 남겼다"
	whyHelper       = "뒤가 본용언에 붙여 쓰는 보조 용언이다"
	whyOne          = "한 음절 줄기의 ㅡ 뒤는 어뿐이다"
	whyLonger       = "줄기에 앞 음절이 있다"
	whyTurned       = "그 낱말 자신이 꾸밈꼴·이름꼴이다"
	whyStemTail     = "줄기의 갈래가 뒤에 붙은 것을 받지 않는다"
	whySole         = "줄기의 낱말이 단위 전체다"
	whyConn         = "본용언의 어/아 꼴 바로 뒤에 붙여 쓴 보조 용언의 줄기다"
	whyMain         = "단위를 열고 붙여 쓴 보조 용언 바로 앞에서 어/아 꼴로 끝나는 본용언의 줄기다"
	whyFusedIn      = "뒤의 조각이 이 글자의 모음에 합쳐진 어/아나 ㅆ이다"
	whySides        = "양쪽이 맞는 글자 열이 아니다"
	whyBoth         = "양쪽이 모두 그 글자 열이다"
	whyAlone        = "점이 하나뿐이다"
	whyRepeated     = "점이 이어진다"
	whyExtension    = "확장자 목록에 없다"
	whyDomain       = "최상위 도메인 목록에 없다"
	whyVersion      = "버전의 꼴이 아니다"
	whySameSigns    = "같은 기호가 이어진다"
	whyNoEnd        = "어절 끝에 있지도 홀로 서지도 않았다"
	whyNoLead       = "뒤에 글자가 바로 붙지도 홀로 서지도 않았다"
	whyNoLetters    = "뒤에 글자가 바로 붙지 않았다"
	whyNoLineHead   = "입력 줄의 맨 처음이 아니다"
	whyNoSpace      = "바로 뒤가 빈칸이 아니다"
	whyNotOpened    = "같은 종류의 여는 기호가 열려 있지 않다"
	whyOpened       = "같은 따옴표가 열려 있다"
	whyOpenedAny    = "같은 종류의 여는 기호가 열려 있을 수 있다"
	whyNotOnWord    = "앞에 글자가 바로 붙어 있지 않다"
	whyEndsSentence = "문장을 끝내는 부호(. ? ! …) 바로 뒤에 있다"
	whyNoHangulOn   = "뒤에 한글이 바로 붙어 있지 않다"
	whyInWord       = "낱말 가운데에 끼어 있다"
	whyAfterDigit   = "숫자 바로 뒤에 있다"
	whyApart        = "앞뒤가 모두 빈칸이다"
	whyArrow        = "붙임표나 등호와 맞닿아 있다"
	whyBySign       = "다른 기호와 맞닿아 있다"

	whyHangulSides     = "양쪽이 한글이다"
	whyNotHangulSides  = "양쪽이 한글이 아니다"
	whyNotLinked       = "앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다"
	whyLinked          = "앞 어절이 연결 꼴(어/아, 고)로 끝난다"
	whyAsked           = "앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다"
	whyNoBecome        = "바로 뒤 어절이 되다·아니다가 아니다"
	whyAsksNext        = "바로 뒤 어절이 의문 대명사다"
	whyNoBecomeAfter   = "의문 대명사 뒤 어절이 되다·아니다가 아니다"
	whyAsksBare        = "의문 대명사에 이·가가 붙지 않았다"
	whyAsksSelf        = "의문 대명사다"
	whyNoSubjectBefore = "바로 앞 어절이 이·가로 끝나지 않는다"
	whyNotShut         = "바로 앞이 닫는 기호가 아니다"
	whyNoNeun          = "앞 어절이 꾸밈 맺음 는으로 끝나지 않는다"
	whyJoinedParticle  = "줄기와 조사를 합친 글자열이 낱말 표의 낱말이다"
	whyAdnBound        = "꾸밈 맺음 바로 뒤에 붙여 쓴 의존명사 것·거의 줄기다"
	whyModifier        = "꾸밈말과 그 꾸밈을 받는 아는 낱말을 붙여 쓴 통짜 새 낱말이다"
)

// The names of the rules of a sentence, as a trace prints them.
const (
	ruleAuxLinked      = "문장:연결 꼴 뒤 보조 용언"
	ruleMainUnlinked   = "문장:연결 꼴 뒤가 아닌 본용언"
	ruleAdverbialFirst = "문장:에서·서는 부사격 먼저"
	ruleComplement     = "문장:되다·아니다 앞 보격"
	ruleQuoted         = "문장:닫는 따옴표 뒤 인용 조사 라고"
	ruleLimitOne       = "문장:꾸밈 는 뒤 한은 체언"
	ruleShutTopic      = "문장:닫는 기호 뒤 보조사 이란·란"
)

// The rule of a sentence that reads the auxiliary verb from the main
// verb's side, and why it falls.
const (
	ruleLinkBeforeAux = "문장:보조 용언 앞 어/아 꼴"
	whyNoAuxNext      = "바로 뒤 어절의 줄기가 보조 용언이 아니다"
)

// The rule of a sentence of 말다 as an auxiliary verb, and why it falls.
const (
	ruleMalLinked   = "문장:지·고 뒤 보조 동사 말다"
	whyNotMalLinked = "앞 어절이 지·고로 끝나지 않는다"
)

// Why the auxiliary verb's rules fall after -고 unless the stem is 있 or
// 보, and why that stands.
const (
	whyGoNotAux = "앞 어절이 -고로 끝난다"
	whyGoAux    = "줄기가 -고 뒤에 서는 보조 용언(있다·보다)이다"
	whyNoGoAux  = "줄기가 -고 뒤에 서는 보조 용언(있다·보다)이 아니다"
)

// Why the rules of an auxiliary verb stand after a linking closing with
// an auxiliary particle after it (하고는, 알고만), and why that falls.
const (
	whyLinkParticle  = "앞 어절의 표지에 연결 꼴이 있다"
	whyNoParticleEnd = "앞 어절이 보조사로 끝나지 않는다"
)

// Why the auxiliary verb's rules fall after 관하여, 대하여 and 통하여.
const whyRegard = "앞 어절이 관하다·대하다·통하다(…에 관하여, …에 대하여, …을 통하여)다"

// The rule of a sentence of 나다 as an auxiliary verb after -고, and why
// it falls.
const (
	ruleNaLinked   = "문장:-고 뒤 보조 동사 나다"
	whyNotNaLinked = "앞 어절이 -고로 끝나지 않는다"
	whyNaTopic     = "꾸밈 는으로 끝난다(-고 뒤 절의 첫머리 나는과 글자가 같다)"
	whyNaNamed     = "이름꼴이 있다(-고 남은의 남다, 체언 남과 글자가 같다)"
)

// Why the main predicate's rule falls where 보다 is an auxiliary in
// -다(가) 보니 and -다(가) 보면.
const (
	whyWhile          = "앞 어절이 다·다가로 끝난다"
	whyNoWhileClosing = "니·니까·면으로 끝나지 않는다"
	whyNotSee         = "줄기가 보다가 아니다"
)

// Why a slash at a unit's edge is a 빗금 although its sides are no
// Hangul (signGraph).
const whyAtEdge = "단위의 가장자리(어절 끝, 읽는 글자 바로 앞, 홀로)에 있다"

// Why a run of middle dots is not read as 줄임표 (dotsGraph).
const whyNotAtEdge = "단위의 가장자리(어절 끝, 읽는 글자 바로 앞, 홀로)에 있지 않다"

// Why the rule of a sentence gluedShortRule does not stand, and its
// name (its claim is that the word is written onto the word before it).
const (
	whyNotAfterNominal = "바로 앞 어절이 체언으로 끝나지 않는다"
	ruleGluedShort     = "문장:체언에 붙은 한 음절 체언+조사"
)

// The name of the rule of a sentence unmodifiedRule (its claim is that
// the word after the candidate is not one an adnominal may stand
// before).
const ruleUnmodified = "문장:꾸밈받을 말이 뒤에 없는 인·일 체언"

// Why the rules of a sentence that read the copula fall.
const (
	whyNotEnded  = "바로 뒤가 마침표·물음표나 문장의 끝이 아니다"
	whyWordAfter = "바로 뒤에 낱말이 있다"
	whyNoComma   = "바로 뒤가 쉼표가 아니다"
)

// The names of the rules of a sentence that read the copula.
const (
	ruleCopulaEnd   = "문장:문장 끝 야·라고는 맺음"
	ruleNumberQuote = "문장:수 뒤 라고는 인용 조사"
	ruleCopulaComma = "문장:쉼표 앞 이나는 서술격"
)

// The names and the reasons of the rules of a sentence about dependent
// nouns and the modifying forms before them.
const (
	ruleBoundBefore    = "문장:의존 명사 앞 꾸밈"
	ruleNounUnmodified = "문장:꾸밈 뒤가 아닌 수·줄은 체언"
	ruleBoundAfter     = "문장:꾸밈 뒤 의존 명사"
	ruleNounHeaded     = "문장:첫머리나 가운뎃점 앞의 의존 명사 글자는 체언"
	ruleNounModified   = "문장:관형어 뒤 체언"
	ruleNounUnheld     = "문장:관형어나 체언 뒤가 아닌 의존 명사 글자는 체언"

	whyNoBoundNext  = "바로 뒤 어절이 의존 명사가 아니다"
	whyNominalNext  = "바로 뒤 의존 명사가 체언 뒤에도 선다"
	whyKnowNext     = "바로 뒤 어절이 알다·모르다다"
	whyNominalSelf  = "체언 뒤에도 서는 의존 명사다"
	whyDotNext      = "바로 뒤가 가운뎃점이다"
	whyNotHeaded    = "문장 첫머리도 여는 기호 바로 뒤도 아니다"
	whyDenyNext     = "바로 뒤 어절이 아니다다"
	whyNoModifier   = "앞 어절이 꾸밈 맺음이나 관형격 조사로 끝나지 않는다"
	whyAfterNominal = "바로 앞 어절이 체언으로 끝난다"
	whyHeaded       = "문장 첫머리이거나 여는 기호 바로 뒤다"
	whyHeldBefore   = "앞 어절이 접미·어림·닫는 기호로 끝난다"
	whyShutBefore   = "앞 어절이 닫는 기호다"
	ruleDegreeUnit  = "문장:수 바로 뒤 어절 끝 도는 단위 먼저"
)

// The name of the rule of a sentence of a demonstrative standing alone
// (determinerAloneRule), and why it falls.
const (
	ruleDeterminerAlone = "문장:홀로 선 이·그·저는 관형사"
	whyNoModifiableNext = "바로 뒤 어절이 꾸밈을 받을 말(체언·관형사·꾸밈꼴·이름꼴)이 아니다"
	whyGluedWord        = "앞말에 빈칸 없이 붙어 있다"
	whyOpenBefore       = "바로 앞이 여는 기호다"
	whyOpenAfter        = "바로 뒤가 여는 기호다"
)

// The name of the rule of a sentence of a demonstrative determiner that
// is also an adjective's modifying form (determinerFormRule), and why it
// falls.
const (
	ruleDeterminerForm = "문장:홀로 선 이런·그런·저런·어떤·다른은 관형사"
	whyComparedBefore  = "앞 어절에 와·과가 있다(다르다가 받는 견줄 대상, 뒤에 보조사가 붙어도)"
	whyNotDiffer       = "다른이 아니다"
)

// The name of the rule of a sentence of 한 after a complement of 하다
// (haOneRule), and why it falls.
const (
	ruleHaOne    = "문장:하다가 받는 꼴 뒤 한은 하다의 꾸밈꼴"
	whyNoHaTakes = "앞 어절이 하다가 받는 꼴(목적격, 로·으로, 에게·한테·께, -게·-도록·-어야·-려(고)·-고자·-다고 …)로 끝나지 않는다"
)

// Why the exception "written onto the word before it" of the rule of 한
// after a complement of 하다 (haOneRule) falls after the modifying 는.
const whyNeunBefore = "앞 어절이 꾸밈 맺음 는으로 끝난다(체언 한, 限)"

// The name of the rule of a sentence of a word of two kinds (a nominal
// and an adverb) standing alone (adverbAloneRule), and why it falls.
const (
	ruleAdverbAlone   = "문장:홀로 선 체언·부사 글자는 부사"
	whyModifierBefore = "앞 어절이 꾸밈 맺음이나 관형격 조사로 끝난다"
	whyGluedNext      = "바로 뒤 어절이 빈칸 없이 붙어 있다"
)

// The name of the rule of a sentence of a word of two kinds (a nominal
// and a determiner) standing alone (tableDeterminerRule).
const ruleTableDeterminer = "문장:홀로 선 체언·관형사 글자는 관형사"

// The name of the rule of a sentence of a word of two kinds (a nominal
// and an interjection) standing alone (interjectionAloneRule).
const ruleInterjectionAlone = "문장:홀로 선 체언·감탄사 글자는 감탄사"

// The name of the rule of a sentence of a dependent noun of the word
// table's name after a modifying form (boundTableRule).
const ruleBoundTable = "문장:꾸밈 뒤 의존 명사(낱말 표 이름)"

// The name of the rule of a sentence of a dependent noun of the word
// table's name written onto a nominal with its particle (gluedBoundRule).
const ruleGluedBound = "문장:체언에 붙은 의존 명사(낱말 표 이름)+조사"

// The name of the rule of a sentence of a dependent noun of the word
// table's name after a nominal with its particle (nominalBoundRule).
const ruleNominalBound = "문장:체언 뒤 의존 명사(낱말 표 이름)+조사"

// The name of the rule of a sentence of a noun of one syllable with its
// particle after a numeral determiner (countedShortRule), and why it
// falls.
const (
	ruleCountedShort = "문장:수 관형사 뒤 한 음절 체언+조사"
	whyNotCounted    = "바로 앞 어절이 관형사로 읽히지 않는다"
	whyNoNumDet      = "바로 앞 어절이 수 관형사가 아니다"
	whyUnitNext      = "바로 뒤 어절이 수 관형사가 세는 단위다"
	whyNotNumDet     = "어절의 글자가 수 관형사가 아니다"
)

// The name of the rule of a sentence of a modifying form between an
// argument and a word it may modify (argumentAdnRule), and why it falls.
const (
	ruleArgumentAdn     = "문장:격 조사와 꾸밈 받을 말 사이의 꾸밈꼴"
	whyNoArgumentBefore = "앞 어절이 주격·보격·목적격·부사격 조사로 끝나지 않는다"
	whyAdnNext          = "바로 뒤 어절이 꾸밈꼴로 끝난다"
)

// The name of the rule of a sentence of 밖에 with no negative after it
// (outsideNounRule), and why it falls.
const (
	ruleOutsideNoun = "문장:부정이 따르지 않는 밖에는 체언 밖+에"
	whyNegNext      = "바로 뒤 어절이 부정을 나타내는 말이다"
	whyNegNext2     = "그다음 어절이 부정을 나타내는 말이다"
)

// Why boundAfterRule falls for an adverbial dependent noun with 이/가
// (미안한 듯이), and why that exception falls.
const (
	whyAdverbialBound = "주격·보격 조사를 받지 않는 부사성 의존 명사다"
	whyNoSubjectMark  = "주격·보격 이·가가 붙지 않았다"
)

// The name of the rule of a sentence of 달다 after the 어/아 form
// (giveLinkedRule), and why it falls.
const (
	ruleGiveLinked = "문장:어/아 꼴 뒤 보조 동사 달다"
	whyNotConn     = "앞 어절이 어/아 꼴로 끝나지 않는다"
)

// The name of the rule of a sentence of a particle written apart
// (apartParticleRule), and why it falls.
const (
	ruleApartParticle = "문장:따로 쓴 조사는 앞말의 조사"
	whyEndsMarked     = "앞 어절이 조사나 맺음으로 끝난다"
	whyHeadOrSign     = "문장 첫머리이거나 바로 앞이 닫는 기호·영숫자 열·붙임표로 이은 이름·코드다"
)

// Why apartParticleRule does not hold for the vocative particle at the
// head of a sentence (the vocative stands right after the one it calls,
// and nothing stands before it there), and why that exception falls.
const (
	whyVocativeHead = "문장 첫머리의 호격 조사다(부를 말이 앞에 없다)"
	whyNotHead      = "문장 첫머리가 아니다"
	whySignsBefore  = "바로 앞이 이름 없는 기호만의 열이다(가린 이름 ○○)"
)

// Why the exception of apartParticleRule "the word before ends in no
// nominal" falls right after a dash.
const whyDashBefore = "바로 앞이 줄표다"

// The name of the rule of a sentence of 내게, 네게 and 제게
// (dativeContractRule), and why it falls.
const (
	ruleDativeContract = "문장:내게·네게·제게는 대명사와 게"
	whyCauseNext       = "바로 뒤 어절이 하다·되다·만들다다"
)

// The name of the rule of a sentence of the -지 form before 않다,
// 못하다 and 말다 (linkBeforeNegRule), and why it falls.
const (
	ruleLinkBeforeNeg = "문장:않다·못하다·말다 앞 -지"
	whyNoNegNext      = "바로 뒤 어절이 않다·못하다·말다가 아니다"
)

// The name of the rule of a sentence of the -어야 form before 하다 or
// 되다 (linkBeforeHaRule), and why it falls.
const (
	ruleLinkBeforeHa = "문장:하다·되다 앞 -어야"
	whyNoHaNext      = "바로 뒤 어절이 하다·되다가 아니다"
)

// The name of the rule of a sentence of the -기 form with a particle
// before 하다 (namedBeforeHaRule), and why it falls.
const (
	ruleNamedBeforeHa = "문장:하다 앞 -기+조사"
	whyNoDoNext       = "바로 뒤 어절이 하다가 아니다"
)

// The name of the rule of a sentence of the predicate after an object
// (predAfterObjectRule), and why it falls.
const (
	rulePredAfterObject = "문장:목적어 뒤 서술어"
	whyNoObjectBefore   = "앞 어절이 목적격 조사로 끝나지 않는다"
	whyAdnSelf          = "꾸밈꼴로 끝난다(격 조사 뒤 꾸밈꼴 규칙이 판정한다)"
)

// The name of the rule of a sentence of 하다 after -어야 or a -기 form
// (haLinkedRule), and why it falls or stands.
const (
	ruleHaLinked   = "문장:-어야·-기 뒤 하다"
	whyNoHaLink    = "앞 어절이 -어야로 끝나지 않는다"
	whyNamedBefore = "앞 어절에 이름꼴 -기가 있다"
)

// The name of the rule of a sentence of 내, 네 and 제 standing alone
// before a word a modifier may stand before (possessiveRule); it falls
// before a unit for whyUnitNext, as the rules of 한 do.
const rulePossessive = "문장:체언 앞 내·네·제는 의가 줄어든 대명사"

// The name of the rule of a sentence of a word of a nominal and an
// adverb with a nominal written onto it (compoundFrontRule), and why it
// falls.
const (
	ruleCompoundFront = "문장:체언이 붙은 체언·부사 글자는 체언"
	whyApartNext      = "바로 뒤 어절이 빈칸 없이 붙어 있지 않다"
	whyPredNext       = "바로 뒤 어절이 맺음·꾸밈꼴·이름꼴로 끝난다"
)

// The name of the rule of a sentence of a word of a nominal and an
// adverb with the suffix 들 after it (pluralNounRule).
const rulePluralNoun = "문장:접미 들이 붙은 체언·부사 글자는 체언"

// The name of the rule of a sentence of a word of a nominal and an
// adverb or a determiner written onto a nominal (gluedNominalRule), and
// why it falls.
const (
	ruleGluedNominal = "문장:체언에 붙여 쓴 체언·부사·관형사 글자는 체언"
	whyApartBefore   = "앞말에 빈칸 없이 붙어 있지 않다"
)

// The name of the rule of a sentence of 전 after a naming form or a
// quantity (priorNounRule), and why it falls; and why the rule of the
// determiner of the same letters falls there (priorExcept).
const (
	rulePriorNoun   = "문장:이름꼴·수량 뒤 전은 체언"
	whyNoPrior      = "앞 어절이 이름꼴이나 수량(수와 단위)으로 끝나지 않는다"
	whyPriorBefore  = "앞 어절이 이름꼴이나 수량(수와 단위)으로 끝난다(그 뒤의 전은 명사)"
	whyNotPriorStem = "줄기가 전이 아니다"
	whyActionBefore = "앞 어절이 동작성 명사(낱말 표에 그 명사+하다가 있다)다"
	whyMarkedBefore = "앞 어절에 조사나 맺음이 붙었다"
)

// The name of the rule of rebuttal of a sentence of 못하다 after the -지
// form read as the other kind than the predicate before it
// (negKindRule), and why it falls.
const (
	ruleNegKind      = "문장:-지 뒤 못하다는 본용언의 품사"
	whyNotVerbBefore = "앞 어절의 줄기가 동사(낱말 표의 동사, 닫힌 집합의 보조 용언)가 아니다"
	whyNotAdjBefore  = "앞 어절의 줄기가 낱말 표의 형용사가 아니다"
)

// The name of the rule of rebuttal of a sentence of 이라고 after 것 read
// as the particle of direct quotation (thingQuoteRule).
const ruleThingQuote = "문장:것 뒤 이라고는 인용 조사가 아니다"
