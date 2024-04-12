package iteration

var repeatCount = 5

// Repeat takes a character and returns five of it
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
