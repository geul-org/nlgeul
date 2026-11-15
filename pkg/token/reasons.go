package token

// Why an exception defeats what it attacks, as a trace prints it. Each
// exception carries its reason on its own node (except), so that one
// rule function says what it means where it is put.
const (
	exceptLabel = "예외"
	preferLabel = "선호"

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
	whyJoinedWord   = "줄기와 접미를 합친 글자열이 낱말 표의 낱말이다"
	whyShape        = "기호가 잇는 숫자 마디들이 이 풀이의 꼴이 아니다"
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
	whySpoken       = "닫힌 표의 구어 꼴이다"
	whySole         = "줄기의 낱말이 단위 전체다"
	whyConn         = "본용언의 어/아 꼴 바로 뒤에 붙여 쓴 보조 용언의 줄기다"
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
	whyNotAsked        = "앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝나지 않는다"
	whyAsked           = "앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다"
	whyNoBecome        = "바로 뒤 어절이 되다·아니다가 아니다"
	whyAsksNext        = "바로 뒤 어절이 의문 대명사다"
	whyNoBecomeAfter   = "의문 대명사 뒤 어절이 되다·아니다가 아니다"
	whyAsksBare        = "의문 대명사에 이·가가 붙지 않았다"
	whyAsksSelf        = "의문 대명사다"
	whyNoSubjectBefore = "바로 앞 어절이 이·가로 끝나지 않는다"
	whyNotShut         = "바로 앞이 닫는 기호가 아니다"
	whyHaTakes         = "앞 어절이 하다가 받는 꼴(목적격, 로·으로, -게·-도록·-다고 …)로 끝난다"
	whyNeun            = "앞 어절이 꾸밈 맺음 는으로 끝난다"
	whyNoNeun          = "앞 어절이 꾸밈 맺음 는으로 끝나지 않는다"
	whyJoinedParticle  = "줄기와 조사를 합친 글자열이 낱말 표의 낱말이다"
	whyNotShutAfter    = "바로 뒤가 닫는 기호가 아니다"
)

// The names of the rules of a sentence, as a trace prints them.
const (
	ruleAuxLinked      = "문장:연결 꼴 뒤 보조 용언"
	ruleMainUnlinked   = "문장:연결 꼴 뒤가 아닌 본용언"
	ruleAuxAsked       = "문장:물음 꼴 뒤 보조 용언 보다"
	ruleAdverbialFirst = "문장:에서·서는 부사격 먼저"
	ruleComplement     = "문장:되다·아니다 앞 보격"
	ruleQuoted         = "문장:닫는 따옴표 뒤 인용 조사 라고"
	ruleDeterminerOne  = "문장:홀로 선 한은 관형사"
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

// Why the auxiliary verb's rules fall after 관하여.
const whyRegard = "앞 어절이 관하다(…에 관하여)다"

// The rule of a sentence of 보다 as an auxiliary in -다(가) 보니 and
// -다(가) 보면, and why it and the main predicate's rule fall.
const (
	ruleAuxWhile      = "문장:다·다가 뒤 보조 용언 보니·보면"
	whyNotWhile       = "앞 어절이 다·다가로 끝나지 않는다"
	whyWhile          = "앞 어절이 다·다가로 끝난다"
	whyNoWhileClosing = "니·니까·면으로 끝나지 않는다"
	whyNotSee         = "줄기가 보다가 아니다"
)

// Why a slash at a unit's edge is a 빗금 although its sides are no
// Hangul (signGraph).
const whyAtEdge = "단위의 가장자리(어절 끝, 읽는 글자 바로 앞, 홀로)에 있다"

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
	whyNotEnded    = "바로 뒤가 마침표·물음표나 문장의 끝이 아니다"
	whyWordAfter   = "바로 뒤에 낱말이 있다"
	whyNotNumber   = "줄기가 수가 아니다"
	whyNotContrast = "마지막 조각이 맺음 나가 아니다"
	whyNoComma     = "바로 뒤가 쉼표가 아니다"
)

// The names of the rules of a sentence that read the copula.
const (
	ruleCopulaEnd   = "문장:문장 끝 야·라고는 맺음"
	ruleNumberQuote = "문장:수 뒤 라고는 인용 조사"
	ruleCopulaComma = "문장:쉼표 앞 이나는 서술격"
	ruleAlone       = "문장:괄호 안에 홀로 선 말은 의존명사·관형사가 아니다"
)
