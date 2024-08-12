package main

import "fmt"

type Day string

const (
	S         Day = "Sunday"
	M         Day = "Monday"
	Tuesday   Day = "Tuesday"
	Wednesday Day = "Wednesday"
	Thursday  Day = "Thursday"
	Friday    Day = "Friday"
)

func main() {
	fmt.Println(S, M == Thursday)
}
