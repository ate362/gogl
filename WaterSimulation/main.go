package main

/*
Adapted from this tutorial: http://www.learnopengl.com/#!Lighting/Colors

Shows how the basic usage of color for 3D objects
*/

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"WaterSimulation/cam"
	"WaterSimulation/gfx"
	"WaterSimulation/objects"
	"WaterSimulation/waves"
	"WaterSimulation/win"
)

// vertices to draw 6 faces of a cube
var cubeVertices = []float32{
	// position        // texture position
	-1.0, -1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	// Top
	-1.0, 1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 1.0,

	// Left
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,

	// Right
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
}

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()

}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Samples, 16)

	window := win.NewWindow(1280, 720, "Water Simulation")

	// Initialize Glow (go function bindings)
	if err := gl.Init(); err != nil {
		panic(err)
	}

	err := programLoop(window)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
 * Creates the Vertex Array Object for a triangle.
 * indices is leftover from earlier samples and not used here.
 */
func createVAO(vertices []float32, indices []uint32) uint32 {

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)

	var EBO uint32
	gl.GenBuffers(1, &EBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 3*4 + 2*4
	var offset uintptr = 0

	// position

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, stride, offset)
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}

func programLoop(window *win.Window) error {

	waterVertShader, err := gfx.NewShaderFromFile("shaders/water.vert", gl.VERTEX_SHADER)
	if err != nil {
		return err
	}

	waterFragShader, err := gfx.NewShaderFromFile("shaders/water.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	// the linked shader program determines how the data will be rendered
	vertShader, err := gfx.NewShaderFromFile("shaders/basic.vert", gl.VERTEX_SHADER)
	if err != nil {
		return err
	}

	fragShader, err := gfx.NewShaderFromFile("shaders/basic.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	waterProgram, err := gfx.NewProgram(waterVertShader, waterFragShader)
	if err != nil {
		return err
	}
	defer waterProgram.Delete()

	program, err := gfx.NewProgram(vertShader, fragShader)
	if err != nil {
		return err
	}
	defer program.Delete()

	lightFragShader, err := gfx.NewShaderFromFile("shaders/light.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	//special shader program so that lights themselves are not affected by lighting
	lightProgram, err := gfx.NewProgram(vertShader, lightFragShader)
	if err != nil {
		return err
	}
	defer lightProgram.Delete()

	Water := objects.GenneratePlane(40, 40, 1000, 1000)
	VBO, VAO := objects.CreateMesh(Water)
	wg := waves.WavGenGPU(waterProgram)

	lightVAO := createVAO(cubeVertices, nil)

	// ensure that triangles that are "behind" others do not draw over top of them
	gl.Enable(gl.DEPTH_TEST)

	camera := cam.NewFpsCamera(mgl32.Vec3{0, 1, 9}, mgl32.Vec3{0, 1, 0}, -90, 0, window.InputManager())

	for !window.ShouldClose() {

		// swaps in last buffer, polls for window events, and generally sets up for a new render frame
		window.StartFrame()

		// update camera position and direction from input evevnts
		camera.Update(window.SinceLastFrame())

		// background color
		gl.ClearColor(0, 0, 0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // depth buffer needed for DEPTH_TEST

		// creates perspective
		fov := float32(60.0)
		projectTransform := mgl32.Perspective(mgl32.DegToRad(fov),
			float32(window.Width())/float32(window.Height()),
			0.1,
			100.0)

		camTransform := camera.GetTransform()
		lightPos := mgl32.Vec3{0, 8, -8}
		lightTransform := mgl32.Translate3D(lightPos.X(), lightPos.Y(), lightPos.Z()).Mul4(mgl32.Scale3D(.25, .25, .25))

		model := mgl32.Vec3{0, 0, 0}
		modelTransform := mgl32.Translate3D(model.X(), model.Y(), model.Z())

		inverseTranspose := modelTransform.Transpose().Inv()

		waterProgram.Use()
		gl.Uniform3fv(waterProgram.GetUniformLocation("lightPos"), 1, &lightPos[0])
		gl.Uniform3fv(waterProgram.GetUniformLocation("eyePosition"), 1, &camera.GetPosition()[0])
		gl.UniformMatrix4fv(waterProgram.GetUniformLocation("camera"), 1, false, &camTransform[0])
		gl.UniformMatrix4fv(waterProgram.GetUniformLocation("project"), 1, false, &projectTransform[0])
		gl.UniformMatrix4fv(waterProgram.GetUniformLocation("world"), 1, false, &modelTransform[0])
		gl.UniformMatrix4fv(waterProgram.GetUniformLocation("inverseTranspose"), 1, false, &inverseTranspose[0])

		gl.BindVertexArray(VAO)

		// draw each cube after all coordinate system transforms are bound

		//float32(time.Since(start).Minutes()
		wg.UpdateGPU(window.SinceLastFrame())
		objects.UpdateBuffer(Water, VBO)

		// obj is colored, light is white
		gl.Uniform3f(waterProgram.GetUniformLocation("objectColor"), 1.0, 0.5, 0.31)
		gl.Uniform3f(waterProgram.GetUniformLocation("lightColor"), 1.0, 1.0, 1.0)
		// gl.DrawArrays(gl.POINTS, 0, int32(len(windices)))

		gl.Enable(gl.MULTISAMPLE)
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

		//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		gl.DrawElements(gl.TRIANGLES, int32(len(Water.Indices)), gl.UNSIGNED_INT, nil)

		gl.BindVertexArray(0)

		gl.Disable(gl.BLEND)

		// Draw the light obj after the other boxes using its separate shader program
		// this means that we must re-bind any uniforms
		lightProgram.Use()
		gl.BindVertexArray(lightVAO)
		gl.UniformMatrix4fv(lightProgram.GetUniformLocation("world"), 1, false, &lightTransform[0])
		gl.UniformMatrix4fv(lightProgram.GetUniformLocation("camera"), 1, false, &camTransform[0])
		gl.UniformMatrix4fv(lightProgram.GetUniformLocation("project"), 1, false, &projectTransform[0])
		gl.DrawArrays(gl.TRIANGLES, 0, 36)

		// gl.BindVertexArray(0)

		// end of draw loop
	}

	return nil
}
