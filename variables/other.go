package variables

import (
	"fmt"
	"time"
)

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
