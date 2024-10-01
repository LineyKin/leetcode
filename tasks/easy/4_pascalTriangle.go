package easy

// https://leetcode.com/problems/pascals-triangle/

func GeneratePascalTriangle(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	prevTriangle := GeneratePascalTriangle(numRows - 1)
	newRow := make([]int, numRows)
	newRow[0] = 1
	newRow[numRows-1] = 1
	if numRows == 2 {
		return append(prevTriangle, newRow)
	}

	newRowBasis := make([]int, numRows-2)
	prevTriangleLastRow := prevTriangle[numRows-2]

	for i := 1; i < numRows-1; i++ {
		newRowBasis[i-1] = prevTriangleLastRow[i-1] + prevTriangleLastRow[i]
	}

	copy(newRow[1:], newRowBasis)

	return append(prevTriangle, newRow)
}
