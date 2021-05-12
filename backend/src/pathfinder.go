package src

import (
	"math"
)


type Pathfinder struct {
	StartId, EndId int
	Start, End Node
	Grid [][]*Node
	Walls []int
	Rows, Cols int
}

func CreatePathFinder(walls []int, startId, endId, rows, cols int) Pathfinder {
	p := Pathfinder{
		StartId:startId,
		EndId: endId,
		Grid:       nil,
		Walls: walls,
		Rows: rows,
		Cols: cols,
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
	for r := 0; r < p.Rows; r++ {
		grid = append(grid, make([]*Node, p.Cols))
		for c := 0; c < p.Cols; c++ {
			id := p.Cols * r + c
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
