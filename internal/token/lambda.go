package token

// Lambda is how much a candidate's verdict weighs against its preference
// in the sorting formula: a candidate counts as its preference plus
// Lambda times log10 of its verdict (decision 014-35). This is the one
// place the number is written; internal/combo takes it from here.
const Lambda = 15
