package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	result := 0
	current := 1

	for current <= number {
		result = result + current
		//current += 1
		current = current + 1
	}

	fmt.Println(result)
}

//	number, _ := strconv.Atoi(os.Args[1])
