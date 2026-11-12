package token

import "github.com/geul-org/nlgeul/pkg/rule"

// kindOrder is the order kinds are named in; kindLabel their names in
// a graph's name (열쇠:체언).
var kindOrder = []int{kindNominal, kindVerb, kindAdj, kindAdverb, kindDeterminer, kindAlone}
var kindLabel = map[int]string{kindNominal: "체언", kindVerb: "동사", kindAdj: "형용사", kindAdverb: "부사", kindDeterminer: "관형사", kindAlone: "홀로"}

// kindTails says what may follow a stem of each kind: a nominal stands
// alone or takes particles or the copula; a verb stem takes slots and any closing,
// an adjective stem slots and the closings that are not the verbs'
// own, and neither stands alone; an adverb stands alone or takes
// auxiliary particles; a determiner and the other words stand alone.
var kindTails = map[int]int{
	kindNominal:    rule.TailBare | rule.TailAux | rule.TailCase | rule.TailCopula,
	kindVerb:       rule.TailPred | rule.TailVerb,
	kindAdj:        rule.TailPred | rule.TailAdj,
	kindAdverb:     rule.TailBare | rule.TailAux,
	kindDeterminer: rule.TailBare,
	kindAlone:      rule.TailBare,
}

// partKinds folds a word table's parts of speech into kinds; a part of
// speech not listed here stands alone.
var partKinds = map[string]int{
	"명사": kindNominal, "대명사": kindNominal, "수사": kindNominal, "식별자": kindNominal,
	"동사": kindVerb, "보조용언": kindVerb, "형용사": kindAdj,
	"부사": kindAdverb, "관형사": kindDeterminer,
}

// auxSet is the name of the closed set of auxiliary verbs in the
// grammar tables (parse.ClosedWords).
const auxSet = "보조 용언"

// closedKinds gives the kind of each closed set of the grammar tables.
// The auxiliary verbs after an 어/아 form are all verbs (있다 as an
// adjective is a main predicate, which the word table lists).
// closedTails says what may follow a closed-set word when there is no
// word table: a dependent noun stands alone or takes particles or the
// copula; an auxiliary verb's stem takes a predicate tail and does not
// stand alone; a determiner takes no slot or closing, but particles
// are let through (이 and 그 are pronouns too: 그의, 그에게); a
// conjunction stands alone or takes auxiliary particles.
var closedTails = map[string]int{
	"의존명사":  rule.TailBare | rule.TailAux | rule.TailCase | rule.TailCopula,
	auxSet:  rule.TailPred | rule.TailVerb | rule.TailAdj,
	"관형사":   rule.TailBare | rule.TailAux | rule.TailCase | rule.TailCopula,
	"접속 부사": rule.TailBare | rule.TailAux,
}

var closedKinds = map[string]int{"의존명사": kindNominal, auxSet: kindVerb, "관형사": kindDeterminer, "접속 부사": kindAdverb}

// verbClosings are the closings only verbs take: the present ㄴ다/는다
// (also before a quotation's 고, 면, 는), 는구나, 는군, 느냐, 느라; the
// proposal 자, ㅂ시다/읍시다; purpose and intent 러, 려, 려고, 려면,
// 려다, 고자. The modifying 는 and 는데, 는지, 는가 are not here: 있다
// and 없다 take them.
var verbClosings = []string{
	"맺음:ㄴ다", "맺음:ㄴ다고", "맺음:ㄴ다면", "꾸밈:ㄴ다는", "맺음:는다", "맺음:는다고", "맺음:는다면", "꾸밈:는다는",
	"맺음:는구나", "맺음:는군", "맺음:느냐", "맺음:느라", "맺음:느라고",
	"맺음:자", "맺음:자마자", "맺음:자고", "맺음:ㅂ시다", "맺음:읍시다",
	"맺음:러", "맺음:으러", "맺음:려", "맺음:으려", "맺음:려고", "맺음:으려고", "맺음:려면", "맺음:으려면", "맺음:려다", "맺음:으려다", "맺음:고자",
}

// boundAux are the stems of the auxiliaries made of a dependent noun
// and 하다 or 싶다 (듯하다, 만하다, 법하다, 성싶다, 척하다 of the
// spelling rules' article 47, and the dictionary's other auxiliaries
// of that make): each is written right after a modifying closing (올
// 듯하다, 아는 척한다), where no other predicate stands. 직하다 is not
// here: it follows the naming closing ㅁ.
var boundAux = []string{"듯하", "듯싶", "만하", "법하", "뻔하", "성싶", "성부르", "성하", "척하", "체하", "양하"}

// adnPrefix and namingPrefix begin the names of the modifying closings
// (꾸밈:은) and of the naming closings (이름:기, 이름:ㅁ).
const (
	adnPrefix    = "꾸밈:"
	namingPrefix = "이름:"
)

// predTails are the tails a predicate takes: a stem before one of them
// stands as a verb or an adjective.
const predTails = rule.TailPred | rule.TailVerb | rule.TailAdj

// pronouns are the determiners of the closed set that are pronouns
// too (이, 그, 저): with no word table one of them before particles or
// the copula is named a pronoun, not a determiner (그의, 저는); alone
// it stays a determiner.
var pronouns = []string{"이", "그", "저"}

// pronounWords are pronouns this layer knows on top of the word
// table, which nlgeul parse does not read: 네 (너 before 가, and 너의),
// 내 (나 before 가, and 나의) and 제 (저 before 가, and 저의). The
// dictionary lists each as a pronoun. With a word table each is a
// nominal whatever rows the table has for it (네 is there only as a
// determiner), named as the table's nominals are; without one they are
// not read (each would cut every word that begins with it).
var pronounWords = []string{"네", "내", "제"}

const (
	determinerLabel = "관형사"
	pronounLabel    = "대명사"
)
