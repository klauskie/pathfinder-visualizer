package models

import (
	"klauskie.com/pathfinder/algorithms"
	"klauskie.com/pathfinder/src"
)

type PathfinderInterface interface {
	Run() ([]src.Node, []src.Node)
	GetPathfinder() src.Pathfinder
}

func Create(algoId int, walls []int, startId, endId int) PathfinderInterface {
	p := src.CreatePathFinder(walls, startId, endId)

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
