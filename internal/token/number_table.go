package token

import "github.com/geul-org/koparse/pkg/rule"

// Numbers (decision 014-31). A run of digits is a number; a sign
// between two runs of digits is judged by a graph for each of its
// readings, and a reading that stands joins the runs into one number.
// numName is what a number is called in the output; numGraph names the
// graph of a run of digits.
const (
	numName  = "수"
	numGraph = "수:숫자열"
)

// signShapes give, sign by sign, the readings that join runs of digits
// and the shape each asks of them, in the order they are named.
// 자릿수 is shown as a number (shownAs).
var signShapes = map[rune][]rule.Shape{
	',': {{Of: "자릿수:,", Min: 2, Lens: [][2]int{{1, 3}, {3, 3}}}},
	'~': {{Of: "범위:~", Min: 2, Max: 2}},
	'-': {
		{Of: "범위:-", Min: 2, Max: 2},
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

var signOrder = []rune{',', '~', '-', ':', '/', '.'}

var shownAs = map[string]string{"자릿수": numName}

// headSigns stand at the very front of a number: a sign, or money.
var headSigns = map[rune]string{'-': "부호:-", '+': "부호:+", '$': "돈:$", '₩': "돈:₩"}

var headOrder = []rune{'-', '+', '$', '₩'}

// unitWords are read as units only right after a number (대, 분, 일,
// 시, 자 and 배 are common inside words). The table is closed. A longer
// unit and a shorter one that ends it are both read where they fit
// (3개월 is 개월; 개 alone leaves 월 to the next word).
var unitWords = []string{
	"개", "명", "번", "건", "줄", "쪽", "장", "권", "대", "마리", "가지", "곳", "군데", "차", "회", "층", "배", "등", "위", "점",
	"년", "월", "일", "시", "분", "초", "주", "개월", "시간", "살", "세",
	"원", "달러", "엔",
	"자", "글자", "단어", "문장", "토큰",
	"%", "℃", "kg", "cm", "km", "ms", "MB", "GB",
}

// latinUnits are the units written in Latin letters: a run of digits
// and one of them is a number and its unit.
var latinUnits = []string{"kg", "cm", "km", "ms", "MB", "GB"}

// Places after a number: a greater number (만, 천), 여 (more than), a
// unit, 째 (an order).
const (
	placeBig = iota + 1
	placeOver
	placeUnit
	placeOrder
)

const (
	overName  = "어림:여"
	orderName = "차례:째"
	tildeName = "범위:~(단위 뒤)"
)
