package variables

import (
	"fmt"
	"time"
)

// Variables declared without an explicit initial value are given their zero value.
// The zero value is:
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
func OtherVariables() {
	var name string
	var isEnabled bool
	var salary float32
	var birthdate time.Time = time.Now()

	name = "Pablo"
	isEnabled = true
	salary = 10.25

	fmt.Println("[string] Name: ", name)
	fmt.Println("[bool] Enabled: ", isEnabled)
	fmt.Println("[float32] Salary:", salary)
	fmt.Println("[time.Time] birthdate: ", birthdate)
}

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.
func Constants() {
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
