package main

import (
	"fmt"
)

type Boy struct {
	No   int
	next *Boy
}

//形成一个环形链表(自动生成各节点)
func AddBoy(num int) *Boy {

	head := &Boy{}
	curboy := &Boy{}
	for i := 1; i < num+1; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			head = boy
			curboy = boy
			curboy.next = head
		} else {
			curboy.next = boy
			curboy = curboy.next
			curboy.next = head

		}

	}
	return head

}

func list(head *Boy) {

	curboy := head
	for {
		fmt.Printf("小朋友的编号是:%d=>", curboy.No)
		if curboy.next == head {
			break

		}
		curboy = curboy.next
	}
}

func main() {
	head := AddBoy(5)
	list(head)
}
