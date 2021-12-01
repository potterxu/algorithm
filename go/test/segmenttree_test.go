package algorithm_test

import (
	"testing"

	algorithm "github.com/potterxu/algorithm/go"
)

func TestIntervalSum(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var st algorithm.SegmentTree = &algorithm.IntervalSum{}
	st.Build(data)

	expect := 36
	actual, _ := st.Query(0, 7)
	if expect != actual {
		t.Fatalf("Expect %d, Actual %d", expect, actual)
	}

	expect = 20
	actual, _ = st.Query(1, 5)
	if expect != actual {
		t.Fatalf("Expect %d, Actual %d", expect, actual)
	}

	st.Update(3, 100)
	// {1, 2, 3, 100, 5, 6, 7, 8}

	expect = 116
	actual, _ = st.Query(1, 5)
	if expect != actual {
		t.Fatalf("Expect %d, Actual %d", expect, actual)
	}
}
