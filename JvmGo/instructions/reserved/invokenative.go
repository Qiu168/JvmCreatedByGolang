package reserved

import (
	"JvmCreatedByGolang/JvmGo/instructions/base"
	"JvmCreatedByGolang/JvmGo/native"
	_ "JvmCreatedByGolang/JvmGo/native/java/io"
	_ "JvmCreatedByGolang/JvmGo/native/java/lang"
	_ "JvmCreatedByGolang/JvmGo/native/java/security"
	_ "JvmCreatedByGolang/JvmGo/native/java/util/concurrent/atomic"
	_ "JvmCreatedByGolang/JvmGo/native/sun/io"
	_ "JvmCreatedByGolang/JvmGo/native/sun/misc"
	_ "JvmCreatedByGolang/JvmGo/native/sun/reflect"
	"JvmCreatedByGolang/JvmGo/rtda"
)

// INVOKE_NATIVE Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
