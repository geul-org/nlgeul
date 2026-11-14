package token

import (
	"slices"

	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/pkg/parse"
)

// The words of the spacing rules (spacing-words.tsv, written for this
// layer; the rules are named by the number of their article of the
// spelling rules and by nothing else). Each list is a closed set: a
// word is added by writing a line.
var spacingSets = data.Must(data.Sets("spacing-words", setLink, setTitle, setSurname, setBoundNoun, setBoundAdn, setAuxApart))

var (
	linkWords      = spacingSets[setLink]
	titleWords     = spacingSets[setTitle]
	surnameWords   = spacingSets[setSurname]
	afterNounWords = spacingSets[setBoundNoun]
)

// dependentWords are the dependent nouns the rule of article 42 marks:
// the closed set of the grammar tables and the two sets of the table
// of this layer's spacing words. auxiliaryWords are the stems of the
// auxiliaries article 47 marks: the closed set of the grammar tables,
// the ones written onto the main verb (auxMore) and the spacing words'.
var (
	dependentWords = slices.Concat(parse.ClosedWords()[boundSet], afterNounWords, spacingSets[setBoundAdn])
	auxiliaryWords = slices.Concat(parse.ClosedWords()[auxSet], auxMore, spacingSets[setAuxApart])
)

// particleKinds are the kinds of the pieces that are particles; the
// copula (서술격) is not one of them. hangulUnits are the units of the
// unit table (units.tsv) whose letters are Hangul syllables.
var particleKinds = []string{"주격", "보격", "목적격", "관형격", "부사격", "호격", "보조사"}

var hangulUnits = hangulWords(unitWords)

// SpacingArticles are the names of the spacing rules, by the number of
// their article of the spelling rules (다만 for its proviso), in the
// order they are tried. Options.Spacing is a list of them.
var SpacingArticles = []string{"2", "41", "42", "43", "43다만", "44", "45", "46", "47", "47다만", "48", "48다만", "49", "50"}

// spacingRules judge a boundary of a candidate by the two words that
// meet there, in the order of SpacingArticles.
var spacingRules = []spacingRule{
	{"2", fitsAny},
	{"41", fitsParticle},
	{"42", fitsBound},
	{"43", fitsUnit},
	{"43다만", fitsUnitDigit},
	{"44", fitsMyriad},
	{"45", fitsLink},
	{"46", fitsShort},
	{"47", fitsAux},
	{"47다만", fitsAuxApart},
	{"48", fitsTitle},
	{"48다만", fitsSurname},
	{"49", fitsNames},
	{"50", fitsNames},
}
