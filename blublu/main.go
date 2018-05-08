package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/glow/gl"
)

func initVao(programID, buffer *uint32, vertexBufferData []float32) {
	*programID, _ = initOpengl()
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	gl.GenBuffers(1, buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, *buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexBufferData)*4, gl.Ptr(vertexBufferData), gl.STATIC_DRAW)
}

func draw(programID, buffer uint32, window *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(programID)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.DisableVertexAttribArray(0)
	window.SwapBuffers()
	glfw.PollEvents()
}

func main() {
	runtime.LockOSThread()

	var programID, buffer uint32
	vertexBufferData := []float32{
		-1.0, -1.0, 0.0,
		1.0, -1.0, 0.0,
		0.0, 1.0, 0.0,
	}

	window := initGlfw()
	defer glfw.Terminate()
	initVao(&programID, &buffer, vertexBufferData)

	for !window.ShouldClose() {
		draw(programID, buffer, window)
		//	time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}
