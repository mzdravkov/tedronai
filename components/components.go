package components

import "engo.io/engo"

// +gen ChunkedVec
type ComponentMask uint

const (
	Nil       ComponentMask = 0
	Spaceable ComponentMask = 1 << iota
	Renderable
	Hoverable
)

// ====================================
//             Components
// ====================================

// +gen ChunkedVec
type SpaceComponent engo.SpaceComponent

// +gen ChunkedVec
type RenderComponent struct {
	Texture string
	Scale   engo.Point
}

// +gen ChunkedVec
type HoverComponent struct{}
