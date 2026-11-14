//ff:type feature=token type=model
//ff:what unitReader — one unit read as the stream asks for its stretches: the stretches made and not yet sent (the signs before its letters at first), the net of its letters (netRun) while it has stretches to give, what makes the stretches of the signs after the letters when the net is done, and the trace lines made since they were last taken
package token

type unitReader struct {
	queue []unitStretch
	run   *netRun
	after func() []unitStretch
	notes []string
}
