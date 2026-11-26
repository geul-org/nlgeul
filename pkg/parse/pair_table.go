package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

// Signs that open and close a pair (written with no space inside:
// opening signs onto the word after, closing signs onto the word
// before). A sentence is cut at them as it is read from its first
// letter on, and nothing is looked up ahead to match them: each is
// read for every way the signs before it may stand open, so that what
// stands between them is read on its own and a particle written onto a
// closing sign is read as one (함수(인자)를). pairRows give, sign by
// sign, the opening sign of its kind, whether it opens, whether it
// closes, whether it is unsure (read as its letter beside its candidates
// as a sign of a pair), and what keeps it from being such a sign
// (bar…). A quotation mark written the same way at both ends opens
// when none of its kind may stand open, and closes when one may.
var pairRows = data.Must(loadPairs(data.Must(data.Rows("pairs", 4))))

// pairOrder is the order the graphs of the signs are made in.
var pairOrder = data.Must(data.Runes("pairs", 4))

// What keeps a sign from being read as a sign of a pair: it stands
// inside a word (it's, O'Brien); a space stands on both sides of it,
// or a hyphen or an equals sign right next to it (x < y, ->, =>, <=).
const (
	barWord = 1 << iota
	barApart
)

// loneShuts are the quotation marks that close a pair with none of
// their kind open (written onto the word before, with Hangul or no
// letter at all right after: 좋다"고, 있다”며, 했다' 로; or before a
// marker that takes a quotation: "라고): when no way the signs before them may
// stand open holds one of their kind, they may still close, and what
// stands open is left as it is. Brackets are not among them (1), 가)).
const loneShuts = "\"'”’"

// quoteShuts are the signs that close a quotation: the quotation marks
// of loneShuts and the closing corner brackets. An ending of the quoting
// verb may be written right onto one of them (quotedClosings).
const quoteShuts = loneShuts + "」』｣"

// quoteSigns are the quotation marks among the signs of a pair: a reading
// of one is lowered where its place speaks against it (pairGraph).
const quoteSigns = "\"'“”‘’"

// swapOpens are the round quotation marks drawn as closing ones that are
// also read as opening where an opening mark is written (a space or the
// sentence's start before them, a letter or digit right after: ”두, ’거래),
// and swapShuts those drawn as opening ones that are also read as closing
// where a closing mark is written (사항‘이라고, 좋겠다“고): a text may
// write the two shapes the wrong way round.
const (
	swapOpens = "”’"
	swapShuts = "“‘"
)

// quoteMarkers are the markers that take a quotation, written right onto
// its closing quotation mark (quote-markers.tsv): "…"라고, "…"고, '…'는.
var quoteMarkers = data.Must(data.Column("quote-markers", 1))

// openName and shutName are what a sign of a pair is called, with its
// depth after a colon (열기:1, 닫기:2). pairName stands for either in a
// unit whose head is such a sign, while the unit is read; codeName is
// what a span of code between backticks is called.
const (
	openName = "열기"
	shutName = "닫기"
	pairName = "묶음"
	codeName = "코드"
)

// midName is what a sign between letters of a unit is called while the
// unit is read, before the stretches are made (expandMid gives it the
// names of its candidates).
const midName = "기호"
