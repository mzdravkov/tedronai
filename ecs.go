package main

import (
	_ "container/list"
	"sync"

	. "github.com/mzdravkov/tedronai/components"
	. "github.com/mzdravkov/tedronai/systems"
)

const (
	CHUNK_SIZE uint = 1024
)

var (
	ComponentMasks     = NewComponentMaskChunkedVec(CHUNK_SIZE)
	ComponentMasksLock = sync.Mutex{}

	Components = map[ComponentMask]interface{}{
		Spaceable:  NewSpaceComponentChunkedVec(CHUNK_SIZE),
		Renderable: NewRenderComponentChunkedVec(CHUNK_SIZE),
		Hoverable:  NewHoverComponentChunkedVec(CHUNK_SIZE),
	}

	Systems = []System{
		NewRenderSystem(),
	}
)

func init() {
	// RunSystems(Systems)
}

func RunSystems(systems []System) {
	for _, system := range systems {
		system.Run(ComponentMasks, Components)
	}
}

func AddEntity(components map[ComponentMask]interface{}) {
	mask := Nil
	for m, _ := range components {
		mask &= m
	}

	// Add the entity to ComponentMasks
	// with the locking, we guarantee that different goroutines won't try to put
	// on the same place different elements at the same time
	ComponentMasksLock.Lock()
	capBefore := ComponentMasks.Cap()
	i, k := ComponentMasks.Add(mask)
	capAfter := ComponentMasks.Cap()
	shouldGrow := false
	if capAfter > capBefore {
		shouldGrow = true
	}
	ComponentMasksLock.Unlock()

	// we then put all the components provided as arguments to the position
	// on which we have put the mask in ComponentsMask
	for m, v := range components {
		switch m {
		case Spaceable:
			spaceComponents := Components[m].(*SpaceComponentChunkedVec)
			if shouldGrow {
				spaceComponents.Grow(1)
			}
			spaceComponents.PutAt(v.(SpaceComponent), i, k)
		case Renderable:
			renderComponents := Components[m].(*RenderComponentChunkedVec)
			if shouldGrow {
				renderComponents.Grow(1)
			}
			renderComponents.PutAt(v.(RenderComponent), i, k)
		case Hoverable:
			hoverComponents := Components[m].(*HoverComponentChunkedVec)
			if shouldGrow {
				hoverComponents.Grow(1)
			}
			hoverComponents.PutAt(v.(HoverComponent), i, k)
		}
	}
}

// // Merges the masks of different components (basically reducing them with &)
// // If called with no arguments, it'll return Nil mask
// func mergeMasks(masks ...ComponentMask) ComponentMask {
// 	var mask uint = 0
// 	for _, c := range masks {
// 		mask &= uint(c)
// 	}

// 	return components.ComponentMask(mask)
// }
