package ltsort

import (
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// 用来测试排序的一些辅助函数

// 生成随机数列
func generateRandomArray(n, min, max int) []int {
	arr := make([]int, n)
	for i:=0; i<n; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}



// 生成基本有序的数列
func generateNearlyOrderedArray(n, swaptimes int) []int {
	arr := make([]int, n)
	// 自然数序列
	for i:=0; i<n; i++ {
		arr[i] = i
	}
	// 交换
	for j:=0; j<swaptimes; j++ {
		posx := rand.Intn(n)	// 生成0~n-1的随机数，作为下标
		posy := rand.Intn(n)
		arr[posx], arr[posy] = arr[posy], arr[posx]
	}
	return arr
}

const (
	ASC = true
	DESC = false
)

// 判断数列是否升序有序序，判断我们的排序算法是否正确
func isSorted(arr []int, isASC bool) bool {
	n := len(arr)
	if isASC {
		for i:=0; i<n-1; i++ {
			if arr[i] > arr[i+1] {return false}
		}
	} else {
		for i:=0; i<n-1; i++ {
			if arr[i] < arr[i+1] {return false}
		}
	}
	return true
}


// 获取调用者函数名称信息。
func CallerName(skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	// skip输为0则输出本身函数信息; 这里 +1 是为了抵消CallerName自身的调用
	if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

// 测试算法性能，统计耗时
func testSort(arr []int, sort func([]int) []int, isASC bool) {
	t1 := time.Now()
	arr2 := sort(arr)
	d := time.Now().Sub(t1).Nanoseconds()
	sortValue := reflect.ValueOf(sort)
	sortPtr := sortValue.Pointer()
	sortName := runtime.FuncForPC(sortPtr)
	if !isSorted(arr2, isASC) {
		log.Fatalf("%s sort failed", sortName)
	}
	log.Printf("%s sort success, used %d ns", sortName, d)
}














