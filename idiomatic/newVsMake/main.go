package main

func main() {
	// newBlabla is for when you return a pointer
	// makeBlabla is for when you return a type
}

type mystruct struct {
	one string
	two string
}

func newStruct(one, two string) *mystruct {
	return &mystruct{
		one: one,
		two: two,
	}
}

func makeString(one, two string) string {
	return one + two
}
