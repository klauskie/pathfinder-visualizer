package algorithms

import (
	"klauskie.com/pathfinder/backend/src"
	"math"
	"reflect"
	"sort"
)

func gridToList(grid [][]*src.Node) []*src.Node {
	var nodeList []*src.Node
	for _, list := range grid {
		for _, node := range list {
			nodeList = append(nodeList, node)
		}
	}
	return nodeList
}

func sortByDistance(nodes []*src.Node) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Distance < nodes[j].Distance
	})
}

func contains(hay []int, needle int) bool {
	for _, a := range hay {
		if a == needle {
			return true
		}
	}
	return false
}

func buildPath(node *src.Node) []src.Node {
	var path []src.Node
	currentNode := node
	for currentNode != nil{
		path = append([]src.Node{*currentNode}, path...)
		currentNode = currentNode.PrevNode
	}
	return path
}

func getNeighbors(node *src.Node, grid [][]*src.Node) []*src.Node {
	var neighbors []*src.Node
	if node.Row > 0 {
		n := grid[node.Row-1][node.Col]
		if !n.IsVisited {
			neighbors = append(neighbors, n)
		}
	}
	if node.Row < len(grid)-1 {
		n := grid[node.Row+1][node.Col]
		if !n.IsVisited {
			neighbors = append(neighbors, n)
		}
	}
	if node.Col > 0 {
		n := grid[node.Row][node.Col-1]
		if !n.IsVisited {
			neighbors = append(neighbors, n)
		}
	}
	if node.Col < len(grid[0])-1 {
		n := grid[node.Row][node.Col+1]
		if !n.IsVisited {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func calcEuclideanDistance(x1,y1,x2,y2 float64) int {
	return int(math.Sqrt(math.Pow(math.Abs(x1-x2), 2) + math.Pow(math.Abs(y1-y2), 2)))
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

type NodeList []src.Node
func (list NodeList) contains(needle src.Node) bool {
	for _, a := range list {
		if a.Id == needle.Id {
			return true
		}
	}
	return false
}

func (list NodeList) findById(needle int) *src.Node {
	for _, a := range list {
		if a.Id == needle {
			return &a
		}
	}
	return nil
}

type NodeRefList []*src.Node
func (list NodeRefList) contains(needle *src.Node) bool {
	for _, a := range list {
		if a.Id == needle.Id {
			return true
		}
	}
	return false
}

func (list NodeRefList) findById(needle int) *src.Node {
	for _, a := range list {
		if a.Id == needle {
			return a
		}
	}
	return nil
}
