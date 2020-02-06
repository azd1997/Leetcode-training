package lt706

import "math"

// 设计哈希映射


// 由于这道题所有值都在[1,1000000]范围
// 作为一道定位easy的题目，大概率不是让你真去设计工业级的哈希表
//
// 这里直接用数组存储，映射关系为 idx+1 -> v
// 而且看示例，本身也是k->v一一对应，大小相等

type MyHashMap struct {
	data [1000001]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	data := [1000001]int{}
	for i:=1; i<=1000000; i++ {data[i] = math.MinInt32}
	return MyHashMap{data: data}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.data[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
// 醉了，题目明明说所有值范围在[1,1000000]范围，结果测例中出现了0值。
// 这使得这里不能直接以数组默认的0值来判断键值对不存在
func (this *MyHashMap) Get(key int) int {
	if this.data[key] == math.MinInt32 {return -1}
	return this.data[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	this.data[key] = math.MinInt32
}