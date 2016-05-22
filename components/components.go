package components

import (
	"image/color"
	"sync"

	"engo.io/engo"
	"engo.io/engo/common"
	"engo.io/gl"
)

// type ChunkedVec interface {
// 	Add(interface{}) (uint, uint)
// 	PutAt(interface{}, uint, uint)
// 	DeleteAt(uint, uint)
// 	Get(uint, uint) interface{}
// 	Shrink()
// 	Len() int
// 	Cap() int
// 	Iter() interface{}
// 	Contains(interface{}) bool
// 	ContainsAll(interface{}) bool
// 	Equal(interface{}) bool
// 	Clone() ChunkedVec
// 	Clear()
// 	String() string
// }

// +gen ChunkedVec
type ComponentMask uint

const (
	Nil       ComponentMask = 0
	Spaceable               = 1 << iota
	Renderable
	Hoverable
)

// ====================================
//             Components
// ====================================

// +gen ChunkedVec
type SpaceComponent struct {
	sync.Mutex
	common.SpaceComponent
}

// +gen ChunkedVec
type RenderComponent struct {
	sync.Mutex
	// Hidden is used to prevent drawing by OpenGL
	Hidden bool
	// Scale is the scale at which to render, in the X and Y axis. Not defining Scale, will default to engo.Point{1, 1}
	Scale engo.Point
	// Color defines how much of the color-components of the texture get used
	Color color.Color
	// Drawable refers to the Texture that should be drawn
	Drawable common.Drawable
	// Repeat defines how to repeat the Texture if the viewport of the texture is larger than the texture itself
	Repeat common.TextureRepeating

	Shader common.Shader
	ZIndex float32

	Buffer *gl.Buffer
	// bufferContent []float32
}

func CastRenderComponentToEngo(renderComponent *RenderComponent) *common.RenderComponent {
	component := &common.RenderComponent{
		Hidden:   renderComponent.Hidden,
		Scale:    renderComponent.Scale,
		Color:    renderComponent.Color,
		Drawable: renderComponent.Drawable,
		Repeat:   renderComponent.Repeat,
	}

	component.SetShader(renderComponent.Shader)
	component.SetZIndex(renderComponent.ZIndex)

	return component
}

// +gen ChunkedVec
type HoverComponent struct {
	sync.Mutex
}
