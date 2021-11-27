package main

import "fmt"

type studentNode struct {
	no   int
	name string
	next *studentNode
	pre  *studentNode
}

//创建
func InsertStudent(head *studentNode, student *studentNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = student
	student.pre = temp

	// for {
	// 	fmt.Printf("[%d,%s]==>", temp.no, temp.name)
	// 	temp = temp.pre
	// 	if temp.next == nil {
	// 		break
	// 	}

	// }

}

//添加  (这里插入的元素不可以是最后一个元素)
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
		student.pre = temp
		if temp.next != nil {
			temp.next.pre = student
			temp.next = student
		}

	}
}

//删除有序双链表的其中一个元素
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
	if temp.next != nil {
		student.next.pre = student.pre

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

func unlistStudent(head *studentNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for {
		if temp.pre == nil {
			break
		}
		fmt.Printf("[%d,%s]==>", temp.no, temp.name)

		temp = temp.pre
	}

}

// func deleteStudent(head *studentNode, student *studentNode) {
// 	temp := head
// 	for {
// 		if temp.next == nil {
// 			break
// 		} else if temp.next.no == student.no {
// 			break
// 		}
// 		temp = temp.next
// 	}
// 	temp.next = student.next
// }

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
	InsertStudent(head, studentNode1)
	// InsertStudent(head, studentNode2)
	// InsertStudent(head, studentNode3)
	InsertStudent(head, studentNode4)
	InsertStudent(head, studentNode5)

	// sortInsert(head, studentNode1)
	sortInsert(head, studentNode2)
	sortInsert(head, studentNode3)
	// sortInsert(head, studentNode4)
	// sortInsert(head, studentNode5)

	listStudent(head)
	fmt.Println()

	unlistStudent(head)
	fmt.Println()

	// deleteStudent(head, studentNode3)
	// listStudent(head)
	// fmt.Println()

	deleteStudent(head, studentNode5)
	listStudent(head)
	fmt.Println()

	unlistStudent(head)

}
