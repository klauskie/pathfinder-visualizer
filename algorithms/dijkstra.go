package algorithms

import (
	"klauskie.com/pathfinder/src"
	"math"
)

type Dijkstra struct {
	Pf src.Pathfinder
}

func (c Dijkstra) Run() ([]src.Node, []src.Node) {
	exploration := dijkstra(c.Pf)
	path := dijkstraGetPath(c.Pf.GetEndNode())
	return exploration, path
}

func (c Dijkstra) GetPathfinder() src.Pathfinder {
	return c.Pf
}

func dijkstra(p src.Pathfinder) []src.Node {
	p.GetStartNode().Distance = 0
	var sortedVisitedNodes []src.Node
	unvisitedNodes := gridToList(p.Grid)
	for ; len(unvisitedNodes) > 0 ; {
		sortByDistance(unvisitedNodes)
		firstNode := unvisitedNodes[0]
		unvisitedNodes = unvisitedNodes[1:]

		if firstNode.Distance == int(math.MaxInt64) {
			break
		}

		if contains(p.Walls, firstNode.Id) {
			continue
		}

		firstNode.IsVisited = true
		sortedVisitedNodes = append(sortedVisitedNodes, *firstNode)
		if firstNode.Id == p.EndId {
			break
		}
		updateNeighbors(firstNode, p.Grid)
	}
	return sortedVisitedNodes
}

func dijkstraGetPath(endNode *src.Node) []src.Node {
	return buildPath(endNode)
}

func updateNeighbors(node *src.Node, grid [][]*src.Node) {
	neighbors := getNeighbors(node, grid)
	for i := 0; i < len(neighbors); i++ {
		neighbors[i].Distance = node.Distance + 1
		neighbors[i].PrevNode = node
	}
}

