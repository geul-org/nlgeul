package token

// adverbialBound are the dependent nouns of the closed set that school
// grammar counts as adverbial (부사성 의존 명사: 대로, 만큼, 듯, 채): they
// make an adverbial with the modifier before them and take no particle of
// the subject or the complement. The dictionary gives 듯이 as a dependent
// noun of its own (듯 is its short form), so 듯이 after a modifying form
// is that noun and not 듯 with the nominative 이 (boundAfterRule).
var adverbialBound = []string{"대로", "만큼", "듯", "채"}
