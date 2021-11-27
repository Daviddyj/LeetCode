package main

import "fmt"

type studentNode struct {
	no   int
	name string
	next *studentNode
	// pre  *studentNode
}

// func unlistStudent(head *studentNode, student *studentNode) {
// 	temp := head
// 	for {
// 		if temp.next == nil {
// 			break
// 		}
// 		temp = temp.next
// 	}
// 		student.pre = temp.pre
// 		temp.next = student
// 	for {
// 		fmt.Printf("[%d,%s]==>", temp.no, temp.name)
// 		temp = temp.pre
// 		if temp.next == nil {
// 			break
// 		}

// 	}

// }

func sortInsert(head *studentNode, student *studentNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > student.no {
			break
		} else if temp.next.no == student.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在no:", student.no)

	} else {
		student.next = temp.next
		temp.next = student
	}
}

func listStudent(head *studentNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("这是一个空链表")
		return
	}

	for {
		fmt.Printf("[%d,%s]==>", temp.next.no, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}

	}
}

func deleteStudent(head *studentNode, student *studentNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == student.no {
			break
		}
		temp = temp.next
	}
	temp.next = student.next
}

func main() {
	head := &studentNode{}
	studentNode1 := &studentNode{
		no:   1,
		name: "daiyijie",
	}
	studentNode2 := &studentNode{
		no:   2,
		name: "wangrui",
	}
	studentNode3 := &studentNode{
		no:   3,
		name: "dashuaige",
	}
	studentNode4 := &studentNode{
		no:   4,
		name: "wangqiang",
	}
	studentNode5 := &studentNode{
		no:   5,
		name: "shangzhengxi",
	}

	sortInsert(head, studentNode3)
	sortInsert(head, studentNode1)
	sortInsert(head, studentNode2)
	sortInsert(head, studentNode4)
	sortInsert(head, studentNode5)

	listStudent(head)
	fmt.Println()

	deleteStudent(head, studentNode3)
	listStudent(head)
	fmt.Println()

}
