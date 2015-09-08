package view

import (
	"fmt"
)

func ViewFail(value string) {
	fmt.Print("Content-Type: text/plain\n\n")
	fmt.Println("Error " + value)
}
