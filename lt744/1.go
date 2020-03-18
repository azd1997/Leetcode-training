package lt744

// 寻找比目标字母大的最小字母

// 方法1：用[26]bool作哈希表（用map也行）记录letter中字母的存在情况。
//       从target的右邻位开始找看存不存在，第一个存在的就是答案
// 方法2： 线性扫描letters，找比目标字母大的第一个字母，返回它；
//		 如果遍历一遍没找到比target大的，则返回letters[0]
// 方法3： 二分查找:要找出target的右邻，而letters是排好序的，那么实际上相当于：
//		 将target插入letters，将会插入在哪？我们要找的是target右邻
// 		 要注意使用labuladong二分两端闭区间模板时，查找target右边界，如果target不存在，最后
// 		 r就会落在第一个比target小的字母。不管存不存在，最后返回的都是r+1对应的那个字母，如果r+1
// 注意，最后如果r+1越界了，需要 模len(letters) 一下

// 二分查找
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r, mid := 0, len(letters)-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return letters[(r+1)%len(letters)]
}

// 尽管letters是循环的，可能有多个target，
// 但是假定最后一个出现的target才是target的位置，我们要找的就是最后一个target的右邻
