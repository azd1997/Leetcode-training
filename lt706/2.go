package lt706



// 完全用数组的话，太bug了，所以，现在参考提交区的一个做法
// 将之换为二维数组。实现如下




type MyHashMap2 struct {
	data [][]*Elem
}

type Elem struct {
	k, v int
}


/** Initialize your data structure here. */
func Constructor2() MyHashMap2 {
	// 将所有数分到50个数组去
	return MyHashMap2{data: make([][]*Elem, 50)}
}


func (this *MyHashMap2) getElem(key int) *Elem {

	// 获取当前key对应的分区编号
	cid := key % 50

	// 查看当前分区是否包含数据
	if this.data[cid]==nil {return nil}

	// 遍历当前分区
	for idx, elem := range this.data[cid] {
		if elem.k == key {return this.data[cid][idx]}
	}
	return nil
}

/** value will always be non-negative. */
func (this *MyHashMap2) Put(key int, value int)  {
	// 尝试获取key对应的elem
	elem := this.getElem(key)
	if elem == nil {
		cid := key % 50
		this.data[cid] = append(this.data[cid], &Elem{
			k: key,
			v: value,
		})
		return
	}
	elem.v = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
// 醉了，题目明明说所有值范围在[1,1000000]范围，结果测例中出现了0值。
// 这使得这里不能直接以数组默认的0值来判断键值对不存在
func (this *MyHashMap2) Get(key int) int {
	elem := this.getElem(key)
	if elem==nil {return -1}
	return elem.v
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap2) Remove(key int)  {
	cid := key % 50
	if this.data[cid] == nil {return}
	for idx, elem := range this.data[cid] {
		if elem.k == key {
			this.data[cid] = append(this.data[cid][:idx], this.data[cid][idx+1:]...)
			return
		}
	}
}
