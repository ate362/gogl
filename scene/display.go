package scene

import (
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitGlfw(width int, height int, hint string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, hint, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func InitOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	// vShader, verr := ioutil.ReadFile("shader/vertexShader.glsl")
	// vShaderStr := string(vShader) + "\x00"
	// pShader, perr := ioutil.ReadFile("shader/fragmentShader.glsl")
	// pShaderStr := string(pShader) + "\x00"

	// if verr != nil || perr != nil {
	// 	log.Fatal(verr, perr)
	// }

	// vertexShader, err := shader.CompileShader(vShaderStr, gl.VERTEX_SHADER)
	// if err != nil {
	// 	panic(err)
	// }
	// fragmentShader, err := shader.CompileShader(pShaderStr, gl.FRAGMENT_SHADER)
	// if err != nil {
	// 	panic(err)
	// }

	// prog := gl.CreateProgram()
	// gl.AttachShader(prog, vertexShader)
	// gl.AttachShader(prog, fragmentShader)
	// gl.LinkProgram(prog)
	// return prog
}

func Clear(r float32, g float32, b float32, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}

func SwapBuffers(win *glfw.Window) {
	glfw.WaitEvents()
	glfw.SwapInterval(1)
	win.SwapBuffers()
}
