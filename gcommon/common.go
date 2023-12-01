package gcommon

// PanicIfError will trigger a panic if the provided error is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
