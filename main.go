// All files starts with package
package main

import (
	/*
		Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
		The format 'verbs' are derived from C's but are simpler.

		By convention, the package name is the same as the last element of the import path.
	*/

	"github.com/pmoron94/go-lessons/keyboard"
)

func main() { // If cloud environment => main() receives params
	keyboard.Execute()
}

// in console type go run .
// in console type go build . to compile executable file
