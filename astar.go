package main

import (
	"a-star/utils"
	"fmt"
	"math"
	"os"
)

var origin, dest utils.Point
var openList, closeList, path []utils.Point

// Set the origin point
func setOrig(s *Scene) {
	origin = utils.Point{utils.GetRandInt(s.rows-2) + 1, utils.GetRandInt(s.cols-2) + 1, 0, 0, 0, nil}
	if s.scene[origin.X][origin.Y] == ' ' {
		s.scene[origin.X][origin.Y] = 'A'
	} else {
		setOrig(s)
	}
}

// Set the destination point
func setDest(s *Scene) {
	dest = utils.Point{utils.GetRandInt(s.rows-2) + 1, utils.GetRandInt(s.cols-2) + 1, 0, 0, 0, nil}

	if s.scene[dest.X][dest.Y] == ' ' {
		s.scene[dest.X][dest.Y] = 'B'
	} else {
		setDest(s)
	}
}

// Init origin, destination. Put the origin point into the openlist by the way
func initAstar(s *Scene) {
	setOrig(s)
	setDest(s)
	openList = append(openList, origin)
}

func findPath(s *Scene) {
	current := getFMin()
	addToCloseList(current, s)
	walkable := getWalkable(current, s)
	for _, p := range walkable {
		addToOpenList(p)
	}
}

func getFMin() utils.Point {
	if len(openList) == 0 {
		fmt.Println("No way!!!")
		os.Exit(-1)
	}
	index := 0
	for i, p := range openList {
		if (i > 0) && (p.F <= openList[index].F) {
			index = i
		}
	}
	return openList[index]
}

func getWalkable(p utils.Point, s *Scene) []utils.Point {
	var around []utils.Point
	row, col := p.X, p.Y
	left := s.scene[row][col-1]
	up := s.scene[row-1][col]
	right := s.scene[row][col+1]
	down := s.scene[row+1][col]
	leftup := s.scene[row-1][col-1]
	rightup := s.scene[row-1][col+1]
	leftdown := s.scene[row+1][col-1]
	rightdown := s.scene[row+1][col+1]
	if (left == ' ') || (left == 'B') {
		around = append(around, utils.Point{row, col - 1, 0, 0, 0, &p})
	}
	if (leftup == ' ') || (leftup == 'B') {
		around = append(around, utils.Point{row - 1, col - 1, 0, 0, 0, &p})
	}
	if (up == ' ') || (up == 'B') {
		around = append(around, utils.Point{row - 1, col, 0, 0, 0, &p})
	}
	if (rightup == ' ') || (rightup == 'B') {
		around = append(around, utils.Point{row - 1, col + 1, 0, 0, 0, &p})
	}
	if (right == ' ') || (right == 'B') {
		around = append(around, utils.Point{row, col + 1, 0, 0, 0, &p})
	}
	if (rightdown == ' ') || (rightdown == 'B') {
		around = append(around, utils.Point{row + 1, col + 1, 0, 0, 0, &p})
	}
	if (down == ' ') || (down == 'B') {
		around = append(around, utils.Point{row + 1, col, 0, 0, 0, &p})
	}
	if (leftdown == ' ') || (leftdown == 'B') {
		around = append(around, utils.Point{row + 1, col - 1, 0, 0, 0, &p})
	}
	return around
}

func addToOpenList(p utils.Point) {
	updateWeight(&p)
	if checkExist(p, closeList) {
		return
	}
	if !checkExist(p, openList) {
		openList = append(openList, p)
	} else {
		if openList[findPoint(p, openList)].F > p.F { //New path found
			openList[findPoint(p, openList)].Parent = p.Parent
		}
	}
}

// Update G, H, F of the point
func updateWeight(p *utils.Point) {
	if checkRelativePos(*p) == 1 {
		p.G = p.Parent.G + 10
	} else {
		p.G = p.Parent.G + 14
	}
	absx := (int)(math.Abs((float64)(dest.X - p.X)))
	absy := (int)(math.Abs((float64)(dest.Y - p.Y)))
	p.H = (absx + absy) * 10
	p.F = p.G + p.H
}

func removeFromOpenList(p utils.Point) {
	index := findPoint(p, openList)
	if index == -1 {
		os.Exit(0)
	}
	openList = append(openList[:index], openList[index+1:]...)
}

func addToCloseList(p utils.Point, s *Scene) {
	removeFromOpenList(p)
	if (p.X == dest.X) && (p.Y == dest.Y) {
		generatePath(p, s)
		s.draw()
		os.Exit(1)
	}
	// if (p.Parent != nil) && (checkRelativePos(p) == 2) {
	// 	parent := p.Parent
	// 	//rdblck := s.scene[p.X][parent.Y] | s.scene[parent.X][p.Y]
	// 	//fmt.Printf("%c\n", rdblck)
	// 	if (s.scene[p.X][parent.Y] == '#') || (s.scene[parent.X][p.Y] == '#') {
	// 		return
	// 	}
	// }
	if s.scene[p.X][p.Y] != 'A' {
		s.scene[p.X][p.Y] = '·'
	}
	closeList = append(closeList, p)
}

func checkExist(p utils.Point, arr []utils.Point) bool {
	for _, point := range arr {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

func findPoint(p utils.Point, arr []utils.Point) int {
	for index, point := range arr {
		if p.X == point.X && p.Y == point.Y {
			return index
		}
	}

	return -1
}

func checkRelativePos(p utils.Point) int {
	parent := p.Parent
	hor := (int)(math.Abs((float64)(p.X - parent.X)))
	ver := (int)(math.Abs((float64)(p.Y - parent.Y)))
	return hor + ver
}

func generatePath(p utils.Point, s *Scene) {
	if (s.scene[p.X][p.Y] != 'A') && (s.scene[p.X][p.Y] != 'B') {
		s.scene[p.X][p.Y] = '*'
	}
	if p.Parent != nil {
		generatePath(*(p.Parent), s)
	}
}
