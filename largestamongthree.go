package main

import "fmt"

func main() {

	var n1, n2, n3 int
	fmt.Println("enter the integers :")
	fmt.Scanln(&n1, &n2, &n3)
	if n1 >= n2 && n1 >= n3 {
		fmt.Println(n1, "is the largest")
	} else if n2 >= n1 && n2 >= n3 {
		fmt.Println(n2, "is the largest ")

	} else {
		fmt.Println(n3, "is the largest")
	}

}
