package largeGroupPositions



import "testing"

func TestLargeGroupPositions(t *testing.T) {


	nums := "abbxxxxzzy" // 1
//	nums := "abcdddeeeeaabbbcd"
	//nums := "aaa"
	//nums := "babaaaabbb"	
	//nums := "ggggz"	
	//	nums := "aaaa"	
	result := largeGroupPositions(nums)

	if result != nil {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
