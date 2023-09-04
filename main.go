package main

import (
	"fmt"
	"net/http"
)

func main() {
	// password := "aakbjkgbejanfn213j5n13n45l1n5"

	// fmt.Println("Password is: ", password)

	// a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// b := a[1:4]
	// fmt.Println("a:", a)
	// fmt.Println("b:", b)

	// // Works fine even though c is indexing past the end of b.
	// c := b[4:7]
	// fmt.Println("c:", c)

	// // This fails with panic: runtime error: index out of range [4] with length 3
	// d := b[4]
	// fmt.Println("d:", d)

	// if true {
	// 	fmt.Println("WOHO")
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %d!", UncoverFunc()+CoverFunc()+AnotherCoverFunc())
	})

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}

func UncoverFunc() int {
	return 20
}

func CoverFunc() int {
	return 15
}

func AnotherCoverFunc() int {
	return 30
}
