package assert

func Panic(expr bool) {
	if !expr {
		panic("assertion failed")
	}
}
