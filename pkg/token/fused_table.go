package token

import (
	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// Letters that hold two grammar pieces (decision 014-29, the simple
// ones: one candidate, rare inside words, the stem ends before them).
//
// honorLetters hold the honorific 시 and what is fused under it: 셨
// (시+었), 셔 (시+어…), 십 (시+ㅂ…). Wherever such a letter is read as a
// stem's last syllable 시 with a piece fused into it, it is read as
// the honorific too. After a consonant the link vowel 으 stands before
// it (읽으셨다). honorCount and euCount are the observed counts of 시
// and of the link vowel 으 in the tables; nothing in this repository
// records where they were taken from.
const (
	honorName  = "선어말:시"
	honorCount = 19566
	euName     = "이음:으"
	linkVowel  = '으'
	euCount    = 225409
	negName    = "합침:잖"
	negLetter  = '잖'
	jiName     = "맺음:지"
	yoName     = "보조사:요"
)

// An auxiliary particle is written onto a linking closing (보지도,
// 먹고는, 조용하게는, 먹어서는, 하면서는; clause 41 of the orthography: a particle after an
// ending is written onto it too). auxClosings are the closings and
// auxParticles the particles this is read for. A pair that is itself a
// closing of the tables (어도, 어야, 고도, 지만) is read as that closing
// only. The lists of names are made from the closings' rows
// (closingNames).
const linkAuxName = "사슬:맺음 뒤 보조사"

var auxClosings = closingNames(parse.Marker{Form: "지"}, parse.Marker{Form: "고"}, parse.Marker{Form: "게"}, parse.Marker{Lead: rule.LeadVowel}, parse.Marker{Form: "서", Lead: rule.LeadVowel}, parse.Marker{Form: "면서", Eu: true})
var auxParticles = []string{"보조사:도", "보조사:는", "보조사:만", "보조사:야"}

// A particle is written onto the closings that ask or wonder too
// (맞는지를, 좋은지는, 갈지도, 누구인지가): the clause they close is
// used as a nominal is. askClosings are those closings, ㄴ지, 는지,
// 은지, ㄹ지 and 을지, and askParticles the forms of the particles read
// right after them: the auxiliary particles of auxParticles and the
// case particles 가, 를, 에 and 와. More particles may follow as the
// chain's kinds allow (는지에+는).
const askAuxName = "사슬:물음 맺음 뒤 조사"

var askClosings = closingNames(parse.Marker{Form: "지", Lead: rule.JongN}, parse.Marker{Form: "는지"}, parse.Marker{Form: "은지", After: rule.AfterBatchim}, parse.Marker{Form: "지", Lead: rule.JongL}, parse.Marker{Form: "을지", After: rule.AfterBatchim})
var askParticles = []string{"도", "는", "만", "야", "가", "를", "에", "와"}

var honorLetters = []rune{'셨', '셔', '십'}

// 겠 right after the closing 어야 is 어야 하겠 with 하 left out
// (가야겠다, 먹어야겠네). mustNames are the candidates of that closing,
// guessName the slot, and mustPart what is written out between them:
// the auxiliary 하, which the slot and the closing after it attach to.
const guessName = "선어말:겠"

var mustNames = closingNames(parse.Marker{Form: "야", Lead: rule.LeadVowel})
var mustPart = split{"(하)", "보조:하"}

// unspelt are the syllables a stem ends in after which an 어 is often
// left out in writing, against the spelling rules (clause 35, note 2:
// 되어 is shortened to 돼, so 되도, 되서, 되야, 되요 and 됬다 stand for
// 돼도, 돼서, 돼야, 돼요 and 됐다). Such a syllable is read as 가, 서,
// 내 and 세 are, with an 어/아 fused in that does not show.
var unspelt = []rune{'되'}

// latinBases are the syllables a piece is taken out of even right
// after a letter that is no Hangul: 하 and 되, which make a predicate
// of the word before them, and 이, the copula's stem (push해, merge된,
// api인데).
var latinBases = []string{"하", "되", "이"}

// negParts is what 잖 is written out as: the closing 지 and the stem
// of 않다, which the closing after it attaches to.
var negParts = []split{{"지", jiName}, {"(않)", "부정:않"}}

// splits are closings of the tables that are written out as their
// pieces: 죠 is 지+요, 십시오 the honorific and ㅂ시오, 세요 the
// honorific, 어 and 요, 더라 the slot 더 and 라. A closing so written
// out is the same candidate as its pieces read one by one, and is sent
// up once.
var splits = map[string][]split{
	"맺음:죠":    {{"지", jiName}, {"(요)", yoName}},
	"맺음:세요":   {{"시", honorName}, {"(어)", vowelName}, {"요", yoName}},
	"맺음:으세요":  {{"으", euName}, {"시", honorName}, {"(어)", vowelName}, {"요", yoName}},
	"맺음:더라":   {{"더", "선어말:더"}, {"라", "맺음:라"}},
	"맺음:십시오":  {{"시", honorName}, {"(ㅂ)시오", "맺음:ㅂ시오"}},
	"맺음:으십시오": {{"으", euName}, {"시", honorName}, {"(ㅂ)시오", "맺음:ㅂ시오"}},
}

// contracts are the fused forms whose first word is fixed, a closed
// set: only what is listed is read. A particle put under a pronoun as
// a final consonant is read for the listed pronouns only (누군 here,
// and the one-syllable ones of shortPronouns), not after any word (a
// choice that keeps candidates few: the particles ㄴ and ㄹ do attach to
// any word that ends in a vowel). The last three are the dependent
// noun 거 with a particle under it, written onto a modifying closing
// (한건, 하는게, 할걸): they are read only right after such a closing,
// next to the closings and words of the same letters. 내게, 네게 and
// 제게 are the pronoun forms 내, 네 and 제 with the case particle 게,
// which the dictionary gives after these three and with the meaning of
// 에게: it is read under 에게's entry and nowhere else. 다 and 다가
// right after 어디, 여기, 거기 and 저기 are the particle the dictionary
// gives after a few nominals of place as well as after 에 (여기다
// 놓다, 어디다가 두다): they are read under the entries of 에다 and
// 에다가, after these four words only (after any nominal they would
// look the same as the copula's 다).
// contractLists are the fused forms as their table gives them, set by
// set (contracts.tsv).
var contractLists = data.Must(loadContracts(data.Must(data.Rows("contracts", 7))))

var contracts = contractLists[setContract]

// nounCopulas are dependent nouns with the copula's stem fused into
// their last vowel: 테 is 터 with 이 (할 테니, 갈 텐데, 할 테야 for
// 터이니, 터인데, 터이야). A closed set, read as 게, 건 and 걸 are:
// only right after a modifying closing, written onto it or apart, next
// to the words of the same letters. The closing after it is one the
// copula takes; the stem 이 leaves its trace in the vowel and is
// written (이).
var nounCopulas = contractLists[setNounCopula]

// shortPronouns are one-syllable pronouns with a particle under them
// (난 = 나+는, 날 = 나+를). The letters are common in other words, so
// they are read only with a word table, next to the words of the same
// letters (전, 날, 절).
var shortPronouns = contractLists[setPronoun]

// typoContracts are the forms of the dependent noun 거 with a particle
// under it that are written with 꺼 for 거 (껀 for 건, 껄 for 걸): read
// as 건 and 걸 are, only right after a modifying closing, with the stem
// written wrong (the last column is the letters of it). A closed
// table, read with a word table as 꺼 is (typoKeys). 께 (거 with 이) is
// not in it.
var typoContracts = contractLists[setTypo]

// Auxiliary verbs written onto the 어/아 form of a main verb (clause 47
// of the orthography lets them be). connNames are the candidates of that
// form. The stems read so are those of the closed set of the grammar
// tables (parse.ClosedWords; 지 apart) and, with a word table, auxMore
// (addAux): the boundary before one of them costs nothing, and an 어/아
// that leaves no trace is read right before one. helpers are 지
// (어지다) and 하 (어하다), which are always written onto the main
// verb: they are read as pieces of the same word (지우+(어)+지+(ㄴ)다,
// 예쁘+(어)+하+(ㄴ)다).
var connNames = closingNames(parse.Marker{Lead: rule.LeadVowel})
var auxMore = wordSets[setAuxMore]
var helpers = map[rune]string{'지': "보조:지", '하': "보조:하"}

// yeoBases are the syllables a stem ends in that take 여 for 어/아 and
// show it in the letters (하여, 해): 하. Such a syllable as it stands in
// the letters is not read with an 어/아 fused into it without a trace.
var yeoBases = []string{"하"}

// copulaBars are the auxiliaries written onto a main predicate's 어/아
// form (helpers) that do not follow the copula's stem: 하 of 어하다
// stands after an adjective's (and some verbs') 어/아 form, never
// after 이어 (학생이어하는 is no candidate). A closed table.
var copulaBars = []string{"보조:하"}
