package main

import (
	"fmt"

	disassemble "github.com/rbrick/java-disassemble"
)

func main() {
	reader := disassemble.NewClassReader("TestClass.class")

	class, err := reader.ReadClass()

	if err != nil {
		panic(err)
	}

	fmt.Println("Is Java?:", class.Magic == disassemble.MagicNumber)
	fmt.Printf("major version: %d, minor version: %d\n", class.MajorVer, class.MinorVer)
	fmt.Println("Is at least Java8?:", class.MajorVer >= disassemble.SE8)
	fmt.Println("Class Pool Size:", class.ConstantPoolSize)
	fmt.Println("Access Flags:", class.AccessFlags)
	acc := disassemble.AccPublic | disassemble.AccFinal
	fmt.Println("Access Flag Test:", acc&disassemble.AccPublic != 0)
}
