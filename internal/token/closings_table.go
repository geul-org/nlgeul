package token

import (
	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// moreClosings are closings this layer reads on top of the closing
// table (which nlgeul parse reads too and is left as it is). Each is an
// ending the standard dictionary lists: 되, 어야지, 어야만, 어다가, the
// quoting endings with 면서, 며 and 니 (다면서, ㄴ다면서, 는다면서,
// 라면서, 자면서, ㄴ다며, 는다며, 라며, 자며, 다니, ㄴ다니, 는다니,
// 라니), 길래, 든가, ㄴ들, ㄹ망정, ㄹ지라도, ㄹ뿐더러, ㄹ세라,
// 으라고, 다만, ㄴ다만, 는다만, 건만, 느냐고, ㄴ지라,
// 는지라, 은지라, 답니다, ㄴ답니다, 는답니다, 랍니다. Not listed, though
// the dictionary has them: 든 and 는걸, ㄴ걸, 은걸, whose letters are
// read as the particle 든 and as a modifying closing with 거 under 를.
var moreClosings = data.Must(data.Each("closings-more", 4, markerRow))

// gatedClosings are closings of this layer read only with a word
// table, after a predicate's stem of the table: the naming closing 기
// (먹기, 있기), which without that would be read at the end of every
// word that ends in 기. nounHeads are the naming closings a particle
// chain is read right after (있기+는, 받기+도): the word so made is a
// nominal. 음 and ㅁ are not among them. Their names are those of the
// gated closings' entries, and every other list that speaks of a
// naming closing is made from this one (dropNaming, dropAdn,
// droppableKnown, namingEnd).
var gatedClosings = []parse.Marker{
	{Form: "기", Kind: rule.KindNoun, NoYo: true},
}

var nounHeads = closingNames(gatedClosings...)

// looseClosings are the modifying closings 은 and 을 written onto a
// stem that ends in ㄹ without dropping the ㄹ (줄은 for 준, 줄을 for
// 줄), against clause 18 of the orthography (a stem's ㄹ goes before
// ㄴ and ㄹ, and such a stem takes no 으). They are read only with a
// word table, after a predicate's stem of the table that ends in ㄹ,
// next to the candidates the tables give, and are rebutted as a form
// that does not fit what stands before it is (forceClash).
var looseClosings = []parse.Marker{
	{Form: "은", After: rule.AfterLiquid, Kind: rule.KindAdn, NoIda: true, NoYo: true},
	{Form: "을", After: rule.AfterLiquid, Kind: rule.KindAdn, NoIda: true, NoYo: true},
}

const nounAuxName = "사슬:이름 꼴 뒤 조사"

// hangulBound are the closings that are written only onto a Hangul
// syllable: the modifying closings 은, 는 and 을 and the closings 아, 어,
// 나 and 여. After a run of Latin letters and digits, a number, a span
// of code, a path or a file name the same letters are a particle
// (push는, 3은, v2나), and a predicate made of such a run has 하 or the
// copula's stem between the two (push하는, API인, API여서). The
// closings 다, 고, 면, 지 and 네 are not here: before them the copula's
// stem may be dropped after such a run. A closed table: only the
// closings listed are held back.
var hangulBound = closingSets[setHangul]
