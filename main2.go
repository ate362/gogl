package main

import (
	scene "WaterSimulation/scene"
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	runtime.LockOSThread()
	window := scene.InitGlfw(800, 600, "Water Simulation")
	camera := scene.InitCamera(0.0, 0.0, -25.0, 0.0, 1.0, 0.0, 0.0, 0.0)
	log.Println(camera)

	defer glfw.Terminate()

	scene.InitOpenGL()

	for !window.ShouldClose() {
		scene.Clear(0.0, 0.0, 0.0, 1.0)

		scene.SwapBuffers(window)
	}
}
