package main

import (
	queue "Proj/Queue"
	"fmt"
)

func main() {

	q, _ := queue.NewCircularQueue(10)

	q.PrintQueue() //[0 0 0 0 0 0 0 0 0 0]

	for i := 1; i <= 10; i++ {
		if err := q.Push(i); err != nil {
			fmt.Println(err)
		}
	}

	q.PrintQueue() //[1 2 3 4 5 6 7 8 9 10]

	if err := q.Pop(); err != nil {
		fmt.Println(err)
	}

	q.PrintQueue() //[0 2 3 4 5 6 7 8 9 10]

	if err := q.Push(11); err != nil {
		fmt.Println(err)
	}

	q.PrintQueue() //[11 2 3 4 5 6 7 8 9 10]
}
