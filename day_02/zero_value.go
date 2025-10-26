// package day02
package main

import "fmt"

func main() {
	var y int     // 0
	var z bool    // false
	var aa string // ""
	var bb *int   // nil
	fmt.Println("Y=", y)
	fmt.Println("z=", z)
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)
	// fmt.Println("bb=", bx) //enabling this line will through undefined error

	// Constants
	const Pi = 3.14
	// Pi = 400 // enbling this line will through error: cannot assign to Pi
	fmt.Println("Pi=", Pi)

	//Intergers
	fmt.Println("Signed integers:")
	var a int8 = 127     // ranges from -128 to 127
	var b int16 = -32768 // ranges from -32768 to 32767
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("Unsigned integers:")
	var c uint8 = 255    // ranges from 0 to 255
	var d uint16 = 65535 // ranges from 0 to 65535
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println("Machine Dependent Types:")
	var e int = 123456789 // size depends on the architecture
	var f uint = 123456789
	var g uintptr = 0xdeadbeef
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)
}
