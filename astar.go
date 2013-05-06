package main

import (
	"a-star/utils"
	"fmt"
	"math"
	"os"
)

var origin, dest utils.Point
var openList, closeList, path []utils.Point

var openCursor, closeCursor int

// Set the origin point
func setOrig(s *Scene) {
	origin = utils.Point{utils.GetRandInt(s.rows-2) + 1, utils.GetRandInt(s.cols-2) + 1, 0, 0, 0, nil}
	if s.scene[origin.X][origin.Y] == ' ' {
		s.scene[origin.X][origin.Y] = 'A'
	} else {
		setOrig(s)
	}
	//fmt.Println("Origin:", origin.X, origin.Y)
}

// Set the destination point
func setDest(s *Scene) {
	dest = utils.Point{utils.GetRandInt(s.rows-2) + 1, utils.GetRandInt(s.cols-2) + 1, 0, 0, 0, nil}

	if s.scene[dest.X][dest.Y] == ' ' {
		s.scene[dest.X][dest.Y] = 'B'
	} else {
		setDest(s)
	}
	//fmt.Println("Destination:", dest.X, dest.Y)
}

func addToOpenList(p utils.Point) {
	updateWeight(&p)
	if checkExist(p, closeList) {
		return
	}
	if !checkExist(p, openList) {
		openList = append(openList, p)
		openCursor++
	} else {
		if openList[findPoint(p, openList)].F > p.F { //New path found
			openList[findPoint(p, openList)].Parent = p.Parent
		}
	}
}

// Update G, H, F of the point
func updateWeight(p *utils.Point) {
	if checkRelativePos(*p) == 0 {
		p.G = p.Parent.G + 10
	} else {
		p.G = p.Parent.G + 14
	}
	absx := (int)(math.Abs((float64)(dest.X - p.X)))
	absy := (int)(math.Abs((float64)(dest.Y - p.Y)))
	p.H = (absx + absy) * 10
	p.F = p.G + p.H
}

func findPoint(p utils.Point, arr []utils.Point) int {
	for index, point := range arr {
		if p.X == point.X && p.Y == point.Y {
			return index
		}
	}

	return -1
}

func removeFromOpenList(p utils.Point) {
	//fmt.Println("Point wait to be removed from openList:", p)
	index := findPoint(p, openList)
	if index == -1 {
		//fmt.Println("Fatal error occured.")
		os.Exit(0)
	}
	openList = append(openList[:index], openList[index+1:]...)
}

func addToCloseList(p utils.Point, s *Scene) {
	//fmt.Println(p.F)
	if (p.X == dest.X) && (p.Y == dest.Y) {
		generatePath(p, s)
		s.draw()
		//fmt.Println("Path generation complete.")
		//fmt.Println(closeList)
		os.Exit(1)
	}
	if s.scene[p.X][p.Y] != 'A' {
		s.scene[p.X][p.Y] = '·'
	}
	//fmt.Println(p.X, p.Y)
	removeFromOpenList(p)
	closeList = append(closeList, p)
	closeCursor++
}

func initLists(s *Scene) {
	//openList, closeList = make([]utils.Point, 1000), make([]utils.Point, 1000)
	openCursor, closeCursor = 0, 0
	openList = append(openList, origin)
	openCursor++
}

func findPath(s *Scene) {
	current := getZMin()
	//fmt.Println("Current:", current)
	addToCloseList(current, s)
	walkable := getWalkable(current, s)
	for _, p := range walkable {
		addToOpenList(p)
	}
	//fmt.Printf("Open List: %v\n", openList)
	//fmt.Println("This is the best:", p)
	//fmt.Println("OpenList:", openList)
	//fmt.Println("CloseList:", closeList)
}

func getZMin() utils.Point {
	if len(openList) == 0 {
		fmt.Println("No way!!!")
		os.Exit(-1)
	}
	index := 0
	for i, p := range openList {
		//fmt.Println(p)
		if (i > 0) && (p.F <= openList[index].F) {
			index = i
			//fmt.Printf("Find best at index %d, F: %d\n", index, openList[index].F)
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

func checkExist(p utils.Point, arr []utils.Point) bool {
	for _, point := range arr {
		if p.X == point.X && p.Y == point.Y {
			//fmt.Println("Found exist:", p)
			return true
		}
	}
	return false
}

func checkRelativePos(p utils.Point) int {
	parent := *(p.Parent)
	hor := (int)(math.Abs((float64)(p.X - parent.X)))
	ver := (int)(math.Abs((float64)(p.Y - parent.Y)))
	return hor + ver - 1
}

func generatePath(p utils.Point, s *Scene) {
	if (s.scene[p.X][p.Y] != 'A') && (s.scene[p.X][p.Y] != 'B') {
		s.scene[p.X][p.Y] = '*'
	}
	if p.Parent != nil {
		generatePath(*(p.Parent), s)
	}
}
