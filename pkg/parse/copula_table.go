package parse

// The copula 이다 (morphemes.md section 2). Its stem 이 after a noun is
// a candidate of its own, and after a vowel the stem may be dropped
// before the closings listed here (학교다, 학교고, 학교며, 학교면,
// 학교지만, 학교라서, 학교라면, 학교지, 누군지, 뭔데, 누군가, 뭘까,
// 학자리라, and with a word table the 하오체 closing 올시다, which
// follows only the copula's stem and 아니: 학교올시다),
// and before 잖 (학교잖아: 지 and 않). copulaCount is its observed
// count in the tables; nothing in this repository records where that
// count was taken from.
const (
	copulaName  = "서술격:이"
	droppedName = "서술격:이(빠짐)"
	copulaCount = 4520000
)

var droppable = closingSets[setDrop]

// bareAdn is the modifying closing ㄴ put under the syllable before it.
const bareAdn = "꾸밈:ㄴ"

// droppableKnown are more closings the stem is dropped before (거네,
// 뭐냐, 사과구나, 뭐야, and the quoting 라고, 라는, 라며, 라면서, 라니:
// 사과라는, 사과라고 beside the particle 라고, and the quoting 냐는, 냐며, 냐면, 냐면서, 냐니까: 누구냐면):
// the dictionary's note on 이다 says the stem
// may be dropped after a vowel, mostly not before a modifying or a
// naming form. They are short and common at the end of other words, so
// the dropped stem is read before them only with a word table, after a
// nominal of the table or a dependent noun.
// The gated naming closing 기 is among them (nounHeads: 문제기도 하다;
// not 음 and ㅁ, which a particle chain is read after as well): before
// it the dropped stem is read only when a particle follows it
// (dropNaming, and the particle chain right after a naming closing).
// The set also names the slots the stem is dropped before, after a
// nominal of the table as well (누구+(이)+시+(ㅂ)니까, 사내+(이)+겠+느냐,
// 증세+(이)+더+냐), and the closings whose first piece is the slot 시
// (세요, 시오: 누구+(이)+시+(어)+요); before 었 the stem is not dropped
// but fused (누구였다).
var droppableKnown = append(append([]string{}, closingSets[setDropKnown]...), nounHeads...)

// droppableNew are the closings of droppableKnown before which the
// dropped stem is read after a word no table lists as well: the quoting
// 라 (-라고 shortened), 라는, 라며, 라면서 and 라니, which name what a
// nominal is called (카담바리라는 책: 카담바리+(이)+라는). The tables know
// no kind of a new word, so it may be a nominal, and these closings
// follow the nominal a quotation names; before the other closings of
// droppableKnown (거네, 고래네) the word before is still asked to be a
// nominal of the table.
var droppableNew = closingSets[setDropNew]

// dropAdn are the closings before which a dropped stem is rebutted: in
// front of a modifying or a naming form the stem mostly stays (의사인,
// not 의산; 의사이기, not 의사기), at the force forceDropAdn.
var dropAdn = append([]string{bareAdn}, nounHeads...)

// dropNaming are the naming closings before which a dropped stem does
// not stand, unless a particle is written right after the closing
// (문제기도, 문제기는: not 문제기 by itself).
var dropNaming = nounHeads

// droppableRun are the closings a dropped stem is read before when the
// letter before it is no Hangul syllable (a run of Latin letters and
// digits, a number, a span of code, a file name: abc야): nothing in
// such letters tells whether they end in a vowel. 야 attaches only to
// the copula, so nothing else is read there; before the other closings
// such a run is read as their stem (abc다), but for a run that ends in
// a digit, which no closing stands right on (stemlessAfter: 3면).
var droppableRun = []string{"맺음:야"}

// notOnly is the one other stem the closings of the copula attach to.
const notOnly = "아니"

// copulaSuffixes are the suffixes the copula stands after (꽃들이다,
// 배씩이다, 너희끼리다, and with a word table 보장률이다, 수리율이다, 보장사이다). 쯤 and the
// suffix 이 of a name are not among them.
var copulaSuffixes = []string{"접미:들", "접미:씩", "접미:끼리", "접미:률", "접미:율", "접미:사"}

// copulaParticles are the particles the copula stands after: the
// range particle 대로 (뜻대로지, 마음대로다) and the range particles 까지
// and 부터 (오늘까지이다, 내일부터이다), as the gold of the evaluation
// reads them: school grammar puts the copula after a phrase that stands
// as a nominal, a particle written onto it included. The dictionary's
// 이다 lists no particle before it. Then 뿐 written onto a nominal
// (생각뿐이다, 사람뿐이었다): the dictionary's particle 뿐 (an auxiliary
// particle after a nominal; the dependent noun 뿐 stands after a
// modifying or a final ending), as the gold of the evaluation now reads
// it after a nominal. copulaBefore are the last pieces of a chain the
// copula follows.
var copulaParticles = []string{"부사격:대로", "보조사:까지", "보조사:부터", "보조사:뿐"}

var copulaBefore = append(append([]string{}, copulaSuffixes...), copulaParticles...)
