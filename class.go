package disassemble

const MagicNumber uint32 = 0xCAFEBABE

const (
	// Access Flags
	AccPublic     uint16 = 0x0001
	AccFinal      uint16 = 0x0010
	AccSuper      uint16 = 0x0020
	AccInterface  uint16 = 0x0200
	AccAbstract   uint16 = 0x0400
	AccSynthetic  uint16 = 0x1000
	AccAnnotation uint16 = 0x2000
	AccEnum       uint16 = 0x4000
)

//JClass represents a compiled Java class
//The class file is read in sequential byte order
type JClass struct {
	// first 4 bytes
	Magic uint32
	// 5-6
	MinorVer uint16
	// 6-7
	MajorVer uint16
	// todo: constant pool
	ConstantPoolSize uint16
	ConstantPool     []*JConstant
	// 10+cpsize
	AccessFlags    uint16
	ClassName      string
	SuperClassName string
	// todo: interface
	// todo: fields
	// todo: method
	// todo: attributes
}

func (j *JClass) HasAccessFlag(bit uint16) bool {
	return j.AccessFlags&bit != 0
}
