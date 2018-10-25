package calculate_test

import (
	"calculate"
	"testing"
)

func TestSum(t *testing.T) {
	s := calculate.Sum(1, 2, 3)

	if s != 6 {
		t.Error("Wrong result")
	}
}
