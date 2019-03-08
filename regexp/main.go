package main

import (
	"fmt"
	"regexp"
)

func main() {
	line := `<script>var goVersion = "go1.12";</script>`
	
	re := regexp.MustCompile(`(go1\.)(..)`) // two "capture groups".
	xs := re.FindStringSubmatch(line)       // xs[0] is the whole string, []xs1 is capgroup1, xs[2] is capgroup2
	
	ver := xs[2]
	fmt.Println(ver)
}