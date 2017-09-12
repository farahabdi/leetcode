package maxProfitf

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {

	//stock := []int{2, 5, 7, 1, 4, 3, 1, 3}
	stock := []int{3, 3, 5, 0, 0, 3, 1, 4}
	result := maxProfit(stock)

	if result != 6 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}

}
