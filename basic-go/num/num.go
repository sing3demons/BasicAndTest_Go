package num

import "fmt"

func Number() error {
	var num int
	fmt.Print("input number : ")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		println("")
	}
	return nil
}
