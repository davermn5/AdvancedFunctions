package main

//Import the fmt package
import "fmt"

// a is a pointer to the same address (which resolves to the actual value 20) that the variable age already points to.
// t is a pointer to the same address (which resolves to the actual value "John") that the variable text points to.
func update(a *int, t *string) {
	*a = *a + 5      //Add 5 to the value which is being stored at the address that pointer 'a' and variable 'age' points to
	*t = *t + " Doe" //Append " Doe" to the value which is being stored at the address that pointer 't' and variable 'text' points to
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	// Pass in age and text as param to func args
	update(&age, &text)

	fmt.Println("After:", text, age)
}
