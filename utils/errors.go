package utils

// Panics if err is not null.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
