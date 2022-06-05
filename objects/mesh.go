package objects

import "github.com/go-gl/mathgl/mgl32"

type Vertex struct {
	Position mgl32.Vec3
	TexCoord mgl32.Vec2
	Normal   mgl32.Vec3
	Tangent  mgl32.Vec3
}
