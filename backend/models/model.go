package models

import (
	"klauskie.com/pathfinder/backend/algorithms"
	"klauskie.com/pathfinder/backend/src"
)

type PathfinderInterface interface {
	Run() ([]src.Node, []src.Node)
	GetPathfinder() src.Pathfinder
}

func Create(algoId int, walls []int, startId, endId, rows, cols int) PathfinderInterface {
	p := src.CreatePathFinder(walls, startId, endId, rows, cols)

	switch algoId {
	case 0:
		p.InitGrid(true)
		return algorithms.Dijkstra{Pf: p}
	case 1:
		p.InitGrid(false)
		return algorithms.Astar{Pf: p}
	case 2:
		p.InitGrid(true)
		return algorithms.BothWays{Pf: p}
	default:
		return nil
	}
}
