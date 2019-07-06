package p22

import "fmt"

func Main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The Value:", m["Answer"])

	m["Anwser"] = 63
	fmt.Println("The Value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Persent?", ok)
}