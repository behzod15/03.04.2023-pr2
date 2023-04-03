package main

import "fmt"

func main () {
	var num1, num2, cnt int
	cnt = 0
	fmt.Println("Enter first number of the range:")
	fmt.Scan(&num1)
	fmt.Println("Enter second number of the range:")
	fmt.Scan(&num2)
	for i := num1; i <= num2; i++{
		for j := 0; j <= cnt; j++{
			fmt.Println(i)
		}
		cnt ++ 
	}
}