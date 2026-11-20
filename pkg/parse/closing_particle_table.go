package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
)

// Some particles are written onto a few closings only (clause 41 of
// the orthography: a particle after an ending is written onto it too).
// pickedClosings gives, by the particle's form, the closings it is read
// right after, with every name its form has: 를 after the linking 지 of
// 지 않다 and 지 못하다 (끄지를 못하다, 오지를 않다; the dictionary's 를
// after the linking closing -지), and 와 after the linking 어서 that a
// comparison follows (있어서와 같이; as the gold of the evaluation reads
// it), and the quoting 라고 after the closing 다 (ㄴ다, 는다) of a
// sentence quoted with no quotation marks (있다라고: the dictionary's
// 라고 marks what is quoted as it was said), and then the 의 after the
// linking 어서 of a phrase that modifies a nominal (…에 있어서의 문제: the
// clause 41 of the orthography writes a particle onto an ending too; as
// the gold of the evaluation reads it), and the choosing 든지 and 든가
// after the closing 다 (ㄴ다, 는다) of a sentence quoted with no marks
// (떨어진다든지, 나온다든가: the dictionary's particles 든지 and 든가,
// '-다고 하든지' shortened). Each form has a graph of its own
// (pickedAuxName and the form), whose exception is that the closing
// before it is not one of those. A closed table: only the pairs listed
// are read.
const pickedAuxName = "사슬:맺음 뒤 "

var pickedClosings = map[string][]string{
	"를":  closingNames(Marker{Form: "지"}),
	"와":  closingNames(Marker{Form: "서", Lead: rule.LeadVowel}),
	"라고": closingNames(Marker{Form: "다"}, Marker{Form: "다", Lead: rule.JongN}, Marker{Form: "는다", After: rule.AfterBatchim}),
	"의":  closingNames(Marker{Form: "서", Lead: rule.LeadVowel}),
	"든지": closingNames(Marker{Form: "다"}, Marker{Form: "다", Lead: rule.JongN}, Marker{Form: "는다", After: rule.AfterBatchim}),
	"든가": closingNames(Marker{Form: "다"}, Marker{Form: "다", Lead: rule.JongN}, Marker{Form: "는다", After: rule.AfterBatchim}),
}
