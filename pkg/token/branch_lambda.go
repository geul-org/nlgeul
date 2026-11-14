//ff:func feature=token type=helper control=sequence
//ff:what Lambda — the place of a candidate in the sorting formula (decision 014-35): its preference plus the package's Lambda times log10 of its verdict; minus infinity when the verdict is 0 or less. The candidates of a node come in this order, and it is the lambda of a candidate in nlgeul token --format json
package token

func (b Branch) Lambda() float64 {
	return lambdaOf(b)
}
