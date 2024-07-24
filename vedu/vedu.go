package vedu

import (
	"fmt"
)

func Vedu() map[string]int {
	testing := make(map[string]int)
	testing = map[string]int{
		"name": 23,
		"age":  23,
	}
	return testing
}

type Student struct {
	Name  string
	Age   int
	Score []int
}

func Jiya() Student {
	student1 := Student{
		Name: "Student",
		Age:  12,
		Score: []int{
			1,
			3,
			4,
		},
	}
	return student1
}

var ch chan int

func recivedata() {
	i := <-ch
	fmt.Println("1111111111111111111111111111")
	fmt.Println(i, "wewewewewewewe")
	ch <- 42

}

func senddata(num int) {
	ch <- num
	i := <-ch
	fmt.Println(i)

}

func Maya(num int) {
	fmt.Println("dsdsdsdsdsddssds")

	ch = make(chan int)

	go senddata(num)
	recivedata()
}
