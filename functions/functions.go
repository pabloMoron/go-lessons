package functions

import "fmt"

// A function can take zero or more arguments.
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// A function can return any number of results.
func IntToString(value1, value2, value3 int) (bool, string) {
	result := fmt.Sprint(value1)
	statusOk := true

	return statusOk, result
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
