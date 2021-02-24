package src

import (
	"math"
)

const (
	cols = 50
	rows = 30
)

type Pathfinder struct {
	StartId, EndId int
	Start, End Node
	Grid [][]*Node
	Walls []int
}

func CreatePathFinder(walls []int, startId, endId int) Pathfinder {
	p := Pathfinder{
		StartId:startId,
		EndId: endId,
		Grid:       nil,
		Walls: walls,
	}
	p.InitGrid(true)
	return p
}

type Node struct {
	Id int
	Row, Col int
	IsStart, IsEnd, IsNormal bool
	IsVisited bool
	Distance int
	PrevNode *Node
}

func (p *Pathfinder) InitGrid(useInfinity bool) {
	var grid [][]*Node
	for r := 0; r < rows; r++ {
		grid = append(grid, make([]*Node, cols))
		for c := 0; c < cols; c++ {
			id := cols* r + c
			isStart := id == p.StartId
			isEnd := id == p.EndId
			distance := 0
			if useInfinity {
				distance = int(math.MaxInt64)
			}

			grid[r][c] = &Node{
				Id:        id,
				Row:       r,
				Col:       c,
				IsStart:   isStart,
				IsEnd:     isEnd,
				IsNormal:  !isStart && !isEnd,
				IsVisited: false,
				Distance: distance,
			}
		}
	}
	p.Grid = grid
}

func (p *Pathfinder) GetEndNode() *Node {
	for _, row := range p.Grid {
		for _, node := range row {
			if node.IsEnd {
				return node
			}
		}
	}
	return nil
}

func (p *Pathfinder) GetStartNode() *Node {
	for _, row := range p.Grid {
		for _, node := range row {
			if node.IsStart {
				return node
			}
		}
	}
	return nil
}
