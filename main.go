// All files starts with package
package main

import (
	/*
		Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
		The format 'verbs' are derived from C's but are simpler.

		By convention, the package name is the same as the last element of the import path.
	*/
	"fmt"
	"math/rand"

	"github.com/pmoron94/go-lessons/functions"
	"github.com/pmoron94/go-lessons/hello_world"
	"github.com/pmoron94/go-lessons/variables"
)

func main() { // If cloud environment => main() receives params
	fmt.Println("Random:", rand.Intn(10)) // Add newline
	fmt.Println("------------")

	hello_world.Greet()
	fmt.Println("------------")

	variables.PrintIntegers()
	fmt.Println("------------")

	variables.OtherVariables()
	fmt.Println("------------")

	integer := 2318
	txt1 := string(integer) // bad conversion
	fmt.Println(txt1)
	fmt.Println("------------")

	integer2 := 8
	ok, txt2 := functions.IntToString(integer2, 0, 0)
	if ok {
		fmt.Println(txt2)
	}
	fmt.Println("------------")

	x, y := functions.Split(11)
	fmt.Println("x: ", x, "y: ", y)
	fmt.Println("------------")
}

// in console type go run .
// in console type go build . to compile executable file
