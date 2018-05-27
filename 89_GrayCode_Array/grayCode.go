package grayCode


func grayCode(num int) int  {

		n := uint(num)
	    // base case
		if n <= 0 {
			return 0
		}
	
		// 'arr' will store all generated codes
		arr := []string{}

		// start with one-bit pattern

		arr = append(arr, "0", "1")

		// Every iteration of this loop generates 2*i codes from previously
		// generated i codes.

		for i := 2; i < 1 << n; i = i<<1 {

				for j := i-1 ; j >= 0 ; j-- {
						arr = append(arr, arr[j])
					}
				// append 0 to the first half
				for j := 0 ; j < i ; j++ {
						arr[j] = "0" + arr[j]
					}

				// append 1 to the second half
				for j := i ; j < 2*i ; j++ {
						arr[j] = "1" + arr[j]
					}
		}
		return 0
}

/*
func grayCodew(n int) []int {


	s1 := []string{"0"}
	s2 := []string{"1"}

	for i :=0; i< 3; i++ {

		s1 = append([]string{"0"}, s1... )

		s2 = reverse(s1)
		s2 = append([]string{"0"}, s2... )

	
	}

	return nil
    
}

func reverse(s string ) string {

	r := []rune(s)
	end := len(s)

	for start :=0; start< end; start++ {
		r[start], r[end] = r[end], r[start]

		end++
	}

	return string(r)
}*/