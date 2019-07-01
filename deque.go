package main

import "fmt"

func main() {

	deq := deque{}

	for i := 0; i < 5; i++ {
		deq.push_front(i)

	}
	fmt.Println(deq)
	for i := 0; i < 5; i++ {
		deq.push_back(i)

	}

	fmt.Println(deq.pop_back())
	fmt.Println(deq.pop_front())
	fmt.Println(deq.peekdeque('f'))
	fmt.Println(deq.peekdeque('b'))
}

type deque struct {

}