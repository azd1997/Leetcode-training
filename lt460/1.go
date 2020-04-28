package lt460

// LFU缓存

// 题目要求操作的时间复杂度为O(1)
// 我已开始直接想到的是优先队列与链表（处理相同频次的键值对）的组合。但是复杂度是O(lgn)
// 常见的O(1)实现方法：
// 1. 哈希表 + 平衡二叉树
// 2. 双哈希表
// 由于go没有树的API，自己写太麻烦
// 这里使用双哈希表实现

type LFUCache struct {
}

func Constructor(capacity int) LFUCache {

}

func (this *LFUCache) Get(key int) int {

}

func (this *LFUCache) Put(key int, value int) {

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
