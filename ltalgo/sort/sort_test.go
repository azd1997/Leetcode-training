package ltsort

import (
	"fmt"
	"math"
	"testing"
)

// 测试，每次都使用较大的两个数组进行测试，一个是随机数组，一个是接近有序的数组

var (
	testArrLength = 1000000
	minnum = 0
	maxnum = math.MaxInt32
	testSwapTimes = 200

)


func TestBubbleSort(t *testing.T) {
	arr1 := generateRandomArray(testArrLength, minnum, maxnum)
	arr2 := generateNearlyOrderedArray(testArrLength, testSwapTimes)

	testSort(arr1, BubbleSort, true)
	testSort(arr2, BubbleSort, true)
}

func TestBubbleSort_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(BubbleSort)
}

func TestSelectionSort_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(SelectionSort)
}

func TestInsertSort_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(InsertSort)
}

func TestShellSort_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(ShellSort)
}

func TestMergeSort_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(MergeSort)
}

func TestQuickSort1_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(QuickSort(QuickSortNormal))
}

func TestQuickSort2_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(QuickSort(QuickSort2Way))
}

func TestQuickSort3_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(QuickSort(QuickSort3Way))
}

func TestHeapSort1_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(HeapSort(HeapSort1))
}

func TestHeapSort2_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(HeapSort(HeapSortHeapify))
}

func TestHeapSort3_SmallDataSet(t *testing.T) {
	testSort_SmallDataSet(HeapSort(HeapSortInplace))
}

func TestIndexMaxHeap_Insert_Remove(t *testing.T) {
	arr := generateRandomArray(20, 0, 20)
	fmt.Println(arr)

	// 插入测试： 通过
	fmt.Println("插入测试")
	heap := NewIndexMaxHeap(20)
	for i:=0; i<20; i++ {
		heap.Insert(i, arr[i])
	}
	fmt.Println(heap.data)
	fmt.Println(heap.indexes)

	// 删除操作： 通过
	fmt.Println("删除操作")
	arr2 := make([]int, 20)
	for i:=0; i<20; i++ {
		arr2[i] = heap.RemoveMax()
	}

	fmt.Println(arr2)
	// 输出应该为降序
	fmt.Println(isSorted(arr2, false))
}

func TestIndexMaxHeap_RemoveMaxAndReturnIndex_GetItem(t *testing.T) {
	arr := generateRandomArray(20, 0, 20)
	fmt.Println(arr)

	// 删除并返回最大索引: 先构建堆，然后删除并返回索引
	// 再用索引去访问data看是否是降序的
	fmt.Println("删除并返回索引、根据索引取值测试")
	heap2 := NewIndexMaxHeap(20)
	for i:=0; i<20; i++ {
		heap2.Insert(i, arr[i])
	}
	fmt.Println("data： ", heap2.data)
	fmt.Println("indexes： ", heap2.indexes)


	arr3 := make([]int, 20)
	idxs := make([]int, 20)
	for i:=0; i<20; i++ {
		idxs[i] = heap2.RemoveMaxAndReturnIndex()
		arr3[i] = heap2.GetItem(idxs[i])
	}
	fmt.Println("idxs", idxs)
	fmt.Println(arr3)
	fmt.Println(isSorted(arr3, false))

	fmt.Println("data： ", heap2.data)
	fmt.Println("indexes： ", heap2.indexes)

	arr4 := make([]int, 20)
	for i:=19; i>=0; i-- {
		arr4[i] = heap2.data[heap2.indexes[i]]
	}

	fmt.Println(arr4)
	fmt.Println(isSorted(arr4, true))
}

func TestIndexMaxHeap_Change(t *testing.T) {
	arr := generateRandomArray(20, 0, 20)
	fmt.Println(arr)
	fmt.Println("change测试")
	heap := NewIndexMaxHeap(20)
	for i:=0; i<20; i++ {
		heap.Insert(i, arr[i])
	}

	heap.Change(7, 80)
	fmt.Println(heap)
}