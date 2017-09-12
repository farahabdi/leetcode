package medianSortedArrays

//Same  length arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	//	A := []int{1, 12, 15, 26, 38}
	//	B := []int{2, 13, 17, 30, 45}

	result := getMedian(nums1, nums2, len(nums1))

	return result
}

func getMedian(A []int, B []int, n int) float64 {

	if n == 0 {
		return float64(A[0]+B[0]) / 2.0
	}

	if n == 2 {
		return float64((max(A[0], B[0]) + min(A[1], B[1])) / 2)
	}

	Am := median(A)
	Bm := median(B)

	if Am < Bm {
		if n%2 == 0 {
			return getMedian(A[n/2-1:], B[:n/2+1], n/2+1)
		}
		return getMedian(A[n/2:], B[:n/2+1], n/2+1)

	}

	if n%2 == 0 {
		return getMedian(A[:n/2+1], B[n/2-1:], n/2+1)
	}
	return getMedian(A[:n/2+1], B[n/2:], n/2+1)
}

func median(nums []int) float64 {

	n := len(nums)

	if n%2 == 0 {
		return (float64(nums[len(nums)/2.0-1]) + float64(nums[len(nums)/2])) / 2.0
	}

	return float64(nums[len(nums)/2.0])
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
