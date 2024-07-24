package main

import (
	"fmt"
)

func test_demo(num int) {

	fmt.Println(num)
	if num%2 != 0 {
		fmt.Println("number is not divisible by 2")
	} else {
		fmt.Println("number is divisible by 2")
	}
}

func demo_test(num int) []int {
	final_list := []int{} // Initialize an empty slice

	// Loop to potentially add elements
	for i := 0; i < num; i++ {
		final_list = append(final_list, i)
	}

	// Return the final list after the loop completes
	return final_list
}
