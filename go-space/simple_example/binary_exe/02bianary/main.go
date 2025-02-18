package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const INT_SIZE int = int(unsafe.Sizeof(0))

//判断我们系统中的字节序类型
func systemEdian() {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
}

func testBigEndian() {

	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use big endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func BigEndian2(data []byte) {
	fmt.Printf("%v, %s use big endian: \n", data, string(data))

	buf := make([]byte, 4+len(data))
	fmt.Printf("buf: %v\n", buf)

	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	fmt.Printf("buf: %v\n", buf)

	copy(buf[4:], data)

	fmt.Printf("buf: %v\n", buf)

	fmt.Println("data to bytes:", data)


	fmt.Println("\n\n")
}

func testLittleEndian() {

	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use little endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
	data := "helloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworld"
	BigEndian2([]byte(data))

	// systemEdian()
	// testBigEndian()
	// testLittleEndian()
}
