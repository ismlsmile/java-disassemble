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

// two types of constants (longs and doubles)
// take up two consecutive slots in the table.
var tagSizeMap = map[Tag]int{
	UTF8String:    2,
	Integer:       4,
	Float:         4,
	Long:          8,
	Double:        8,
	ClassRef:      2,
	StringRef:     2,
	FieldRef:      4,
	MethodRef:     4,
	InterfaceRef:  4,
	NameTypeDesc:  4,
	MethodHandle:  3,
	MethodType:    2,
	InvokeDynamic: 4,
}

func (t Tag) Size() int {
	return tagSizeMap[t]
}

//JConstant represents a constant in the constant pool
type JConstant struct {
	Tag Tag
	Buf []byte
}

// Read function
