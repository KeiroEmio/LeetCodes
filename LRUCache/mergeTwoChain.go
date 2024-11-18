package main

//合并两个有序链表：
//
//题目描述：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//示例：
//go
//复制代码
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

type Node1 struct {
	value int
	prev  *Node1
	next  *Node1
}

type LinkedList1 struct {
	head *Node1
	tail *Node1
}

func ConstructorChain(nums []int) *LinkedList1 {
	ll := &LinkedList1{} // 创建一个空链表

	// 遍历数组并创建节点
	for _, num := range nums {
		node := &Node1{value: num} // 假设 key 和 value 都使用数组的值
		if ll.tail == nil {
			// 如果链表为空，设置头和尾指向新节点
			ll.head = node
			ll.tail = node
		} else {
			// 将新节点添加到链表末尾
			ll.tail.next = node
			node.prev = ll.tail
			ll.tail = node
		}
	}

	return ll
}

func ConstructorChain1(nums []int) *LinkedList1 {
	ll := &LinkedList1{}

	for _, num := range nums {
		node := &Node1{value: num}

		if ll == nil {
			ll.head = node
			ll.tail = node
		} else {
			//三，两节点绑定，一连结尾定义
			ll.tail.next = node
			node.prev = ll.tail
			ll.tail = node
		}
	}

	return ll
}

func main() {
	num1 := []int{1, 2, 3, 4, 5}
	num2 := []int{1, 2, 3}

}
