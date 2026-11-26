package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

// spokenMarkers are the closings and particles written as they are
// spoken or as an older or regional way of speaking has them, each with
// the standard marker it stands for: 구 for the closing 고 (서구, 마시구),
// 두 for the particle 도 (마당두, 작게두), 으루 for 으로, 우 and 수 for the
// 하오체 closings 오 and 소 (가우, 좋수), 든 for the modifying 던. The
// piece is printed with the letters as they stand and the name of the
// standard marker (구[맺음:고]): the entry is a copy of the standard
// marker's with the written form (spokenEntry). The rows are written
// down from the grammar of the standard markers, not copied from a list
// or a dictionary, and the lines of spoken-markers.tsv in the data
// package are the whole table. A row is read only with a word table,
// after a stem of the table (addSpoken): the letters are common at the
// end of words.
var spokenMarkers = data.Must(data.Each("spoken-markers", 3, spokenOf))

// spokenFlags are the conditions a row of spokenMarkers may name: a
// form read only after the copula stem (요 for 오: 마당이요).
var spokenFlags = map[string]int{"서술격만": fNeedIda}
