//ff:func feature=data type=helper control=sequence
//ff:what Must — a table as it was read, or a panic with the error: a table held in the program that is written wrong stops the program as it starts
package data

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return v
}
