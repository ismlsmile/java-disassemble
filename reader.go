package disassemble

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"sync"
)

var (
	buf32bitPool = sync.Pool{
		New: func() interface{} { return make([]byte, 4) },
	}

	buf16bitPool = sync.Pool{
		New: func() interface{} { return make([]byte, 2) },
	}
)

type ClassReader struct {
	reader *bufio.Reader
}

func (c *ClassReader) readUint32() uint32 {
	buffer := buf32bitPool.Get().([]byte)
	// return buffer upon completion
	defer buf32bitPool.Put(buffer)

	c.reader.Read(buffer)

	return binary.BigEndian.Uint32(buffer)
}

func (c *ClassReader) readUint16() uint16 {
	buffer := buf16bitPool.Get().([]byte)
	// return buffer upon completion
	defer buf16bitPool.Put(buffer)

	c.reader.Read(buffer)

	return binary.BigEndian.Uint16(buffer)
}

func (c *ClassReader) readUTF8String() string {
	length := c.readUint16()

	buf := make([]byte, length)

	c.reader.Read(buf)

	return string(buf)
}

func (c *ClassReader) readString() string {
	return ""
}

func (c *ClassReader) readByte() byte {
	b, _ := c.reader.ReadByte()
	return b
}

func (c *ClassReader) readConstant() *JConstant {
	t := Tag(c.readByte())

	size := t.Size()

	if t == UTF8String {
		size += int(c.readUint16())
	}

	buf := make([]byte, size)

	c.reader.Read(buf)

	return &JConstant{
		Tag: t,
		Buf: buf,
	}
}

func (c *ClassReader) ReadClass() (class *JClass, err error) {
	class = &JClass{}

	class.Magic = c.readUint32()
	class.MinorVer = c.readUint16()
	class.MajorVer = c.readUint16()
	class.ConstantPoolSize = c.readUint16()

	// The constant_pool table is indexed from 1 to constant_pool_count-1.
	class.ConstantPool = make([]*JConstant, class.ConstantPoolSize-1)

	// TODO: fix me.
	// I am reading these constants incorrectly
	for i := 0; i < int(class.ConstantPoolSize-1); i++ {
		fmt.Printf("%v\n", c.readConstant())
	}

	return
}

func NewClassReader(file string) *ClassReader {
	f, _ := ioutil.ReadFile(file)

	return &ClassReader{
		reader: bufio.NewReader(bytes.NewReader(f)),
	}
}
