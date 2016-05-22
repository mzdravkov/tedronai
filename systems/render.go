package systems

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	. "github.com/mzdravkov/tedronai/components"
)

func init() {
	initShaders()
}

type RenderSystem struct {
	// common system fields
	Mask ComponentMask
	Stop chan struct{}

	// TODO: maybe remove this "optimization" as we don't sort the entities by shader
	// render specific
	currentShader common.Shader
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{Mask: Spaceable & Renderable, Stop: make(chan struct{})}
}

func (rs *RenderSystem) Run(componentMasks *ComponentMaskChunkedVec, components map[ComponentMask]interface{}) {
	go func() {
		for {
			select {
			case <-rs.Stop:
				return
			default:
				// for each node from the list
				for i, node := 0, componentMasks.List.Front(); node != nil; i, node = i+1, node.Next() {
					// for each entity in the array-node
					for k := 0; k < len(node.Value.([]ComponentMask)); k++ {
						// if the entity has all the components for the system
						if rs.Mask == node.Value.([]ComponentMask)[k] {
							spaceable := components[Spaceable].(*SpaceComponentChunkedVec).Get(uint(i), uint(k))

							// spaceable := spaceables.Get(uint(i), uint(k))
							renderables := components[Renderable].(*RenderComponentChunkedVec)
							renderable := renderables.Get(uint(i), uint(k))

							Render(&spaceable, &renderable, rs)
						}
					}
				}
				if rs.currentShader != nil {
					rs.currentShader.Post()
					rs.currentShader = nil
				}
			}
		}
	}()
}

// This a (little bit modified) copy of the engo's common.RenderSystem.Update method
// We use this, because the normal rendering system of engo is tightly entangled with their ecs
func Render(spaceComponent *SpaceComponent, renderComponent *RenderComponent, rs *RenderSystem) {
	fmt.Println("yo, doing some rendering here")
	if engo.Headless() {
		return
	}

	engo.Gl.Clear(engo.Gl.COLOR_BUFFER_BIT)

	if renderComponent.Hidden {
		return // skip this entity and continue with others
	}

	// Retrieve a shader, may be the default one -- then use it if we aren't already using it
	shader := renderComponent.Shader
	if shader == nil {
		shader = common.DefaultShader
	}

	// Change Shader if we have to
	if shader != rs.currentShader {
		if rs.currentShader != nil {
			rs.currentShader.Post()
		}
		shader.Pre()
		rs.currentShader = shader
	}

	// Setting default scale to 1
	if renderComponent.Scale.X == 0 && renderComponent.Scale.Y == 0 {
		renderComponent.Scale = engo.Point{1, 1}
	}

	// Setting default to white
	if renderComponent.Color == nil {
		renderComponent.Color = color.White
	}

	rs.currentShader.Draw(CastRenderComponentToEngo(renderComponent), &spaceComponent.SpaceComponent)
}

func initShaders() error {
	shaders := []common.Shader{
		common.DefaultShader,
		common.HUDShader,
		common.LegacyShader,
		common.LegacyHUDShader,
	}
	var err error

	for _, shader := range shaders {
		err = shader.Setup(&ecs.World{})
		if err != nil {
			return err
		}
	}

	return nil
}
