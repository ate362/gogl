package scene

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Handler struct {
	camera *Camera
	delta  *float64
}

func Init(c *Camera, d *float64) Handler {
	handler := &Handler{}
	handler.camera = c
	handler.delta = d
	return *handler
}

func (handler *Handler) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {

	case key == glfw.KeyEscape && action == glfw.Press:
		w.SetShouldClose(true)

	case key == glfw.KeyW && (action == glfw.Repeat || action == glfw.Press):
		ProcessKeyboard(handler.camera, 0, float32(*handler.delta))

	case key == glfw.KeyS && (action == glfw.Repeat || action == glfw.Press):
		ProcessKeyboard(handler.camera, 1, float32(*handler.delta))

	case key == glfw.KeyA && (action == glfw.Repeat || action == glfw.Press):
		ProcessKeyboard(handler.camera, 2, float32(*handler.delta))

	case key == glfw.KeyD && (action == glfw.Repeat || action == glfw.Press):
		ProcessKeyboard(handler.camera, 3, float32(*handler.delta))

	case key == glfw.KeyE && (action == glfw.Repeat || action == glfw.Press):
		ProcessKeyboard(handler.camera, 4, float32(*handler.delta))

	case key == glfw.KeyQ && (action == glfw.Repeat || action == glfw.Press):
		ProcessKeyboard(handler.camera, 5, float32(*handler.delta))
	}

}

func (handler *Handler) MouseCallback(w *glfw.Window, xpos float64, ypos float64) {

	xpos = xpos - float64(800)/2
	ypos = ypos - float64(600)/2
	ProcessMouseMovement(handler.camera, float32(xpos), float32(ypos))
}
