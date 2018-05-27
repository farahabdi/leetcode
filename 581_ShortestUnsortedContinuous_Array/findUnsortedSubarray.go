
package findUnsortedSubarray

import ("sort")
func findUnsortedSubarray(nums []int) int {
	snums := append([]int(nil), nums...)

	sort.Ints(snums);

	start := len(snums)
	end := 0;

	for i := 0; i < len(snums); i++ {
			if snums[i] != nums[i] {
					start = min(start, i);
					end = max(end, i);
			}
	}
	

	if end - start >= 0 {
		return end - start + 1
	} 

  return 0
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