package disassemble

type Tag byte

const (
	UTF8String    Tag = 1        // 1
	Integer       Tag = iota + 3 // 3
	Float                        // 4
	Long                         // 5
	Double                       // 6
	ClassRef                     // 7
	StringRef                    // 8
	FieldRef                     // 9
	MethodRef                    // 10
	InterfaceRef                 // 11
	NameTypeDesc                 // 12
	MethodHandle  Tag = 15       // 15
	MethodType    Tag = 16       // 16
	InvokeDynamic Tag = 18       // 18
)

//JConstant represents a constant in the constant pool
type JConstant struct {
	Tag byte
}

// Read function
