package main

import (
	"bou.ke/monkey"
	"fmt"
)

func Yes() string {
	return "yes"
}

func No() string {
	return "no"
}

func main() {
	fmt.Println(Yes())
	monkey.Patch(Yes,No)
	fmt.Println(No())
}