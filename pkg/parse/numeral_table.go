package parse

import "github.com/geul-org/nlgeul/pkg/data"

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
// sinoMyriads are the greater numbers a numeral is grouped by, four
// digits a group (clause 44 of the orthography: 십이억 삼천만).
const (
	sinoDigits  = "일이삼사오육칠팔구"
	sinoMyriads = "만억조"
	sinoBigs    = "십백천" + sinoMyriads
	sinoSome    = '수'
	// sinoSix is written 륙 (sinoSixAfter) right after sinoFive in an
	// approximate number (오륙: approxPair).
	sinoFive     = "오"
	sinoSix      = "육"
	sinoSixAfter = "륙"
)

// numeralCount is the number of entries of the two tables: a numeral
// is taken as one of them, as the tables give no counts. 13 is the
// native numerals (nativeOnes and nativeAlone), 16 the Sino-Korean ones
// (sinoDigits, sinoBigs and sinoSome); the sum is written by hand and
// follows the tables only as long as it is kept so.
const numeralCount = 13 + 16

// numberLetters are Latin letters written where a number stands (N개,
// k번째). Only a letter that is a run by itself counts.
var numberLetters = []string{"N", "n", "X", "x", "K", "k"}

// bigName is the name of a greater number after a run of digits (7만,
// 3천): a piece of its own before the unit.
const bigName = "큰 수:"

// numeralReachMax is how far back from a place a numeral of the tables
// is looked for (numeralReach).
var numeralReachMax = numeralReach()

// orderLetters are the letters that number the items (목) of a law's
// article, in their order (가목, 나목, …, 하목, then 거목, …, 허목). They
// stand where a number stands before letterUnit only (orderLetter), each
// taken as one of them.
const (
	orderLetters = "가나다라마바사아자차카타파하거너더러머버서어저처커터퍼허"
	letterUnit   = "목"
)

// hanNumerals are the Han characters a number is written with
// (han-numerals.tsv): the digits 〇 and 零 to 九, the greater numbers 十,
// 百, 千, 萬, 億 and 兆, and 万 as 萬 is also written. A run of them is one
// number however they are put together (十二, 二三, 二〇二四), named
// 수:한자십진 by its letters (numberNameOf).
var hanNumerals = data.Must(data.Column("han-numerals", 1))

// englishNumerals are the numerals of English a number is written with
// (english-numerals.tsv), in small letters: zero to nineteen, the tens
// twenty to ninety, and hundred, thousand, million, billion and
// trillion. A number written with them is named 수:영문십진
// (englishNumber).
var englishNumerals = data.Must(data.Column("english-numerals", 1))
