package rtda

import (
	"JvmCreatedByGolang/JvmGo/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// SetFloat GetFloat
// Java 虚拟机内部通常将浮点数表示为有符号整数，
// 但在解释浮点数位表示时，需要将其强制转换为无符号整数以确保正确解释符号位。
// 这是因为 IEEE 754 格式中的符号位在位表示中的位置不同于有符号整数。
// 这种处理方式有助于正确处理浮点数的位表示，以便在 Java 字节码中进行操作。
// /*
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// SetLong long consumes two slots
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// SetDouble double consumes two slots
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}
