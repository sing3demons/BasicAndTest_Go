package main

import "fmt"

func Number() {
	var num int
	fmt.Print("input number : ")
	fmt.Scanf("%d", &num)

	if num >= 1 {

		for i := 1; i <= num; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			println("")
		}
		return
	}
	fmt.Print("input number only")

}

func main() {
	Number()
}
