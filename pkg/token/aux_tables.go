package token

import "github.com/geul-org/nlgeul/pkg/parse"

// malStem is the stem of 말다, which the dictionary gives as an
// auxiliary verb after -지 (가지 마라: do not) and after -고 (지고 말았다:
// in the end); the word tables have it under the verb's kind only
// (mainKindOf), as the answer sheet writes it. malLinks are those two
// closings (malLinkedRule).
const malStem = "말"

var malLinks = closingNames(parse.Marker{Form: "지"}, parse.Marker{Form: "고"})

// regardStems are the stems of the verbs that after …에 name what is
// talked or thought about (관하다, the dictionary's "말하거나 생각하는
// 대상으로 하다": …에 관하여, …에 관한). The verb after their -여 form is
// a main verb that takes that phrase, not an auxiliary of theirs (…에
// 관하여 보건대). Only 관하다 is listed, the one the answer sheet shows.
var regardStems = []string{"관하"}

// whileLinks and whileClosings make the frame the dictionary gives 보다
// as an auxiliary verb in after -다(가): -다(가) 보니, -다(가) 보면 (in
// the course of doing it, one finds or comes to what follows: 살다 보니,
// 걷다가 보면). whileLinks are the closings of the word before (다,
// 다가), whileClosings those of 보다 itself (니, 니까, 면; mainUnlinkedRule).
var (
	whileLinks    = closingNames(parse.Marker{Form: "다"}, parse.Marker{Form: "다가"})
	whileClosings = closingNames(parse.Marker{Form: "니", Eu: true}, parse.Marker{Form: "니까", Eu: true}, parse.Marker{Form: "면", Eu: true})
)

// seeAux is the stem of the auxiliary 보다, as the rules of a sentence
// read it after -다(가) (whileLinks: mainUnlinkedRule).
const seeAux = "보"
