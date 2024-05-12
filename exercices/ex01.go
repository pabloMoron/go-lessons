package exercices

import (
	"fmt"
	"strconv"
)

func StringToInt(number string) (parsedInt int, phrase string) {
	parsedInt, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Err", err)
	}
	if parsedInt >= 100 {
		phrase = "Is greater or equals to 100"
	} else {
		phrase = "Is lower than 100"
	}
	return
}
