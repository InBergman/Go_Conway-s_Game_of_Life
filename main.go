package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/glow/gl"
)

func draw(vao uint32, window *glfw.Window, prog uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()

	prog := initOpengl()
	vao := makeVao(square)
	for !window.ShouldClose() {
		draw(vao, window, prog)
	}

}
