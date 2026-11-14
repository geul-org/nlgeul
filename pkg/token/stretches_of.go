//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stretchesOf — contests read the old way (the signs at a unit's edges, a unit with a sign of a pair written onto) as the stretches of a unit a stream takes: each with its branches and no cells
package token

func stretchesOf(contests []Contest) []unitStretch {
	out := make([]unitStretch, 0, len(contests))
	for _, c := range contests {
		out = append(out, unitStretch{contest: c})
	}
	return out
}
