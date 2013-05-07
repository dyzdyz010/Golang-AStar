package main

import (
	//"a-star/utils"

	//"fmt"
	"time"
)

func main() {
	var scene Scene
	scene.initScene(23, 70)
	scene.addWalls(10)
	initAstar(&scene)

	for {
		//utils.Clear()
		findPath(&scene)
		scene.draw()
		time.Sleep(100 * time.Millisecond)
	}
}
