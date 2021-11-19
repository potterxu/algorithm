package algorithm_test

import (
	"testing"

	algorithm "github.com/potterxu/algorithm/go"
)

func TestUnionFind(t *testing.T) {
	uf := algorithm.UnionFind()
	for i := 0; i <= 10; i++ {
		uf.Add(i)
	}

	// Union 1
	uf.Union(0, 1)
	uf.Union(0, 2)
	uf.Union(0, 3)
	uf.Union(2, 5)
	uf.Union(4, 5)

	// Union 2
	uf.Union(6, 7)
	uf.Union(8, 9)
	uf.Union(8, 10)
	uf.Union(9, 7)

	for i := 0; i < 6; i++ {
		expect := 0
		if actual, ok := uf.Find(i); ok != nil || actual != expect {
			t.Fatalf("Root of %d: Expect %d, Actual %d", i, expect, actual)
		}
	}
	for i := 6; i < 11; i++ {
		expect := 8
		if actual, ok := uf.Find(i); ok != nil || actual != expect {
			t.Fatalf("Root of %d: Expect %d, Actual %d", i, expect, actual)
		}
	}

	// union Union1 and Union2
	uf.Union(3, 9)
	for i := 0; i < 11; i++ {
		expect := 8
		if actual, ok := uf.Find(i); ok != nil || actual != expect {
			t.Fatalf("Root of %d: Expect %d, Actual %d", i, expect, actual)
		}
	}
}
