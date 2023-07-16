package main

import (
	"dsa/sll"
	"fmt"
)

func main() {
	l := &sll.SingleLinkedList[int]{}
	fmt.Println(l.GetAll())
}
