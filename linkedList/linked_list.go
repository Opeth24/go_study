package linkedlist

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	elem := queue.Front()
	if elem == nil {
		return nil
	}
	queue.Remove(elem)
	return elem
}

func printQueue(queue *list.List) {
	for head := queue.Front(); head != nil; head = head.Next() {
		fmt.Printf("%v", head.Value)
	}
}

func LinkedStart() {
	mylist := list.New()
	Push(1, mylist)
	Push(2, mylist)
	Push(3, mylist)
	printQueue(mylist)
	Pop(mylist)
	printQueue(mylist)
}
