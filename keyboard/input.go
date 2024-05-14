package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var phrase string

var err error

func Execute() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("1 - !!!!!!" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("2 - !!!!!!" + err.Error())
		}
	}

	fmt.Println("Ingrese frase: ")
	if scanner.Scan() {
		phrase = scanner.Text()
		fmt.Println(phrase, num1*num2)
	}
}
