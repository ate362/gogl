package scene

import (
	"git.maze.io/go/math32"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3
	WorldUp  mgl32.Vec3

	Yaw   float32
	Pitch float32

	MovementSpeed    float32
	MouseSensitivity float32
	Zoom             float32
}

func InitCamera(posX float32, posY float32, posZ float32, upX float32, upY float32, upZ float32, yaw float32, pitch float32) *Camera {
	camera := new(Camera)
	camera.Position = mgl32.Vec3{posX, posY, posZ}
	camera.WorldUp = mgl32.Vec3{upX, upY, upZ}
	camera.Yaw = yaw
	camera.Pitch = pitch
	camera.updateCameraVectors()

	return camera
}

func GetViewMatrix(camera *Camera) mgl32.Mat4 {
	return mgl32.LookAtV(camera.Position, camera.Position.Add(camera.Front), camera.Up)
}

func (camera *Camera) updateCameraVectors() {
	fx := math32.Cos(mgl32.DegToRad(camera.Yaw)) * math32.Cos(mgl32.DegToRad(camera.Pitch))
	fy := math32.Sin(mgl32.DegToRad(camera.Pitch))
	fz := math32.Sin(mgl32.DegToRad(camera.Yaw)) * math32.Cos(mgl32.DegToRad(camera.Pitch))

	camera.Front = mgl32.Vec3{fx, fy, fz}.Normalize()
	camera.Right = camera.Front.Cross(camera.WorldUp).Normalize()
	camera.Up = camera.Right.Cross(camera.Front).Normalize()

}
