package main

import (
	"time"
)

func main() {
	var scene Scene
	scene.initScene(23, 70)
	scene.addWalls(10)
	initAstar(&scene)

	for {
		findPath(&scene)
		scene.draw()
		time.Sleep(50 * time.Millisecond)
	}
}
