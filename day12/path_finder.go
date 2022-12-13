package main

import (
	"errors"
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"math"
)

type PathFinder struct {
	grid            [][]rune
	start           Point
	end             Point
	unvisitedVertex *VertexQueue
	visited         map[Point]bool
	distances       map[Point]int
}

func NewPathFinder(grid [][]rune, start Point, end Point) *PathFinder {
	unvisited := NewVertexQueue()
	distances := make(map[Point]int)
	for y, row := range grid {
		for x, _ := range row {
			point := Point{
				X: x,
				Y: y,
			}
			if y == start.Y && x == start.X {
				distances[point] = 0
			} else {
				distances[point] = math.MaxInt
			}
		}
	}
	return &PathFinder{
		grid:            grid,
		start:           start,
		end:             end,
		unvisitedVertex: unvisited,
		distances:       distances,
		visited:         make(map[Point]bool),
	}
}

func (pf *PathFinder) ShortestPath() int {
	pf.unvisitedVertex.Enqueue(&QueueNode{
		X:        pf.start.X,
		Y:        pf.start.Y,
		Distance: 0,
	})
	for !pf.unvisitedVertex.Empty() {
		node, _ := pf.unvisitedVertex.Dequeue()
		point := nodeToPoint(node)
		currentDist := pf.distances[point]
		if currentDist < math.MaxInt {
			neighbours := pf.getNeighbours(point)
			for _, neighbour := range neighbours {
				if !pf.visited[neighbour] {
					newDistance := currentDist + 1
					if pf.distances[neighbour] > newDistance {
						pf.distances[neighbour] = newDistance
						neighNode := pointToNode(neighbour, newDistance)
						pf.unvisitedVertex.Enqueue(neighNode)
					}
				}
			}
		}
		pf.visited[point] = true
	}
	return pf.distances[pf.end]
}

func nodeToPoint(node *QueueNode) Point {
	return Point{
		X: node.X,
		Y: node.Y,
	}
}

func pointToNode(point Point, distance int) *QueueNode {
	return &QueueNode{
		X:        point.X,
		Y:        point.Y,
		Distance: distance,
	}
}

func (pf *PathFinder) getNeighbours(point Point) []Point {
	var neighbours []Point
	if point.X > 0 {
		if !pf.isTooHigh(point, point.X-1, point.Y) {
			neighbours = append(neighbours, Point{Y: point.Y, X: point.X - 1})
		}
	}
	if point.X < len(pf.grid[0])-1 {
		if !pf.isTooHigh(point, point.X+1, point.Y) {
			neighbours = append(neighbours, Point{Y: point.Y, X: point.X + 1})
		}
	}
	if point.Y > 0 {
		if !pf.isTooHigh(point, point.X, point.Y-1) {
			neighbours = append(neighbours, Point{Y: point.Y - 1, X: point.X})
		}
	}

	if point.Y < len(pf.grid)-1 {
		if !pf.isTooHigh(point, point.X, point.Y+1) {
			neighbours = append(neighbours, Point{Y: point.Y + 1, X: point.X})
		}
	}
	return neighbours
}

func (pf *PathFinder) isTooHigh(currentPosition Point, targetX int, targetY int) bool {
	targetHeight := pf.grid[targetY][targetX]
	return targetHeight-pf.grid[currentPosition.Y][currentPosition.X] > 1
}

type QueueNode struct {
	Y        int
	X        int
	Distance int
}

func (qn *QueueNode) CoordString() string {
	return fmt.Sprintf("(%v,%v)", qn.X, qn.Y)
}

type VertexQueue struct {
	pq *priorityqueue.Queue
}

func (vq *VertexQueue) Dequeue() (*QueueNode, error) {
	value, ok := vq.pq.Dequeue()
	if !ok {
		return nil, errors.New("queue empty")
	}
	return value.(*QueueNode), nil
}

func (vq *VertexQueue) Enqueue(node *QueueNode) {
	vq.pq.Enqueue(node)
}

func (vq *VertexQueue) Empty() bool {
	return vq.pq.Empty()
}

func NewVertexQueue() *VertexQueue {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		distA := a.(*QueueNode).Distance
		distB := b.(*QueueNode).Distance
		return utils.IntComparator(distA, distB)
	})
	return &VertexQueue{
		pq: pq,
	}
}
