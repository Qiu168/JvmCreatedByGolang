package classfile

type MarkerAttribute struct {
}
type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}

func (self MarkerAttribute) readInfo(reader *ClassReader) {
	//只起标志作用无数据
}
