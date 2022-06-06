package objects

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const SizeofVertex = 3 * 2 * 3 * 3 * 4
const Stride = 3 + 2 + 3 + 3

type Vertex struct {
	Position mgl32.Vec3
	TexCoord mgl32.Vec2
	Normal   mgl32.Vec3
	Tangent  mgl32.Vec3
}

func CreateMesh(geo *Geometry) (uint32, uint32) {

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
	gl.BufferData(gl.ARRAY_BUFFER, len(geo.Vertices)*4, gl.Ptr(geo.Vertices), gl.DYNAMIC_DRAW)

	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(geo.Indices)*4, gl.Ptr(geo.Indices), gl.STATIC_DRAW)

	var offset uintptr = 0

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, Stride*4, offset)
	gl.EnableVertexAttribArray(0)

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VBO, VAO
}

func UpdateBuffer(geo *Geometry, vbo uint32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(geo.Vertices)*4, gl.Ptr(geo.Vertices))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func GetPosX(i int) int {
	return i
}

func GetPosY(i int) int {
	return i + 1
}

func GetPosZ(i int) int {
	return i + 2
}
