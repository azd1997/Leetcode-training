package lt706

// 拉链法，数组加链表


type MyHashMap3 struct {
	data [500]*Node		// 500个分组
}

// 单链表节点  每一个分组都是一条单链表
type Node struct {
	k, v int
	next *Node
}


/** Initialize your data structure here. */
func Constructor3() MyHashMap3 {
	// 将所有数分到500个分组去
	return MyHashMap3{data: [500]*Node{}}
}



func (this *MyHashMap3) hash(key int) int {
	Mod := 500
	return key % Mod
}

/** value will always be non-negative. */
func (this *MyHashMap3) Put(key int, value int)  {
	// 尝试获取key对应的分组（单链表头结点）
	cid := this.hash(key)
	node := this.data[cid]

	// 情况1， 该分组还没有节点
	if node == nil {
		this.data[cid] = &Node{key, value, nil}
		return
	}

	// 情况2，该分组下有节点，遍历该链表找是否有key相等的节点
	for node.next != nil {		// 遍历到倒数第二个节点
		if node.k == key {		// 找到
			node.v = value
			return
		}
		node = node.next
	}
	// node为最后一个节点
	if node.k == key {
		node.v = value
	} else {node.next = &Node{key, value, nil}}
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap3) Get(key int) int {
	cid := this.hash(key)
	node := this.data[cid]
	if node == nil {return -1}
	for node != nil {
		if node.k == key {
			return node.v
		}
		node = node.next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap3) Remove(key int)  {
	cid := this.hash(key)
	node := this.data[cid]
	if node == nil {return}

	// 为了方便删除链表头部节点，设置哨兵节点（伪头）
	// 为了能够删除尾部节点，设置pre移动哨兵
	head := &Node{next:node}
	pre, cur := head, head.next
	for cur != nil {
		if cur.k == key {
			pre.next = cur.next	// 删除cur.next
			return
		}
		pre, cur = cur, cur.next
	}

	this.data[cid] = head.next	// 记得这一步
}

