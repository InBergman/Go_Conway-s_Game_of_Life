package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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

func initOpengl() (uint32, error) {
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
	programID := gl.CreateProgram()
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Opengl Version = ", version)
	gl.AttachShader(programID, vertexShader)
	gl.AttachShader(programID, fragmentShader)
	gl.LinkProgram(programID)

	gl.DetachShader(programID, vertexShader)
	gl.DetachShader(programID, fragmentShader)

	var status int32
	gl.GetProgramiv(programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(programID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(programID, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to link program: %v", log)
	}
	return programID, nil
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
