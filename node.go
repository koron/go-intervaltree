package intervaltree

import (
	"sort"
)

type Node struct {
	Start int64
	End   int64
	Index int
	Max   int64
}

type Nodes []Node

func (nn *Nodes) Add(start, end int64) int {
	idx := len(*nn)
	*nn = append(*nn, Node{
		Index: idx,
		Start: start,
		End:   end,
	})
	return idx
}

func (nn Nodes) sort() {
	sort.Slice(nn, func(i, j int) bool {
		a, b := nn[i].Start, nn[j].Start
		switch {
		case a < b:
			return true
		case a > b:
			return false
		default:
			a, b = nn[i].End, nn[j].End
			return a < b
		}
	})
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func (nn Nodes) fillMax(off int) int64 {
	l := len(nn)
	mid := l / 2
	v := nn[mid].End
	if mid > 0 {
		v = max(v, nn[:mid].fillMax(off))
	}
	if mid < l-1 {
		v = max(v, nn[mid+1:].fillMax(off+mid+1))
	}
	nn[mid].Max = v
	return v
}

func (nn *Nodes) Build() {
	nn.sort()
	nn.fillMax(0)
}

func (nn Nodes) Query(q int64) []int {
	return nn.query(q, nil)
}

func (nn Nodes) query(q int64, r []int) []int {
	l := len(nn)
	mid := l / 2
	if q > nn[mid].Max {
		return r
	}
	if q >= nn[mid].Start && q <= nn[mid].End {
		r = append(r, nn[mid].Index)
	}
	if mid > 0 {
		r = nn[:mid].query(q, r)
	}
	if mid < l-1 {
		r = nn[mid+1:].query(q, r)
	}
	return r
}

func (nn Nodes) QueryNodes(q int64) []Node {
	return nn.queryNodes(q, nil)
}

func (nn Nodes) queryNodes(q int64, r []Node) []Node {
	l := len(nn)
	mid := l / 2
	if q > nn[mid].Max {
		return r
	}
	if q < nn[mid].Start {
		if mid > 0 {
			return nn[:mid].queryNodes(q, r)
		}
		return r
	}
	if q <= nn[mid].End {
		r = append(r, nn[mid])
	}
	if mid > 0 {
		r = nn[:mid].queryNodes(q, r)
	}
	if mid < l-1 {
		r = nn[mid+1:].queryNodes(q, r)
	}
	return r
}
