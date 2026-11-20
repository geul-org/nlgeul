package parse

// formDeterminers are the demonstrative determiners whose letters are
// also the modifying form of an adjective (이렇다, 그렇다, 저렇다, 어떻다,
// 다르다), a closed set the rule of a sentence prefers as determiners
// where they stand alone (determinerFormRule). differStem is 다른, the
// one of them whose adjective takes the comparison in 와/과, and
// comparedNames are the names of that particle.
var (
	formDeterminers = []string{"이런", "그런", "저런", "어떤", "다른"}
	differStem      = "다른"
	comparedNames   = []string{"부사격:와", "부사격:과"}
)
