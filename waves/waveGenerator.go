package waves

import (
	"WaterSimulation/objects"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const SIZE = 6

type WaveGenerator struct {
	amplitude  [SIZE]float32
	wavelength [SIZE]float32
	speed      [SIZE]float32
	direction  [SIZE]mgl32.Vec2
}

func WaveGen() *WaveGenerator {
	w := WaveGenerator{}

	for i := 0; i < SIZE; i++ {
		w.amplitude[i] = float32(0.25) / float32(i+1)
		w.wavelength[i] = float32(2*math.Pi) / float32(i+1)
		w.speed[i] = float32(.01) + float32(2*math.Pi)*float32(i+1)
		v := mgl32.DegToRad(360.0 / float32(SIZE*(i+1)))
		w.direction[i] = mgl32.Vec2{float32(math.Cos(float64(v))), float32(math.Sin(float64(v)))}
	}

	return &w
}

func (w *WaveGenerator) waveHeight(x float32, z float32, time float32) float32 {
	var height float32
	for i := 0; i < SIZE; i++ {
		frequency := float32(2.0*math.Pi) / w.wavelength[i]
		phase := w.speed[i] * frequency
		theta := w.direction[i].Dot(mgl32.Vec2{x, z})

		height += w.amplitude[i] * float32(math.Sin(float64(theta*frequency+time*phase)))
	}
	return height
}

func (w *WaveGenerator) Update(dt float64, geo *objects.Geometry, time float32) {
	for i := 0; i < len(geo.Vertices); i += objects.Stride {
		geo.Vertices[i+1] = w.waveHeight(geo.Vertices[i], geo.Vertices[i+2], float32(time))
	}
}
