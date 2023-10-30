package arraystring

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	//gas := []int{1, 2, 3, 4, 5}
	//cost := []int{3, 4, 5, 1, 2}

	//gas := []int{3, 1, 1}
	//cost := []int{1, 2, 2}

	gas := []int{2}
	cost := []int{2}
	t.Logf("result:%d", canCompleteCircuit(gas, cost))
}
