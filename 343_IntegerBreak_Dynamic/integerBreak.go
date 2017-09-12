package integerBreak

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func integerBreak(n int) int {
	count := 0
	for n > 0 {
		count = count + 1
		n = n & (n - 1)
	}

	return count
}
