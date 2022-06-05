package objects

import "github.com/go-gl/mathgl/mgl32"

type Vertex struct {
	position mgl32.Vec3
	texCoord mgl32.Vec2
	normal   mgl32.Vec3
	tangent  mgl32.Vec3
}
