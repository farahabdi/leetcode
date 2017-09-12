package merge

import "fmt"

func mergeArray(nums1 []int, m int, nums2 []int, n int) {

	// 	array1 := []int{1, 3, 4, 5}
	//	array2 := []int{-1, 0, 4,  6, 8}

	//array1 := []int{1, 3, 4, 5}
	//array2 := []int{2, 4,  6, 8}
	//	n := min(m, n)

	//nums1 = append(nums1[:0], append(nums1, nums2[0:]...)...)
	nums1 = append(nums1[2:], nums2[0])
	fmt.Println(nums1)

	for i, j := 0, 0; i < n; i++ {
		if nums1[i] == nums2[j] {
			j++
			continue
		}

		if nums1[i] > nums2[j] {

			//nums2 = append(nums2[:i, )

		}
	}

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
