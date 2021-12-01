package algorithm

import "fmt"

type SegmentTree interface {
	Build(data []int)
	Query(l, r int) (int, error)
	Update(pos, val int) error
}

type IntervalSum struct {
	tree []int
	num  int
}

func leftChild(n int) int {
	return 2 * n
}

func rightChild(n int) int {
	return 2*n + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func parent(n int) int {
// 	return n / 2
// }

func (st *IntervalSum) Build(data []int) {
	st.num = len(data)
	st.tree = make([]int, 4*st.num)
	st.build(1, data, 0, st.num-1)
}

func (st *IntervalSum) Query(l, r int) (int, error) {
	if l < 0 || r >= st.num {
		return -1, fmt.Errorf("invalid range [%v, %v]", l, r)
	}
	return st.query(1, 0, st.num-1, l, r), nil
}

func (st *IntervalSum) Update(i, val int) error {
	if i < 0 || i >= st.num {
		return fmt.Errorf("invalid index %v", i)
	}
	st.update(1, 0, st.num-1, i, val)
	return nil
}

func (st *IntervalSum) build(pos int, data []int, l, r int) {
	if l == r {
		st.tree[pos] = data[l]
	} else {
		m := l + (r-l)/2
		st.build(leftChild(pos), data, l, m)
		st.build(rightChild(pos), data, m+1, r)
		st.tree[pos] = st.tree[leftChild(pos)] + st.tree[rightChild(pos)]
	}
}

func (st *IntervalSum) query(pos, il, ir, l, r int) int {
	if l > r {
		return 0
	} else if il == l && ir == r {
		return st.tree[pos]
	} else {
		im := il + (ir-il)/2
		return st.query(leftChild(pos), il, im, l, min(r, im)) + st.query(rightChild(pos), im+1, ir, max(l, im+1), r)
	}
}

func (st *IntervalSum) update(pos, l, r, i, v int) {
	if l == r {
		st.tree[pos] = v
	} else {
		m := l + (r-l)/2
		if i <= m {
			st.update(leftChild(pos), l, m, i, v)
		} else {
			st.update(rightChild(pos), m+1, r, i, v)
		}
		st.tree[pos] = st.tree[leftChild(pos)] + st.tree[rightChild(pos)]
	}
}
