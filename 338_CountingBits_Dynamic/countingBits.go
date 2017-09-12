package countingBits

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func countingBits(n int) int {
	count := 0
	for n > 0 {
		count = count + 1
		n = n & (n - 1)
	}

	return count
}
