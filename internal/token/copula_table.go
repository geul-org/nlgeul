package token

// The copula 이다 (morphemes.md section 2). Its stem 이 after a noun is
// a reading of its own, and after a vowel the stem may be dropped
// before the closings listed here (학교다, 학교고, 학교며, 학교면,
// 학교지만, 학교라서, 학교라면, 학교지, 누군지, 뭔데, 누군가, 뭘까),
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
// 뭐냐, 사과구나, 뭐야, and the quoting 라는, 라며, 라면서, 라니:
// 사과라는): the dictionary's note on 이다 says the stem
// may be dropped after a vowel, mostly not before a modifying or a
// naming form. They are short and common at the end of other words, so
// the dropped stem is read before them only with a word table, after a
// nominal of the table or a dependent noun.
// The naming closings a particle chain is read after are among them
// (nounHeads: 문제기도 하다): before one the dropped stem is read only
// when a particle follows it (dropNaming, and the particle chain right
// after a naming closing).
var droppableKnown = append(append([]string{}, closingSets[setDropKnown]...), nounHeads...)

// dropAdn are the closings before which a dropped stem is rebutted: in
// front of a modifying or a naming form the stem mostly stays (의사인,
// not 의산; 의사이기, not 의사기). forceDropAdn is the rebuttal's force,
// chosen among a few values tried on a thousand sentences.
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
// such a run is read as their stem (abc다).
var droppableRun = []string{"맺음:야"}

const forceDropAdn = 0.5

// notOnly is the one other stem the closings of the copula attach to.
const notOnly = "아니"

// copulaSuffixes are the suffixes the copula stands after (꽃들이다,
// 배씩이다, 너희끼리다). 쯤 and the suffix 이 of a name are not among
// them.
var copulaSuffixes = []string{"접미:들", "접미:씩", "접미:끼리"}
