//ff:type feature=rule type=model
//ff:what Not — a criterion that turns a yes/no rule round: with it the rule applies where it would not, and does not where it would (the engine has no negation, so the direction rides on the criteria)
package rule

// The rules that take it say so (Holds): PieceIs, Listed, StemKind,
// MarkIs and MarkBoth.
type Not struct{}
