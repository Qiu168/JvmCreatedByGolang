package heap

import "JvmCreatedByGolang/JvmGo/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		//不考虑java8
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
