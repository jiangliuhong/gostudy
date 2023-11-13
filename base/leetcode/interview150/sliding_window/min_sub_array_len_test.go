package sliding_window

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	//nums := []int{2, 3, 1, 2, 4, 3}
	//target := 7
	nums := []int{1, 2, 3, 4, 5}
	target := 11

	fmt.Println(minSubArrayLen(target, nums))
}
