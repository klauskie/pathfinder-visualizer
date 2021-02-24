package algorithms

import (
	"fmt"
	"klauskie.com/pathfinder/src"
	"sort"
)

type Astar struct {
	Pf src.Pathfinder
}

func (c Astar) Run() ([]src.Node, []src.Node) {
	path := astar(c.Pf)
	fmt.Println(path)
	return []src.Node{}, path
}

func (c Astar) GetPathfinder() src.Pathfinder {
	return c.Pf
}

func astar(p src.Pathfinder) []src.Node {
	openList := NodeRefList{p.GetStartNode()}
	closedList := NodeRefList{}
	path := NodeList{}

	for ; len(openList) > 0 ; {
		sortByLowestFscore(openList, *p.GetEndNode())
		currentNode := openList[0]
		openList = openList[1:]
		closedList = append(closedList, currentNode)
		currentNode.IsVisited = true

		if currentNode.Id == p.EndId {
			path = buildPath(currentNode)
			break
		}

		if contains(p.Walls, currentNode.Id) {
			continue
		}

		updateNeighbors(currentNode, p.Grid)

		for _, neighbor := range getNeighbors(currentNode, p.Grid) {
			if closedList.contains(neighbor) {
				continue
			}

			if addToOpen(openList, *neighbor, *p.GetEndNode()) {
				openList = append(openList, neighbor)
			}
		}
	}

	return path
}

func f(n, finalNode src.Node) int {
	g := n.Distance
	h := calcEuclideanDistance(float64(n.Col), float64(n.Row), float64(finalNode.Col), float64(finalNode.Row))
	return g + h
}

func addToOpen(openList NodeRefList, neighbor, endNode src.Node) bool {
	for _, node := range openList {
		if neighbor.Id == node.Id && f(neighbor, endNode) >= f(*node, endNode) {
			return false
		}
	}
	return true
}

func sortByLowestFscore(nodes []*src.Node, finalNode src.Node) {
	sort.Slice(nodes, func(i, j int) bool {
		return f(*nodes[i], finalNode) < f(*nodes[j], finalNode)
	})
}
