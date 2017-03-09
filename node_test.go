package intervaltree

import (
	"fmt"
	"testing"
)

func TestNodeAdd(t *testing.T) {
	var nn Nodes
	if len(nn) != 0 {
		t.Fatalf("len(Nodes{}) should be 0: %d", len(nn))
	}
	nn.Add(0, 10)
	if len(nn) != 1 {
		t.Fatalf("should be 1: %d", len(nn))
	}
	nn.Add(2, 3)
	nn.Add(5, 10)
	if len(nn) != 3 {
		t.Fatalf("should be 3: %d", len(nn))
	}
}

func TestDummyData(t *testing.T) {
	var nn Nodes
	nn.Add(16, 21)
	nn.Add(8, 9)
	nn.Add(25, 30)
	nn.Add(5, 8)
	nn.Add(15, 23)
	nn.Add(17, 19)
	nn.Add(26, 26)
	nn.Add(0, 3)
	nn.Add(6, 10)
	nn.Add(19, 20)
	nn.Build()
	for i, x := range nn {
		fmt.Printf("#%-2d %2d %2d %2d %d\n",i, x.Start, x.End, x.Max, x.Index)
	}
	r := nn.QueryNodes(20)
	fmt.Printf("%+v\n", r);
}
