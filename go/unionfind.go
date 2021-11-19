package algorithm

import "fmt"

type UnionFindType struct {
	vToNodeMap map[int]*unionFindNodeType
}

func UnionFind() *UnionFindType {
	return &UnionFindType{
		vToNodeMap: make(map[int]*unionFindNodeType),
	}
}

// Add one node
func (uf *UnionFindType) Add(v int) error {
	if uf.vToNodeMap[v] != nil {
		return fmt.Errorf("value %d already existed in UnionFind", v)
	}
	uf.vToNodeMap[v] = unionFindNode(v)
	uf.vToNodeMap[v].root = uf.vToNodeMap[v]
	return nil
}

func (uf *UnionFindType) Find(v int) (int, error) {
	if uf.vToNodeMap[v] == nil {
		return -1, fmt.Errorf("value %d not found in UnionFind", v)
	}
	return uf.find(uf.vToNodeMap[v]).val, nil
}

func (uf *UnionFindType) Union(a, b int) {
	if uf.vToNodeMap[a] == nil {
		uf.Add(a)
	}
	if uf.vToNodeMap[b] == nil {
		uf.Add(b)
	}
	uf.union(uf.vToNodeMap[a], uf.vToNodeMap[b])
}

// union by rank
func (uf *UnionFindType) union(a, b *unionFindNodeType) {
	aRoot := uf.find(a)
	bRoot := uf.find(b)
	if aRoot != bRoot {
		var small, large *unionFindNodeType
		if aRoot.rank >= bRoot.rank {
			small = bRoot
			large = aRoot
		} else {
			small = aRoot
			large = bRoot
		}
		small.root = large
		if small.rank == large.rank {
			large.rank++
		}
	}
}

func (uf *UnionFindType) find(node *unionFindNodeType) *unionFindNodeType {
	if node == node.root {
		return node
	}
	node.root = uf.find(node.root)
	return node.root
}

type unionFindNodeType struct {
	val  int
	root *unionFindNodeType
	rank int
}

func unionFindNode(val int) *unionFindNodeType {
	return &unionFindNodeType{
		val:  val,
		root: nil,
		rank: 0,
	}
}
