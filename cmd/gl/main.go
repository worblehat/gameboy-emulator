package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// TODO specifiying any version here leads to an error when calling gl.VertexAttribPointer*
	// glfw.WindowHint(glfw.ContextVersionMajor, 3)
	// glfw.WindowHint(glfw.ContextVersionMinor, 2)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(160, 144, "Game Boy", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Using OpenGL version", version)

	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	vertexShader := compileShader(vertexShaderSrc, gl.VERTEX_SHADER)
	fragmentShader := compileShader(fragmentShaderSrc, gl.FRAGMENT_SHADER)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	gl.UseProgram(program)

	v := []float32{
		0.0, 0.5,
		0.5, -0.5,
		-0.5, -0.5,
	}

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(v)*4, gl.Ptr(v), gl.STATIC_DRAW)
	positionAttrib := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))
	gl.VertexAttribPointerWithOffset(positionAttrib, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(positionAttrib)

	if glErr := gl.GetError(); glErr != gl.NO_ERROR {
		fmt.Printf("OpenGL Error: %x\n", glErr)
		os.Exit(1)
	}

	//TODO
	//glfw.SwapInterval(1)

	for !window.ShouldClose() {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(v)/2))
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func compileShader(src string, xtype uint32) (shader uint32) {
	shader = gl.CreateShader(xtype)
	cSrc, free := gl.Strs(src + "\x00")
	gl.ShaderSource(shader, 1, cSrc, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		fmt.Printf("Could not compile shader: %v\n%v", src, log)
		os.Exit(1)
	}
	return
}

var vertexShaderSrc = `
#version 150 core

in vec2 position;

void main()
{
	// Vertices are already in normalized device coordinates,
	// so no transformation needed here.
    gl_Position = vec4(position, 0.0, 1.0);
}
`
var fragmentShaderSrc = `
#version 150 core

out vec4 outColor;

void main()
{
    outColor = vec4(1.0, 0.0, 0.0, 1.0);
}
`
