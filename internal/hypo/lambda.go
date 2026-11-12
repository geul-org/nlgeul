package hypo

// Lambda is how much a reading's verdict weighs against its preference
// in the sorting formula of order: a reading counts as its preference
// plus Lambda times log10 of its verdict. At 0 the verdict counts for
// nothing; the greater it is, the later a hypothesis with a weak
// reading comes. Change this one number to move it.
const Lambda = 15
