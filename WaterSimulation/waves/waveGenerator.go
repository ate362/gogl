package waves

import (
	"WaterSimulation/gfx"
	"WaterSimulation/objects"
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const SIZE = 6

type WaveGeneratorGPU struct {
	wg          *WaveGenerator
	uTime       int32
	uNumWaves   int32
	uAmplitude  int32
	uWavelength int32
	uSpeed      int32
	uDirection  int32
}

type WaveGenerator struct {
	amplitude  [SIZE]float32
	wavelength [SIZE]float32
	speed      [SIZE]float32
	direction  [SIZE][2]float32
	time       float32
}

func WaveGen() *WaveGenerator {
	w := WaveGenerator{}

	for i := 0; i < SIZE; i++ {
		w.amplitude[i] = float32(0.25) / float32(i+1)
		w.wavelength[i] = float32(2.0*math.Pi) / float32(i+1)
		w.speed[i] = float32(1) + float32(2)*float32(i+1)
		v := mgl32.DegToRad(360.0 / float32(SIZE*(i+1)))
		w.direction[i] = mgl32.Vec2{float32(math.Cos(float64(v))), float32(math.Sin(float64(v)))}
	}

	return &w
}

func WavGenGPU(program *gfx.Program) *WaveGeneratorGPU {
	w := WaveGen()
	gwg := WaveGeneratorGPU{}

	gwg.wg = w
	gwg.uTime = gl.GetUniformLocation(program.Handle, gl.Str("waterDeformer.time\x00"))
	gwg.uNumWaves = gl.GetUniformLocation(program.Handle, gl.Str("waterDeformer.numWaves\x00"))
	gwg.uAmplitude = gl.GetUniformLocation(program.Handle, gl.Str("waterDeformer.amplitude\x00"))
	gwg.uWavelength = gl.GetUniformLocation(program.Handle, gl.Str("waterDeformer.wavelength\x00"))
	gwg.uSpeed = gl.GetUniformLocation(program.Handle, gl.Str("waterDeformer.speed\x00"))
	gwg.uDirection = gl.GetUniformLocation(program.Handle, gl.Str("waterDeformer.direction\x00"))

	return &gwg
}

func (w *WaveGenerator) waveHeight(x float32, z float32) float32 {
	var height float32
	for i := 0; i < SIZE; i++ {
		frequency := float32(2.0*math.Pi) / w.wavelength[i]
		phase := w.speed[i] * frequency
		// theta := w.direction[i].Dot(mgl32.Vec2{x, z})
		theta := -mgl32.Vec2{x, z}.Sub(mgl32.Vec2{0, 0}).Normalize().Dot(mgl32.Vec2{x, z})

		height += w.amplitude[i] * float32(math.Sin(float64(theta*frequency+w.time*phase)))
	}
	return height
}

func (w *WaveGenerator) UpdateCPU(time float64, geo *objects.Geometry) {
	w.time += float32(time / 10)
	for i := 0; i < len(geo.Vertices); i += objects.Stride {
		geo.Vertices[i+1] = w.waveHeight(geo.Vertices[i], geo.Vertices[i+2])
	}
}

func (gwg *WaveGeneratorGPU) UpdateGPU(time float64) {
	gwg.wg.time = gwg.wg.time + float32(time/10)
	gl.Uniform1f(gwg.uTime, gwg.wg.time)
	gl.Uniform1i(gwg.uNumWaves, SIZE)
	gl.Uniform1fv(gwg.uAmplitude, SIZE, &(gwg.wg.amplitude)[0])
	gl.Uniform1fv(gwg.uWavelength, SIZE, &(gwg.wg.wavelength)[0])
	gl.Uniform1fv(gwg.uSpeed, SIZE, &(gwg.wg.speed)[0])
	gl.Uniform2fv(gwg.uDirection, SIZE, &(gwg.wg.direction)[0][0])
}
