//ff:func feature=token type=helper control=sequence
//ff:what registry — the registry of a scan, made when the scan first needs it
package token

func (sc *scan) registry() *registry {
	if sc.reg == nil {
		sc.reg = newRegistry(0)
	}
	return sc.reg
}
