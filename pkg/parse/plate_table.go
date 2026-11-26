package parse

import "github.com/geul-org/nlgeul/pkg/data"

// The shape of a motor vehicle's registration number (자동차 등록 번호,
// as the notice on registration plates gives it): the number of the kind
// of vehicle (plateHeadLens digits: two, and three on the plates of
// 2019 on), the letter of its use (plateLetters: one Hangul syllable of
// a closed list, 가 to 주 for private cars, 바 사 아 자 배 for hire, 하
// 허 호 for rent) and the serial number (plateSerialLen digits), written
// with no space between (206너9499, 74모2238). An older plate begins
// with the name of the region that issued it (plateRegions: 서울, 경기)
// and then has a kind of plateRegionLens digits (서울34가5678).
var plateLetters = data.Must(data.Column("plate-letters", 1))
var plateRegions = data.Must(data.Column("plate-regions", 1))
var plateHeadLens = []int{2, 3}
var plateRegionLens = []int{1, 2}

const plateSerialLen = 4

// plateKind names a registration number. While a unit is read the whole
// number is one stem named so (an atom), as a court case number is; it
// is sent up as the pieces of one word (plateWord): the kind and the
// serial number, each a number (numName), and between them the letter,
// named plateKind with the letter as its form
// (206[수]+너[차등록:너]+9499[수]); the region before them is a nominal
// (regionName: 서울[열쇠:체언]+34[수]+가[차등록:가]+5678[수]).
const plateKind = "차등록"

var regionName = Name{keyKind, "체언"}

// plateUnit is the unit a registration number is numbered with when one
// is written after it (74모2238호), as a case number is.
const plateUnit = "호"
