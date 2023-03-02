package main

import "fmt"

func main() {

	var n int
	fmt.Print("enter the integer :")
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println(n, "number is even")
	} else {
		fmt.Println(n, "number is odd")
	}

}
