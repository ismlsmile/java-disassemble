package disassemble

const MagicNumber = 0xCAFEBABE

//JClass represents a compiled Java class
//The class file is read in sequential byte order
type JClass struct {
	// first 4 bytes
	Magic [4]byte
	// 5-6
	MinorVer uint16
	// 6-7
	MajorVer uint16
	// todo: constant pool
	ConstantPoolSize uint16
	ConstantPool     []JConstant
	// 10+cpsize
	AccessFlags    int
	ClassName      string
	SuperClassName string
	// todo: interface
	// todo: fields
	// todo: method
	// todo: attributes
}
