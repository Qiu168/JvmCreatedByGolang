package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	SourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.cp.getUtf8(self.SourceFileIndex)
}
