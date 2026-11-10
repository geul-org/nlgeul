package parse

// printed is the order in which equal verdicts are printed: the order
// school grammar lists the constituents in.
var printed = []string{labelSubj, labelObj, labelComp, labelPred, labelAdn, labelAdv, labelVoc}

// conclusions names the constituent each claim concludes, by the claim
// function's name.
var conclusions = map[string]string{
	"subjectCase":      labelSubj,
	"topicSubject":     labelSubj,
	"objectCase":       labelObj,
	"topicObject":      labelObj,
	"complementCase":   labelComp,
	"closingPredicate": labelPred,
	"lastPredicate":    labelPred,
	"adnominalCase":    labelAdn,
	"adnominalClosing": labelAdn,
	"adverbialCase":    labelAdv,
	"adverbialClosing": labelAdv,
	"rangeAdverbial":   labelAdv,
	"determinerWord":   labelAdn,
	"conjunctiveWord":  labelVoc,
	"vocativeCase":     labelVoc,
}

// adverbialClosings is the criterion of adverbialClosing: 게, which
// morphemes.md section 4 says is used like an adverb, and its emphatic
// 게끔.
var adverbialClosings = []string{"게", "게끔"}

// chainClaims and tailClaims are the claims that rest on the particle
// chain and on a predicate tail; wordClaims rest on the word as a whole.
var chainClaims = []string{"subjectCase", "topicSubject", "objectCase", "topicObject", "complementCase", "adnominalCase", "adverbialCase", "vocativeCase", "rangeAdverbial"}
var tailClaims = []string{"closingPredicate", "lastPredicate", "adnominalClosing", "adverbialClosing"}
var wordClaims = []string{"determinerWord", "conjunctiveWord"}
