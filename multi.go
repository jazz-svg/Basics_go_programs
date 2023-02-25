package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the integer : ")
	fmt.Scan(&n)
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, " X ", i, " = ", n*i)
		i++
	}
}
