package main

import (
	"fmt"
	sandy "go_tutorials/sandy"
	vedu "go_tutorials/vedu"
	"math/rand"
)

//	func main() {
//		vedant := "vedant"
//		vedu := 23
//		fmt.Println("testing")
//		if vedant == "vedant" && vedu == 23 {
//			fmt.Println("hello vedant")
//		} else if vedu != 23 && vedant != "vedant" {
//			fmt.Println("both are not matching")
//		} else if vedant != "vedant" {
//			fmt.Println("vedant not matching")
//		} else if vedu != 23 {
//			fmt.Println("vedu are not matching")
//		}
//	}
// func main() {
// 	vedu := []int{1, 2, 4, 8, 16, 32, 64, 128}

// 	for test := range vedu {
// 		if test%2 == 0 {
// 			fmt.Println(test)
// 		}
// 	}

// 	fmt.Println("testing")
// }

func main() {
	test_demo(rand.Intn(10000))
	demo := demo_test(3)
	fmt.Println(demo)
	testing := vedu.Vedu()
	fmt.Println(testing)
	fmt.Println(vedu.Jiya())

	vedu.Maya(122)
	sandy.Sandy()
}
