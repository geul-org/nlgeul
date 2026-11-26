package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

// The signs that are named where they stand at the edges of a unit or
// alone (the full stop has its own graphs: stopNames), each with its
// candidates in the order they are printed. A candidate is a graph of its
// own (signGraph). A sign with no candidate that stands, and a sign that
// is not in the table, is sent up as its letters with no name: no
// letter of a sentence is dropped.
var signNames = data.Must(loadSigns(data.Must(data.Rows("signs", 2))))
var signOrderNamed = data.Must(data.Runes("signs", 2))

// roleNames are the candidates of a sign by what it does for the words
// after it (a heading, a topic, the name of what is meant) and not as
// the sign itself. When only such candidates stand, the sign's letters
// with no name are sent up beside them. headingName is printed with
// the number of the signs in a row and no colon (## is 제목2).
var roleNames = []string{headingName, "태그", "명칭"}

const headingName = "제목"

// bareSignName is the name of a sign that is told only to be a sign,
// with no name of its own (the backslash: \[기호]).
const bareSignName = "기호"

// arrowSigns stand right next to < or > in a run that is no bracket
// (->, =>, <=, >=).
const arrowSigns = "-="

// stopSigns are the signs that close the word before them inside a
// unit, as a piece of that word (stopAt): the full stop, the comma (and
// its full-width form ，) and the ellipsis (…, the same sign as ...).
const stopSigns = ".,…，"

// apartSigns are the signs that, between letters of a unit, are also read
// as a word of their own (a candidate added beside the one that keeps
// them in the word): the middle dots and their variants (‧ ․ •), the
// slash, the tilde and its variants (∼ ～ 〜), the dashes (― —) and the
// colon (when its sides are not both Latin letters or digits: EPI:입자)
// and the signs of bareApart;
// the slash and the tildes only with Hangul on both sides (their graphs).
// cutSigns are the signs that cut a sentence into units where they stand
// between letters (cutter): the plus and its full-width form (＋).
const (
	apartSigns = "·ㆍ・‧․•/" + tildeSigns + "―—:" + plainDashes + bareApart
	cutSigns   = "+＋"
)

// bareApart are the signs with no name that are also read as a word of
// their own between letters of a unit (apartGraph): the equals sign with
// a letter that is no Latin letter or digit on one side (값=3 is 값, =
// and 3; a=1 stays one name). They are printed as their letters with no
// name, as they are at a unit's edge. apartName names their graphs
// (글자 사이:=).
const (
	bareApart = "="
	apartName = "글자 사이:"
)

// plainDashes are the signs written for a dash that are drawn with too:
// the hyphen, one or a row of any length (a typescript's -- and ---),
// and the box-drawing line ─ (U+2500). A row of them is also a rule
// across a page, the line of a table, a frame, the shaft of an arrow,
// and one hyphen a list's mark, a minus or the sign that binds two
// words; so each has a graph that reads its place (dashGraph), and a
// place that speaks against the dash leaves the letters with no name.
// Between letters of a unit a row of two hyphens or more and the
// box-drawing line are also read as a word of their own (apartSigns);
// one hyphen there is not (it binds words: 원-달러).
const plainDashes = "-─"

// minusName is the candidate of one hyphen alone as the minus of a sum
// (가 - 나), printed as the minus between two numbers is (10-20[빼기]).
const minusName = "빼기"

// textMarks are the punctuation marks of a sentence that may stand
// right next to a dash (다---. and ─.” and (密告狀)——?): the marks that
// end or pause a sentence and the signs of a pair but < and >, which
// with a hyphen draw an arrow. Any other sign next to a row of plain
// dashes draws something with it (drawn).
const textMarks = ".,?!…，:;·\"'“”‘’()[]{}「」『』《》〈〉｢｣【】〔〕（）"

// sumBrackets are the opening brackets a sum is written in (가 × [나 -
// 다], (= 가 - 나), {가 - 나}): one hyphen alone inside one that stands
// open is read as a minus before a dash (sumPlace).
const sumBrackets = "([{（［｛"
