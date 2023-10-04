package classfile

type MarkerAttribute struct {
}

/*
DeprecatedAttribute

	Deprecated_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
SyntheticAttribute

	Synthetic_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

func (self MarkerAttribute) readInfo(reader *ClassReader) {
	//只起标志作用无数据
}
