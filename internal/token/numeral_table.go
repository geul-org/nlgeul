package token

// What counts as a number right before a unit, besides a run of digits
// (decision 014-32). The tables are closed; a form is added by writing
// it here.
//
// Native numerals in the form they take before a unit: one of
// nativeOnes, alone or after 열 (열두); 열 and 스무 alone; 몇 and 여러.
var nativeOnes = []string{"한", "두", "세", "네", "다섯", "여섯", "일곱", "여덟", "아홉"}

var nativeAlone = []string{"열", "스무", "몇", "여러"}

// Sino-Korean numerals are read letter by letter: digits and the
// greater numbers, in any run that has no two digits side by side
// (삼십, 이백, 만). 수 stands only before a greater number (수십).
const (
	sinoDigits = "일이삼사오육칠팔구"
	sinoBigs   = "십백천만억조"
	sinoSome   = '수'
)

// numeralCount is the number of entries of the two tables: a numeral
// is taken as one of them, as the tables give no counts.
const numeralCount = 13 + 16

// numberLetters are Latin letters written where a number stands (N개,
// k번째). Only a letter that is a run by itself counts.
var numberLetters = []string{"N", "n", "X", "x", "K", "k"}

// bigName is the name of a greater number after a run of digits (7만,
// 3천): a piece of its own before the unit.
const bigName = "큰 수:"
