package scene

import (
	"git.maze.io/go/math32"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera_Movment int

const (
	FORWARD  Camera_Movment = iota
	BACKWARD Camera_Movment = iota
	LEFT     Camera_Movment = iota
	RIGHT    Camera_Movment = iota
	UP       Camera_Movment = iota
	DOWN     Camera_Movment = iota
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

	camera.MovementSpeed = 10.0
	camera.MouseSensitivity = 0.25

	updateCameraVectors(camera)

	return camera
}

func GetViewMatrix(camera *Camera) mgl32.Mat4 {
	return mgl32.LookAtV(camera.Position, camera.Position.Add(camera.Front), camera.Up)
}

func ProcessKeyboard(camera *Camera, direction Camera_Movment, deltaTime float32) {
	velocity := camera.MovementSpeed * deltaTime
	if direction == FORWARD {
		camera.Position = camera.Position.Add(camera.Front.Mul(velocity))
	}

	if direction == BACKWARD {
		camera.Position = camera.Position.Add(camera.Front.Mul(velocity).Mul(-1))
	}

	if direction == LEFT {
		camera.Position = camera.Position.Add(camera.Right.Mul(velocity).Mul(-1))
	}

	if direction == RIGHT {
		camera.Position = camera.Position.Add(camera.Right.Mul(velocity))
	}

	if direction == UP {
		camera.Position = camera.Position.Add(mgl32.Vec3{0, 1, 0}.Mul(velocity))
	}

	if direction == DOWN {
		camera.Position = camera.Position.Add(mgl32.Vec3{0, 1, 0}.Mul(velocity).Mul(-1))
	}
}

func ProcessMouseMovement(camera *Camera, xpos float32, ypos float32) {

	xpos = -xpos * camera.MouseSensitivity
	ypos = -ypos * camera.MouseSensitivity

	camera.Yaw = (-180 / 80.0) * (xpos + 40)
	camera.Pitch = (180.0/60.0)*(ypos+30) + -90.0

	updateCameraVectors(camera)
}

func updateCameraVectors(camera *Camera) {
	fx := math32.Cos(mgl32.DegToRad(camera.Yaw)) * math32.Cos(mgl32.DegToRad(camera.Pitch))
	fy := math32.Sin(mgl32.DegToRad(camera.Pitch))
	fz := math32.Sin(mgl32.DegToRad(camera.Yaw)) * math32.Cos(mgl32.DegToRad(camera.Pitch))

	camera.Front = mgl32.Vec3{fx, fy, fz}.Normalize()
	camera.Right = camera.Front.Cross(camera.WorldUp).Normalize()
	camera.Up = camera.Right.Cross(camera.Front).Normalize()

}
