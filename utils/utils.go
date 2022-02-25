package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Point - path point structure
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

// Clear - clear the scene
func Clear() {
	fmt.Printf("\033[100B")
	for i := 0; i < 100; i++ {
		fmt.Printf("\033[1A")
		fmt.Printf("\033[K")
	}
}

// GetRandInt - get random integer in range [0, limit)
func GetRandInt(limit int) int {
	return r.Intn(limit)
}
