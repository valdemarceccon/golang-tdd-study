package iteration

const repeatCount = 5

func Repeat(value string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += value
	}
	return repeated
}
