package variables

import "fmt"

func PrintIntegers() {
	var integer int              // Declaration, variables always are initialized with default value, not null
	integer_32_bits := int32(10) // Declaration + assignation operator (:=)
	integer_64_bits := int64(9223372036854775807)

	// Cannot assign different data_types
	// integer_32_bits = integer_64_bits

	fmt.Println(`integer: `, integer)
	fmt.Println(`integer_32_bits: `, integer_32_bits)
	fmt.Println(`integer_64_bits: `, integer_64_bits)

	fmt.Println(`64 to 32 overflow: `, int32(integer_64_bits))
}
