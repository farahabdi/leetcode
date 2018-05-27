package jumpGame

const (  // iota is reset to 0
	GOOD = iota  // c0 == 0
	BAD = iota  // c1 == 1
	UNKNOWN = iota  // c2 == 2
)

func canJumpFromPosition1(position int, nums []int, memo []int) bool {
	
	if memo[position] != UNKNOWN {

		if memo[position] == GOOD {
			return true
		} else {
			return false
		}
	}

	furthestJump := min(position + nums[position], len(nums) - 1)
	for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
			if canJumpFromPosition1(nextPosition, nums, memo) {
					memo[position] = GOOD;
					return true
			}
	}

	memo[position] = BAD
	return false
}

func canJumpBacktrack(nums []int) bool {

	memo := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
			memo[i] = UNKNOWN
	}
	memo[len(nums) - 1] = GOOD
	return canJumpFromPosition1(0, nums, memo);
}
