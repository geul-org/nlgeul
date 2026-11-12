package hypo

// Lambda is how much a candidate's verdict weighs against its preference
// in the sorting formula of order: a candidate counts as its preference
// plus Lambda times log10 of its verdict. At 0 the verdict counts for
// nothing; the greater it is, the later a combination with a weak
// candidate comes. Change this one number to move it.
const Lambda = 15
