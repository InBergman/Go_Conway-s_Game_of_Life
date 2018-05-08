package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/glow/gl"
)

const (
	width  = 1080
	heigth = 720

	rows    = 10
	columns = 10

	fps = 2
)

func initOpengl() uint32 {
	vertexShaderSource, vertexErr := ioutil.ReadFile("blublu/res/shaders/vertexShader.glsl")
	fragmentShaderSource, fragmentErr := ioutil.ReadFile("blublu/res/shaders/fragmentShader.glsl")

	if vertexErr != nil {
		panic(vertexErr)
	} else if fragmentErr != nil {
		panic(fragmentErr)
	}
	if err := gl.Init(); err != nil {
		panic(err)
	}

	vertexShader, err := compileShader(string(vertexShaderSource)+"\x00", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteShader(vertexShader)
	fragmentShader, err := compileShader(string(fragmentShaderSource)+"\x00", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteShader(fragmentShader)
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Opengl Version = ", version)
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	gl.ValidateProgram(prog)
	return prog
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, heigth, "My game", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}
