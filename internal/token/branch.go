//ff:type feature=token type=model
//ff:what branch — one way of reading a stretch of a unit: the words printed, the graph names, the least verdict and the summed preference
package token

type branch struct {
	marked  string
	names   string
	verdict float64
	pref    float64
}
