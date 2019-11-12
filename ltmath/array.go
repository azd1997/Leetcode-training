package ltmath

func TwoDimZeroArray(row, col int) [][]int {
	array := make([]int, col)
	var twoDimArray [][]int
	for i:=0; i<row; i++ {
		twoDimArray = append(twoDimArray, array)
	}
	return twoDimArray
}

func TwoDimBoolArray(row, col int) [][]bool {
	array := make([]bool, col)
	var twoDimArray [][]bool
	for i:=0; i<row; i++ {
		twoDimArray = append(twoDimArray, array)
	}
	return twoDimArray
}