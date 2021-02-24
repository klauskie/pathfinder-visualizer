package algorithms

import (
	"fmt"
	"klauskie.com/pathfinder/src"
	"math"
	"sync"
)

var mainSortedVisitedNodes NodeList
var endSortedVisitedNodes NodeList

type BothWays struct {
	Pf src.Pathfinder
}

func (c BothWays) Run() ([]src.Node, []src.Node) {
	exploration := bothWays(c.Pf)
	path := bothWays(c.Pf)
	return exploration, path
}

func (c BothWays) GetPathfinder() src.Pathfinder {
	return c.Pf
}

func bothWays(p src.Pathfinder) NodeList {
	mainSortedVisitedNodes = NodeList{}
	endSortedVisitedNodes = NodeList{}

	wg := sync.WaitGroup{}
	wg.Add(2)
	q := p

	go func(p src.Pathfinder) {
		defer wg.Done()
		exploreStart(p, mainSortedVisitedNodes)
	}(p)

	go func(p src.Pathfinder) {
		defer wg.Done()
		exploreEnd(p, endSortedVisitedNodes)
	}(q)

	wg.Wait()

	fmt.Println(mainSortedVisitedNodes)
	fmt.Println(endSortedVisitedNodes)

	return joinPaths()
}

func joinPaths() NodeList {
	if len(endSortedVisitedNodes) == 0 || len(mainSortedVisitedNodes) == 0 {
		return NodeList{}
	}
	endTerminal := endSortedVisitedNodes[len(endSortedVisitedNodes)-1]
	startTerminal := mainSortedVisitedNodes.findById(endTerminal.Id)

	path := buildPath(startTerminal)
	endPath := buildPath(endTerminal.PrevNode)

	reverseAny(endPath)
	path = append(path, endPath...)
	return path
}

func exploreStart(p src.Pathfinder, sortedVisitedNodes NodeList) NodeList {
	p.GetStartNode().Distance = 0
	p.GetEndNode().Distance = int(math.MaxInt64)

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

		if endSortedVisitedNodes.contains(*firstNode) {
			break
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

func exploreEnd(p src.Pathfinder, sortedVisitedNodes NodeList) NodeList {
	p.GetEndNode().Distance = 0
	p.GetStartNode().Distance = int(math.MaxInt64)

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

		if mainSortedVisitedNodes.contains(*firstNode) {
			break
		}

		firstNode.IsVisited = true
		sortedVisitedNodes = append(sortedVisitedNodes, *firstNode)
		if firstNode.Id == p.StartId {
			break
		}
		updateNeighbors(firstNode, p.Grid)
	}
	return sortedVisitedNodes
}
