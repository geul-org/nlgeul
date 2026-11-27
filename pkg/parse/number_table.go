package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// Numbers (decision 014-31). A run of digits is a number; a sign
// between two runs of digits is judged by a graph for each of its
// candidates, and a candidate that stands joins the runs into one number.
// numName is what a number is called in the output; numGraph names the
// graph of a run of digits.
const (
	numName  = "수"
	numGraph = "수:숫자열"
)

// The forms of a number's name by the letters it is written in
// (numberNameOf): digits, Hangul syllables, Chinese characters, Roman
// numeral letters. Any other letters (a circled number, one Latin letter)
// keep the name with no form.
const (
	decimalForm = "십진"
	hangulForm  = "한글십진"
	hanjaForm   = "한자십진"
	romanForm   = "로마십진"
)

// numberNames are the names a number is given (numberNameOf).
var numberNames = []string{numName, numName + ":" + decimalForm, numName + ":" + hangulForm, numName + ":" + hanjaForm, numName + ":" + romanForm}

// loneNumberGraph names the graph of a number that is no run of digits
// read where nothing in its word tells that it is a number (no unit is
// written onto it: 사십 before 년 written apart): loneNumbers.
const loneNumberGraph = "수(홀로 섬)"

// hexName and binForm name the other ways a run reads as a number: a
// run that is 0x and then hexadecimal digits is the number named
// 수:16진법 (the only candidate of it); a number of digits 0 and 1 only
// is read a second time as 수:2진법. The prefix is the lower case 0x
// only.
const (
	hexName  = "수:16진법"
	binForm  = "2진법"
	hexLead  = "0x"
	hexDigit = "0123456789abcdefABCDEF"
)

// signShapes give, sign by sign, the candidates that join runs of digits
// and the shape each asks of them, in the order they are named (the
// tilde and its variants, a range each, are added by tildeShapes); of the
// same letters read two ways the one named first comes first. Numbers
// joined by a slash are a fraction before a date: the orthography's
// appendix on punctuation gives the slash between numbers as the way to
// write a fraction (빗금 (3): 3/4분기, 3/20) and no date written so.
// A ratio joins two numbers or more with colons, with no bound on how
// many (57:23:20; of three numbers that fit a time the time is named
// first: 10:30:15).
// 자릿수 is shown as a number (shownAs).
var signShapes = tildeShapes(map[rune][]rule.Shape{
	',': {{Of: "자릿수:,", Min: 2, Lens: [][2]int{{1, 3}, {3, 3}}}},
	'–': {
		{Of: "범위:–", Min: 2, Max: 2},
		{Of: "빼기:–", Min: 2, Max: 2},
		{Of: "날짜:–", Min: 3, Max: 3, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}, {1, 31}}},
		{Of: "전화:–", Min: 3, Max: 3, Lens: [][2]int{{2, 4}, {3, 4}, {4, 4}}, Head: "0"},
	},
	'-': {
		{Of: "범위:-", Min: 2, Max: 2},
		{Of: "빼기:-", Min: 2, Max: 2},
		{Of: "날짜:-", Min: 3, Max: 3, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}, {1, 31}}},
		{Of: "전화:-", Min: 3, Max: 3, Lens: [][2]int{{2, 4}, {3, 4}, {4, 4}}, Head: "0"},
	},
	':': {
		{Of: "시각::", Min: 2, Max: 3, Lens: [][2]int{{1, 2}, {2, 2}}, Vals: [][2]int{{0, 24}, {0, 59}}},
		{Of: "비율::", Min: 2},
	},
	'/': {
		{Of: "나눔:/", Min: 2, Max: 2},
		{Of: "날짜:/", Min: 2, Max: 2, Lens: [][2]int{{1, 2}}, Vals: [][2]int{{1, 12}, {1, 31}}},
	},
	'.': {
		{Of: "날짜:.", Min: 3, Max: 3, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}, {1, 31}}},
		{Of: "IP:.", Min: 4, Max: 4, Lens: [][2]int{{1, 3}}, Vals: [][2]int{{0, 255}}},
	},
})

// stopShapes are the shapes of a date of two numbers joined by a full
// stop, read only where a full stop stands right after the second
// number (the orthography's appendix on punctuation, 마침표 (3): it is
// written after the numbers of a date, 1919. 3. 1., 10. 1.~10. 12.):
// a month and a day (1.28.) and a year and a month (1961.12.). They
// come before the full stop's own joins: two numbers with a full stop
// after them are a date before a decimal. Without the full stop after
// them the date stands behind the decimal (plainShapes: 2.5, 3.1이닝).
var stopShapes = []rule.Shape{
	{Of: "날짜:.", Min: 2, Max: 2, Lens: [][2]int{{1, 2}}, Vals: [][2]int{{1, 12}, {1, 31}}},
	{Of: "날짜:.", Min: 2, Max: 2, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}}},
}

// plainShapes are the shapes of a date and a time of two numbers joined
// by a full stop with no full stop asked after them: a month and a day
// (5.18, the name of a day: 4.3), a year and a month (1963.7), an hour
// and a minute (19.00). They come after the full stop's own joins: of
// the same letters the decimal and the version stay first, and the date
// and the time stand behind them as candidates. The orthography's
// appendix on punctuation writes a full stop between the numbers of a
// month and a day that name a day of a special meaning (3.1 운동, 8.15
// 광복); a time written so is no use of the norm (it parts the hour and
// the minute with a colon) but is met in writing. The ranges are those
// of the shapes of the other signs (날짜:/, 시각::).
var plainShapes = []rule.Shape{
	{Of: "날짜:.", Min: 2, Max: 2, Lens: [][2]int{{1, 2}}, Vals: [][2]int{{1, 12}, {1, 31}}},
	{Of: "날짜:.", Min: 2, Max: 2, Lens: [][2]int{{4, 4}, {1, 2}}, Vals: [][2]int{{1, 9999}, {1, 12}}},
	{Of: "시각:.", Min: 2, Max: 2, Lens: [][2]int{{1, 2}, {2, 2}}, Vals: [][2]int{{0, 24}, {0, 59}}},
}

// nameSigns are the signs that join two runs of Latin letters and
// digits into one name when they are not both numbers: a slash joins
// the parts of a path (a/b), a hyphen two words that belong together
// (h-Cat).
var nameSigns = map[rune]string{'/': "경로:/", '-': "붙임:-"}
var nameOrder = []rune{'/', '-'}

// signOrder is the order the graphs of the signs of numbers are made in.
var signOrder = []rune("," + tildeSigns + "-:/.–")

var shownAs = map[string]string{"자릿수": numName}

// headSigns stand at the very front of a number: a sign, or money.
var headSigns = data.Must(loadHeads(data.Must(data.Rows("number-heads", 2))))

var headOrder = data.Must(data.Runes("number-heads", 2))

// unitWords are read as units only right after a number (대, 분, 일,
// 시, 자 and 배 are common inside words). The table is closed. A longer
// unit and a shorter one that ends it are both read where they fit
// (3개월 is 개월; 개 alone leaves 월 to the next word). unitWords are
// the letters a unit is written with, in the order of the table's lines
// (units.tsv), each once; unitRows are the table's lines, letters and
// name (the second column) that follows unitLabel. One letter may stand
// on two lines with two names (m is 미터 and 분): each line makes its
// own graph, so that both candidates are read.
var unitWords = data.Must(data.Column("units", 2))
var unitRows = data.Must(data.Rows("units", 2))

// Places after a number: a greater number (만, 천), 여 (more than), a
// unit, 째 (an order).
const (
	placeBig = iota + 1
	placeOver
	placeUnit
	placeOrder
)

// tildeSigns are the range signs read after a number and its unit: the
// tilde and its variants (tilde operator, fullwidth tilde, wave dash).
const tildeSigns = "~∼～〜"

const (
	unitLabel = "단위:"
	overName  = "어림:여"
	overForm  = "여"
	orderName = "차례:째"
	tildeName = "범위:~(단위 뒤)"
	// loneUnitGraph names the graph of a unit written in letters that are
	// no Hangul syllables (m, km, kg) read as the stem of a word with no
	// number written onto it (700 m): loneUnits.
	loneUnitGraph = "단위(수와 띄어 씀)"
)

// orderHead is the prefix 제 (第) right before a number (제1항, 제36조): it
// tells an order, as 째 after a number does, and its graph orderHeadName
// takes the branch of 째's name (차례). It joins the number where a word
// begins (nothing or a sign before it) or right after the end of a number
// (a digit, or a unit right after a digit: 제36조제1항, 제3조의2제1항).
const (
	orderHead     = '제'
	orderHeadName = "차례:제"
)

// sinoUnits are the units read after a run of digits or a Sino-Korean
// numeral only, not after a native numeral (sino-units.tsv): 인 (人, the
// standard dictionary's dependent noun that counts people after a
// Sino-Korean number: 삼 인, 5인; 한인 and 네인 hold no number).
var sinoUnits = data.Must(data.Column("sino-units", 1))

// unitSuffixes stand after a number and its unit, where 째 stands
// (백 원짜리, 다섯 개가량, 세 명당): suffixes, as 째 and 여 are. The table
// (unit-suffixes.tsv) files each under what stands before it.
// bareSuffixes are those read after a unit with no number before it
// too (bareUnits: ㎡당, kg당): the standard dictionary's suffix -당 stands
// after a number or a unit (마리당 삼천 원, 시간당 얼마).
const (
	setAfterUnit = "수와 단위 뒤"
	setAfterBare = "수 없는 단위 뒤에도"
)

var unitSuffixSets = data.Must(data.Sets("unit-suffixes", setAfterUnit, setAfterBare))
var unitSuffixes = append(slices.Clone(unitSuffixSets[setAfterUnit]), unitSuffixSets[setAfterBare]...)
var bareSuffixes = unitSuffixSets[setAfterBare]

// enclosedRanges are the circled and parenthesized numbers, each one
// letter a number (name 수), as a run of digits is: U+2460-2473 ①-⑳
// (circled digit and number 1-20), U+2474-2487 ⑴-⒇ (parenthesized 1-20),
// U+3251-325F ㉑-㉟ (circled 21-35), U+32B1-32BF ㊱-㊿ (circled 36-50).
// The table is closed; circled Hangul and Latin letters are not numbers.
var enclosedRanges = [][2]rune{{0x2460, 0x2473}, {0x2474, 0x2487}, {0x3251, 0x325F}, {0x32B1, 0x32BF}}
