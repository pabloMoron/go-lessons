package exercises

import (
	"fmt"
	"strconv"
)

func StringToInt(number string) (parsedInt int, phrase string) {
	parsedInt, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Err", err.Error())
		parsedInt = 0
		phrase = "Error"
		return
	}

	if parsedInt >= 100 {
		phrase = "Is greater or equals to 100"
	} else {
		phrase = "Is lower than 100"
	}
	return
}
