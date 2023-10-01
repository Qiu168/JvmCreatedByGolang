package rtda

import "JvmCreatedByGolang/JvmGo/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
