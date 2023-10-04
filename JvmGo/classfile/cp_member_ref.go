package classfile

/*
ConstantMemberRefInfo
ConstantFieldRefInfo

	CONSTANT_Fieldref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
	}

ConstantMethodRefInfo

	CONSTANT_Methodref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
	}

ConstantInterfaceMethodRefInfo

	CONSTANT_InterfaceMethodref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
	}
*/
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
