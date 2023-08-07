package main

import "fmt"

func main() {
	fmt.Println("Hello Guys", CoverFunc()+UncoverFunc())
}

func UncoverFunc() int {
	return 20
}

func CoverFunc() int {
	return 15
}
