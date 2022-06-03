// package main

// import (
// 	"log"
// 	"runtime"

// 	"io/ioutil"

// 	"HelloWorldGL/shader"

// 	"github.com/go-gl/gl/v4.6-core/gl"
// 	"github.com/go-gl/glfw/v3.3/glfw"
// )

// const (
// 	width  = 800
// 	height = 600
// )

// var (
// 	triangle = []float32{
// 		0, 0.5, 0, 1.0, 0.0, 0.0, // top
// 		-0.5, -0.5, 0, 0.0, 1.0, 0.0, // left
// 		0.5, -0.5, 0, 0.0, 0.0, 1.0, // right
// 	}
// )

// // initGlfw initializes glfw and returns a Window to use.
// func initGlfw() *glfw.Window {
// 	if err := glfw.Init(); err != nil {
// 		panic(err)
// 	}

// 	glfw.WindowHint(glfw.Resizable, glfw.False)
// 	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
// 	glfw.WindowHint(glfw.ContextVersionMinor, 1)
// 	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
// 	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

// 	window, err := glfw.CreateWindow(width, height, "Hello World!", nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	window.MakeContextCurrent()

// 	return window
// }

// // initOpenGL initializes OpenGL and returns an intiialized program.
// func initOpenGL() uint32 {
// 	if err := gl.Init(); err != nil {
// 		panic(err)
// 	}
// 	version := gl.GoStr(gl.GetString(gl.VERSION))
// 	log.Println("OpenGL version", version)

// 	vShader, verr := ioutil.ReadFile("shader/vertexShader.glsl")
// 	vShaderStr := string(vShader) + "\x00"
// 	pShader, perr := ioutil.ReadFile("shader/fragmentShader.glsl")
// 	pShaderStr := string(pShader) + "\x00"

// 	if verr != nil || perr != nil {
// 		log.Fatal(verr, perr)
// 	}

// 	vertexShader, err := shader.CompileShader(vShaderStr, gl.VERTEX_SHADER)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fragmentShader, err := shader.CompileShader(pShaderStr, gl.FRAGMENT_SHADER)
// 	if err != nil {
// 		panic(err)
// 	}

// 	prog := gl.CreateProgram()
// 	gl.AttachShader(prog, vertexShader)
// 	gl.AttachShader(prog, fragmentShader)
// 	gl.LinkProgram(prog)
// 	return prog
// }

// func makeVao(points []float32) uint32 {
// 	var vbo uint32
// 	gl.GenBuffers(1, &vbo)
// 	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
// 	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

// 	var vao uint32
// 	gl.GenVertexArrays(1, &vao)
// 	gl.BindVertexArray(vao)

// 	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)
// 	gl.EnableVertexAttribArray(0)

// 	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)
// 	gl.EnableVertexAttribArray(1)

// 	return vao
// }

// func main() {
// 	runtime.LockOSThread()

// 	window := initGlfw()
// 	defer glfw.Terminate()

// 	program := initOpenGL()

// 	vao := makeVao(triangle)
// 	for !window.ShouldClose() {
// 		draw(vao, window, program)
// 	}
// }

// func draw(vao uint32, window *glfw.Window, program uint32) {
// 	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
// 	gl.UseProgram(program)

// 	gl.BindVertexArray(vao)
// 	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

// 	glfw.SwapInterval(1)
// 	glfw.WaitEvents()
// 	window.SwapBuffers()
// }
