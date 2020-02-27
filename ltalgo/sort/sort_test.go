package ltsort

import (
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