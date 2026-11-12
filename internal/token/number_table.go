package token

import (
	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// Numbers (decision 014-31). A run of digits is a number; a sign
// between two runs of digits is judged by a graph for each of its
// readings, and a reading that stands joins the runs into one number.
// numName is what a number is called in the output; numGraph names the
// graph of a run of digits.
const (
	numName  = "수"
	numGraph = "수:숫자열"
)

// hexName and binForm name the other ways a run reads as a number: a
// run that is 0x and then hexadecimal digits is the number named
// 수:16진법 (the only reading of it); a number of digits 0 and 1 only
// is read a second time as 수:2진법. The prefix is the lower case 0x
// only.
const (
	hexName  = "수:16진법"
	binForm  = "2진법"
	hexLead  = "0x"
	hexDigit = "0123456789abcdefABCDEF"
)

// signShapes give, sign by sign, the readings that join runs of digits
// and the shape each asks of them, in the order they are named.
// 자릿수 is shown as a number (shownAs).
var signShapes = map[rune][]rule.Shape{
	',': {{Of: "자릿수:,", Min: 2, Lens: [][2]int{{1, 3}, {3, 3}}}},
	'~': {{Of: "범위:~", Min: 2, Max: 2}},
	'-': {
		{Of: "범위:-", Min: 2, Max: 2},
		{Of: "빼기:-", Min: 2, Max: 2},
		{Of: "날짜:-", Min: 3, Max: 3, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}, {1, 31}}},
		{Of: "전화:-", Min: 3, Max: 3, Lens: [][2]int{{2, 4}, {3, 4}, {4, 4}}, Head: "0"},
	},
	':': {
		{Of: "시각::", Min: 2, Max: 3, Lens: [][2]int{{1, 2}, {2, 2}}, Vals: [][2]int{{0, 24}, {0, 59}}},
		{Of: "비율::", Min: 2, Max: 2},
	},
	'/': {
		{Of: "날짜:/", Min: 2, Max: 2, Lens: [][2]int{{1, 2}}, Vals: [][2]int{{1, 12}, {1, 31}}},
		{Of: "나눔:/", Min: 2, Max: 2},
	},
	'.': {
		{Of: "날짜:.", Min: 3, Max: 3, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}, {1, 31}}},
		{Of: "IP:.", Min: 4, Max: 4, Lens: [][2]int{{1, 3}}, Vals: [][2]int{{0, 255}}},
	},
}

// nameSigns are the signs that join two runs of Latin letters and
// digits into one name when they are not both numbers: a slash joins
// the parts of a path (a/b), a hyphen two words that belong together
// (h-Cat).
var nameSigns = map[rune]string{'/': "경로:/", '-': "붙임:-"}
var nameOrder = []rune{'/', '-'}

var signOrder = []rune{',', '~', '-', ':', '/', '.'}

var shownAs = map[string]string{"자릿수": numName}

// headSigns stand at the very front of a number: a sign, or money.
var headSigns = data.Must(loadHeads(data.Must(data.Rows("number-heads", 2))))

var headOrder = data.Must(data.Runes("number-heads", 2))

// unitWords are read as units only right after a number (대, 분, 일,
// 시, 자 and 배 are common inside words). The table is closed. A longer
// unit and a shorter one that ends it are both read where they fit
// (3개월 is 개월; 개 alone leaves 월 to the next word). unitWords are
// the letters a unit is written with, in the order of the table's lines
// (units.tsv), and unitNames what each is named with after unitLabel
// (the table's second column).
var unitWords = data.Must(data.Column("units", 2))
var unitNames = data.Must(data.Map("units"))

// Places after a number: a greater number (만, 천), 여 (more than), a
// unit, 째 (an order).
const (
	placeBig = iota + 1
	placeOver
	placeUnit
	placeOrder
)

const (
	unitLabel = "단위:"
	overName  = "어림:여"
	orderName = "차례:째"
	tildeName = "범위:~(단위 뒤)"
)

// afterOthers is how many pieces besides the unit words are counted
// among the entries a piece after a number is taken as one of: 여 and
// 째. The three unitSuffixes are not counted; nothing records why (the
// count was set before they were added and left as it was).
const afterOthers = 2

// unitSuffixes stand after a number and its unit, where 째 stands
// (백 원짜리, 다섯 개가량, 세 명당): suffixes, as 째 and 여 are.
var unitSuffixes = data.Must(data.Column("unit-suffixes", 1))
