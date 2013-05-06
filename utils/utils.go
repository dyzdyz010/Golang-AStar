package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Point struct {
	X, Y    int
	H, G, F int
	Parent  *Point
}

func (p Point) String() string {
	return "[" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + ", " + strconv.Itoa(p.F) + "]"
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Clear() {
	fmt.Printf("\033[100B")
	for i := 0; i < 100; i++ {
		fmt.Printf("\033[1A")
		fmt.Printf("\033[K")
	}
}

func GetRandInt(limit int) int {
	return r.Intn(limit)
}
