package blocks

import (
	"fmt"
	"runtime"
)

func Conditional() {
	if runtime.GOOS == "darwin" {
		fmt.Println("This is a macOS")
	} else {
		fmt.Printf("This is a %s, not a macOS \n", runtime.GOOS)
	}

	// We can declare, assign a variable and declare the condition inline
	if arch := runtime.GOARCH; arch == "arm64" {
		fmt.Println("This is an arm64 arch")
	} else {
		fmt.Printf("This is %s arch, not an arm64 arch \n", arch)
	}
}

func Switch() {
	// We can declare, assign a variable and evaluate inline
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("macOS")
	default:
		fmt.Println("Windows")
	}
}
