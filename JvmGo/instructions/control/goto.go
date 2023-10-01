package control

import (
	"JvmCreatedByGolang/JvmGo/instructions/base"
	"JvmCreatedByGolang/JvmGo/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
