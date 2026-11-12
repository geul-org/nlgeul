package token

import (
	"github.com/geul-org/nlgeul/data"
)

// Signs that open and close a pair (written with no space inside:
// opening signs onto the word after, closing signs onto the word
// before). A sentence is cut at them as it is read from its first
// letter on, and nothing is looked up ahead to match them: each is
// read for every way the signs before it may stand open, so that what
// stands between them is read on its own and a particle written onto a
// closing sign is read as one (함수(인자)를). pairRows give, sign by
// sign, the opening sign of its kind, whether it opens, whether it
// closes, whether it is unsure (read as its letter beside its readings
// as a sign of a pair), and what keeps it from being such a sign
// (bar…). A quotation mark written the same way at both ends opens
// when none of its kind may stand open, and closes when one may.
var pairRows = data.Must(loadPairs(data.Must(data.Rows("pairs", 4))))

// pairOrder is the order the graphs of the signs are made in.
var pairOrder = data.Must(data.Runes("pairs", 4))

// What keeps a sign from being read as a sign of a pair: it stands
// inside a word (it's, O'Brien); it stands right after a digit
// (5'10"); a space stands on both sides of it, or a hyphen or an
// equals sign right next to it (x < y, ->, =>, <=).
const (
	barWord = 1 << iota
	barDigit
	barApart
)

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
// names of its readings).
const midName = "기호"
