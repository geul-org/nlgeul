//ff:type feature=token type=model
//ff:what Name — one reading of a part: its kind (주격, 꾸밈, 열쇠, 맺음, 수, 경로 …) and its form (가, 은, 체언 …; "" when the kind says it all)
package token

type Name struct {
	Kind string
	Form string
}
