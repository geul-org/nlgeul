package token

import (
	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
)

// moreClosings are closings this layer reads on top of the closing
// table (which kopas parse reads too and is left as it is). Each is an
// ending the standard dictionary lists: 되, 어야지, 어야만, 어다가, the
// quoting endings with 면서, 며 and 니 (다면서, ㄴ다면서, 는다면서,
// 라면서, 자면서, ㄴ다며, 는다며, 라며, 자며, 다니, ㄴ다니, 는다니,
// 라니), 길래, 든가, ㄴ들, ㄹ망정, ㄹ지라도, ㄹ뿐더러, ㄹ세라,
// 으라고, 다만, ㄴ다만, 는다만, 건만, 느냐고, ㄴ지라,
// 는지라, 은지라, 답니다, ㄴ답니다, 는답니다, 랍니다. Not listed, though
// the dictionary has them: 든 and 는걸, ㄴ걸, 은걸, whose letters are
// read as the particle 든 and as a modifying closing with 거 under 를.
var moreClosings = []parse.Marker{
	{Form: "되", Kind: rule.KindLink},
	{Form: "야지", Lead: rule.LeadVowel, Kind: rule.KindFinal},
	{Form: "야만", Lead: rule.LeadVowel, Kind: rule.KindLink},
	{Form: "다가", Lead: rule.LeadVowel, Kind: rule.KindLink},
	{Form: "다면서", Kind: rule.KindFinal | rule.KindLink},
	{Form: "다면서", Lead: rule.JongN, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "는다면서", After: rule.AfterBatchim, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "라면서", Eu: true, Kind: rule.KindFinal | rule.KindLink},
	{Form: "자면서", Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "다며", Lead: rule.JongN, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "는다며", After: rule.AfterBatchim, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "라며", Eu: true, Kind: rule.KindFinal | rule.KindLink},
	{Form: "자며", Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "다니", Kind: rule.KindFinal | rule.KindLink},
	{Form: "다니", Lead: rule.JongN, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "는다니", After: rule.AfterBatchim, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "라니", Eu: true, Kind: rule.KindFinal | rule.KindLink},
	{Form: "길래", Kind: rule.KindLink},
	{Form: "든가", Kind: rule.KindLink},
	{Form: "들", Lead: rule.JongN, Kind: rule.KindLink},
	{Form: "은들", After: rule.AfterBatchim, Kind: rule.KindLink},
	{Form: "망정", Lead: rule.JongL, Kind: rule.KindLink},
	{Form: "을망정", After: rule.AfterBatchim, Kind: rule.KindLink},
	{Form: "지라도", Lead: rule.JongL, Kind: rule.KindLink},
	{Form: "을지라도", After: rule.AfterBatchim, Kind: rule.KindLink},
	{Form: "뿐더러", Lead: rule.JongL, Kind: rule.KindLink},
	{Form: "을뿐더러", After: rule.AfterBatchim, Kind: rule.KindLink},
	{Form: "세라", Lead: rule.JongL, Kind: rule.KindLink},
	{Form: "을세라", After: rule.AfterBatchim, Kind: rule.KindLink},
	{Form: "으라고", After: rule.AfterBatchim, Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "다만", Kind: rule.KindLink},
	{Form: "다만", Lead: rule.JongN, Kind: rule.KindLink, NoIda: true},
	{Form: "는다만", After: rule.AfterBatchim, Kind: rule.KindLink, NoIda: true},
	{Form: "건만", Kind: rule.KindLink},
	{Form: "느냐고", Kind: rule.KindFinal | rule.KindLink, NoIda: true},
	{Form: "지라", Lead: rule.JongN, Kind: rule.KindLink},
	{Form: "는지라", Kind: rule.KindLink, NoIda: true},
	{Form: "은지라", After: rule.AfterBatchim, Kind: rule.KindLink, NoIda: true},
	{Form: "답니다", Kind: rule.KindFinal, NoYo: true},
	{Form: "답니다", Lead: rule.JongN, Kind: rule.KindFinal, NoYo: true, NoIda: true},
	{Form: "는답니다", After: rule.AfterBatchim, Kind: rule.KindFinal, NoYo: true, NoIda: true},
	{Form: "랍니다", Eu: true, Kind: rule.KindFinal, NoYo: true},
}

// gatedClosings are closings of this layer read only with a word
// table, after a predicate's stem of the table: the naming closing 기
// (먹기, 있기), which without that would be read at the end of every
// word that ends in 기. nounHeads are the naming closings a particle
// chain is read right after (있기+는, 받기+도): the word so made is a
// nominal. 음 and ㅁ are not among them.
var gatedClosings = []parse.Marker{
	{Form: "기", Kind: rule.KindNoun, NoYo: true},
}

var nounHeads = []string{"이름:기"}

const (
	nounAuxName = "사슬:이름 꼴 뒤 조사"
	nounLetter  = '기'
)
