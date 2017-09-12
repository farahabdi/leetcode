package rangeSumQuery

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {

	return NumArray{nums: nums}

}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum = this.nums[k] + sum
	}

	return sum

}

/**∂∂
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
