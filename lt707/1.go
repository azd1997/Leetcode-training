package lt707

// 设计链表

// 单链表版本

type MyLinkedList struct {
	head *Node
	length int
}


type Node struct {
	val int
	next *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &Node{},
		length: 0,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index<0 || index>=this.length {return -1}
	node := this.head
	for i:=0; i<=index; i++ {
		node = node.next
	}
	return node.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	node := &Node{
		val: val,
	}
	if this.length > 0 {
		node.next = this.head.next
	}
	this.head.next = node
	this.length++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	node := this.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{
		val: val,
	}
	this.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	//fmt.Println("length=", this.length)
	if index > this.length {return}
	if index == this.length {
		this.AddAtTail(val)
		return
	}
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	// 一般情况
	node := this.head
	for i:=0; i<index; i++ {
		node = node.next
	}   // node到达index node前一个
	//fmt.Println("nodeval=", node.val)
	newNode := &Node{
		val: val,
	}
	newNode.next = node.next
	node.next = newNode
	//fmt.Println("nodeval=", newNode.next.val)
	this.length++
	//fmt.Println("length=", this.length)
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index<0 || index>=this.length {return}
	node := this.head
	for i:=0; i<index; i++ {
		node = node.next
	}
	node.next = node.next.next
	this.length--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
