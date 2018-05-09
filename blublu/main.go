package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	runtime.LockOSThread()

	var programID, buffer uint32
	vertexBufferData := []float32{
		-0.2, -0.2, 0.0,
		0.2, -0.2, 0.0,
		0.0, 0.2, 0.0}

	window := initGlfw()
	defer glfw.Terminate()
	initVao(&programID, &buffer, vertexBufferData)

	for !window.ShouldClose() {
		draw(programID, buffer, window)
	}
}
