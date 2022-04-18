package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	var temp *Node
	var head *Node
	var count int

	fmt.Println("Enter number of nodes")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		temp = new(Node)
		fmt.Println("Enter node data")
		fmt.Scan(&temp.data)
		temp.next = head
		head = temp
	}

	fmt.Println("*************************************")

	for head != nil {
		fmt.Println(head.data, head.next)
		head = head.next
	}

}
