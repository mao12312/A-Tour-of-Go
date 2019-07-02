package Switch

import (
	"fmt"
	"runtime"
)

func Main(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Printf("%s.", os)

	}
}