package systems

import (
	. "github.com/mzdravkov/tedronai/components"
)

type System interface {
	Run(*ComponentMaskChunkedVec, map[ComponentMask]interface{})
}
