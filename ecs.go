package main

import (
	"github.com/mzdravkov/tedronai/components"
)

type Components struct {
	Masks components.ComponentMaskChunkedVec

	Spaceable  components.SpaceComponentChunkedVec
	Renderable components.RenderComponentChunkedVec
	Hoverable  components.HoverComponentChunkedVec
}
