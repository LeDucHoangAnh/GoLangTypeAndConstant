package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Boolean
	var myBool bool = true //false
	fmt.Println(myBool)

	var mySecondBool bool = false
	fmt.Println(mySecondBool)

	//String
	var myString string = "hello"
	fmt.Println(myString)

	//Int
	var myInt int = 123
	fmt.Println(myInt)

	//int 8, 16, 32, 64:số lượng bit tối đa có thể đại diện cho các biến có kiểu dữ liệu dạng này
	//1. Range
	//Range Int8 (-128 -> 127)
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	//Range Int16 (-32768 -> 32767)
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	//Range Int32 (-2147483648 -> 2147483647)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	//Range Int64 (-9223372036854775808 -> 9223372036854775807)
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	//2. Bits
	fmt.Println("==========")
	fmt.Println(bits.OnesCount8(math.MaxInt8))
	fmt.Println(bits.OnesCount16(math.MaxInt16))
	fmt.Println(bits.OnesCount32(math.MaxInt32))
	fmt.Println(bits.OnesCount64(math.MaxInt64))

	//uint: Số nguyên dương
	var myUInt uint = 10
	fmt.Println(myUInt)
	//byte:tối đa 0->255
	var myByte byte = 255
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)
	fmt.Println()

	var a byte = 'E'
	fmt.Printf("%X", a)
	fmt.Println()

	//float: Số thực
	//float64
	var myFloat float64 = 10.01
	fmt.Println(myFloat)

	//complex z = a + bi:số phức
	var z complex64 = 10 + 1i
	fmt.Println(z)

	var z1 complex64 = 10 + 1i
	fmt.Println(z1)
	fmt.Println(z1 + z)

	//Rune:
	var myString1 string = "Nhẫn"

	runes := []rune(myString1)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	println()
	var myRune rune = 'A'
	fmt.Printf("%T", myRune) // bản chất rune giống int32
	println()
	//Zero values
	//Type conversions
	var zero1 int = 1
	var zero2 float64 = 2.1
	fmt.Println(zero1 + int(zero2))
	fmt.Println(zero2 + float64(zero1))
}
