package lt380

import "math/rand"

// 常数时间内插入、删除、获取随机元素

// 常数时间插入删除，显然是哈希表咯 (链表号称也是O(1)，但意义不一样)
// 获取随机元素： 假设直接使用map来做，那么就只能把所有键再存到一个数组中，
// 然后数组下标随机选取。
// 这就是第一版基于map的实现
// 第二版将自己写哈希表

// 在第一版中，由于使用了数组来实现随机获取元素，
// 那么删除元素时要想实现O(1)，必须使哈希表指向其在数组中的位置
// 而且在数组中删除时要使用交换删除，而不是删除当前值，然后后面元素全部前移

type RandomizedSet struct {
	m map[int]int	// map 键为元素， 值为元素在数组k中的下标
	k []int
	size int	// 用来标记实际数据的数量，不要也行，加上方便一点
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		k: make([]int, 0, 100),		// 先预设一个空间大小为100的数组
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = this.size
		this.k = append(this.k, val)
		this.size++	// 这句一定要在后面
		return true
	}
	return false	// 已存在
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.m[val]; ok {
		// 注意要先删数组中的，后删map中的
		this.m[this.k[this.size-1]] = idx	// 千万记得把原本末尾的这个元素的下标给修改
		this.k[idx], this.k[this.size-1] = this.k[this.size-1], this.k[idx]
		this.k = this.k[:this.size-1]
		this.size--
		delete(this.m, val)
		return true
	}
	return false	// 不存在
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	randi := rand.Intn(this.size)
	return this.k[randi]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
