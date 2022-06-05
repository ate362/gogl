package objects

// vertexCount := m * n
// 	faceCount := 2 * (m - 1) * (n - 1)

// 	var waterVertices = []Vertex{}
// 	var waterIndices = []int{}

// 	waterVertices = waterVertices[0:vertexCount]
// 	waterIndices = waterIndices[0 : faceCount*3]

// 	halfWidth := 0.5 * float32(width)
// 	halfDepth := 0.5 * float32(depth)
// 	dx := float32(width) / float32(n-1)
// 	dz := float32(depth) / float32(m-1)
// 	du := 1.0 / float32(n-1)
// 	dv := 1.0 / float32(m-1)

// 	for i := 0; i < m; i++ {
// 		z := halfDepth - float32(i)*dz
// 		for j := 0; j < n; j++ {
// 			waterVertices[i*n+j] = Vertex{
// 				position: mgl32.Vec3{float32(i) / float32(m-1), 0, float32(j) / float32(n-1)},
// 				texCoord: mgl32.Vec2{float32(i) / float32(m-1), float32(j) / float32(n-1)},
// 				normal:   mgl32.Vec3{0, 1, 0},
// 				tangent:  mgl32.Vec3{1, 0, 0},
// 			}
// 		}
// 	}

func GenneratePlane(width int, depth int, m int, n int) ([]float32, []int32) {

	waterVertices := []float32{}
	waterIndices := []int32{}

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
			waterVertices = append(waterVertices, x)
			waterVertices = append(waterVertices, float32(0))
			waterVertices = append(waterVertices, z)

			waterVertices = append(waterVertices, float32(float32(j)*du))
			waterVertices = append(waterVertices, float32(float32(i)*dv))

			waterVertices = append(waterVertices, float32(0))
			waterVertices = append(waterVertices, float32(1))
			waterVertices = append(waterVertices, float32(0))

			waterVertices = append(waterVertices, float32(1))
			waterVertices = append(waterVertices, float32(0))
			waterVertices = append(waterVertices, float32(0))
		}
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			waterIndices = append(waterIndices, int32(i*n+j))
			waterIndices = append(waterIndices, int32(i*n+j+1))
			waterIndices = append(waterIndices, int32((i+1)*n+j))
			waterIndices = append(waterIndices, int32((i+1)*n+j))
			waterIndices = append(waterIndices, int32(i*n+j+1))
			waterIndices = append(waterIndices, int32((i+1)*n+j+1))
		}
	}

	return waterVertices, waterIndices
}