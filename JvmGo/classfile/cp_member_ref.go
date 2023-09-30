package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberRefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
