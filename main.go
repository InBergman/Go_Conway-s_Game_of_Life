package main

import (
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/glow/gl"
)

func draw(cells [][]*cell, window *glfw.Window, prog uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)
	for x := range cells {
		for _, c := range cells[x] {
			c.draw()
		}
	}
	glfw.PollEvents()
	window.SwapBuffers()
}

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()

	prog := initOpengl()
	cells := makeCells()
	//vao := makeVao(square)
	for !window.ShouldClose() {
		t := time.Now()
		draw(cells, window, prog)
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}
