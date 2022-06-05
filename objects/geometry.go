package objects

type Geometry struct {
	Vertices     []float32
	Indices      []int32
	SizeOfVertex int
	Offset       int
}

func GenneratePlane(width int, depth int, m int, n int) *Geometry {

	vertices := []float32{}
	indices := []int32{}

	halfWidth := 0.5 * float32(width)
	halfDepth := 0.5 * float32(depth)
	dx := float32(width) / float32(n-1)
	dz := float32(depth) / float32(m-1)
	du := 1.0 / float32(n-1)
	dv := 1.0 / float32(m-1)

	for i := 0; i < m; i++ {
		z := halfDepth - float32(i)*dz
		for j := 0; j < n; j++ {
			x := -halfWidth + float32(j)*dx
			vertices = append(vertices, x)
			vertices = append(vertices, float32(0))
			vertices = append(vertices, z)

			vertices = append(vertices, float32(float32(j)*du))
			vertices = append(vertices, float32(float32(i)*dv))

			vertices = append(vertices, float32(0))
			vertices = append(vertices, float32(1))
			vertices = append(vertices, float32(0))

			vertices = append(vertices, float32(1))
			vertices = append(vertices, float32(0))
			vertices = append(vertices, float32(0))
			// vertices = append(vertices, Vertex{
			// 	mgl32.Vec3{x, float32(0), z},
			// 	mgl32.Vec2{float32(j) * du, float32(i) * dv},
			// 	mgl32.Vec3{0, 1, 0},
			// 	mgl32.Vec3{1, 0, 0}})
		}
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			indices = append(indices, int32(i*n+j))
			indices = append(indices, int32(i*n+j+1))
			indices = append(indices, int32((i+1)*n+j))
			indices = append(indices, int32((i+1)*n+j))
			indices = append(indices, int32(i*n+j+1))
			indices = append(indices, int32((i+1)*n+j+1))
		}
	}

	geo := Geometry{vertices, indices, 4, 3*4 + 2*4 + 3*4 + 3*4}

	return &geo
}
