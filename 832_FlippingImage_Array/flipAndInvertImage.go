package flipAndInvertImage

func flipAndInvertImage(A [][]int) [][]int {

	for _, row := range A {
		reverseRows(row)
		flipBits(row)

	}

    return A
}

func reverseRows(row []int) {
	end := len(row)-1
	for start :=0; start < end; end-- {
			row[start], row[end] = row[end], row[start]
			start++
	}
}

func flipBits(row []int) {
	for start :=0; start < len(row); start++ {
			if row[start] == 0 {
				row[start] = 1
			} else {
				row[start] = 0
			}
	}
}