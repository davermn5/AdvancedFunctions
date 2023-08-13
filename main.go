package main

import "fmt"

// a is a pointer to the same address (which resolves to the actual value 20) that the variable age already points to.
// t is a pointer to the same address (which resolves to the actual value "John") that the variable text points to.
func update(a *int, t *string) {
	*a = *a + 5
	*t = *t + " Doe"
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After:", text, age)
}
