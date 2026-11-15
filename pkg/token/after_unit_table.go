package token

import "github.com/geul-org/nlgeul/pkg/data"

// afterUnitWords are the words written right after a number's unit as
// words of their own (after-unit-words.tsv): the suffixes -간 (間, for a
// time) and -경 (頃, about that time) and the nouns 전 (前), 후 (後) and
// 차 (差, a difference), as the standard dictionary has them after a
// number and its unit (이틀간, 1시경, 3일 전, 3점 차). A known word whose
// letters are a unit and one of these, right after a number (5년간,
// 6월경, 2일전), is rebutted (unitWordAt, unitWordRebuttal), so that the
// number keeps its unit (한글 맞춤법 제43항: a unit is written with its
// number).
var afterUnitWords = data.Must(data.Column("after-unit-words", 1))
