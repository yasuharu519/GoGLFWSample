package main
 
import (
	"runtime"
 
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)
 
func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}
 
func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
 
	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
 
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
 
	if err := gl.Init(); err != nil {
		panic(err)
	}
 
	for !window.ShouldClose() {
		w, h := window.GetFramebufferSize()
		ratio := w / h
 
		gl.Viewport(0, 0, int32(w), int32(h))
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(float64(-ratio), float64(ratio), -1, 1, 1, -1)
		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
		gl.Rotatef(float32(glfw.GetTime()*50), 0., 0., 1.)
		gl.Begin(gl.TRIANGLES)
		gl.Color3f(1., 0., 0.)
		gl.Vertex3f(-0.6, -0.4, 0.)
		gl.Color3f(0., 1., 0.)
		gl.Vertex3f(0.6, -0.4, 0.)
		gl.Color3f(0., 0., 1.)
		gl.Vertex3f(0., 0.6, 0.)
		gl.End()
 
		window.SwapBuffers()
		glfw.PollEvents()
	}
}