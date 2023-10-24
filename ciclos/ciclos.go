package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	i := 0
	for {

		fmt.Println(i)
		i++

		// if i == 2 {
		// 	continue
		// }
		if i > 10 {
			break
		}
	}

}
