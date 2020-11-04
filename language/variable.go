package main

import "fmt"

func main() {
	// --------------
	// 内置类型
	// --------------

	// Type provides integrity and readability.
	// 类型提供完整性和可读性
	// - What is the amount of memory that we allocate?
	// - 我们分配的内存数量是多少？
	// - What does that memory represent?
	// -内存代表了什么？

	// Type can be specific such as int32 or int64.
	// 类型可以是特定的，比如int32或者int64
	// For example,
	// 打个比方，
	// - uint8 contains a base 10  number using one byte of memory
	// - uint8 使用1个字节的内存来容纳10个数字
	// - int32 contains a base 10 number using 4 bytes of memory.
	// - int32 使用4分字节的内存来容纳10个数字

	// When we declare a type without being very specific, such as uint or int, it gets mapped
	// 当我们声明一个不是很具体的类型（例如uint或int）时，
	// based on the architecture we are building the code against.
	// 它会根据我们要构建代码的体系结构（操作系统）进行映射。
	// On a 64-bit OS, int will map to int64. Similarly, on a 32 bit OS, it becomes int32.
	// 在64位操作系统上，int将映射到int64。同样，在32位OS上，它变为int32。

	// The word size is the number of bytes in a word, which matches our address size.
	// 字长是一个字中的字节数，与我们的地址大小匹配。
	// For example, in 64-bit architecture, the word size is 64 bit (8 bytes), address size is 64
	// 例如，在64位体系结构中，字长为64位（8字节），地址长为64位，
	// bit then our integer should be 64 bit.
	// 则我们的整数应为64位。

	// ------------------
	// 零值概念
	// ------------------

	// Every single value we create must be initialized. If we don't specify it, it will be set to
	// 我们创建的每一个单值都必须被初始化。如果我们不指定它，它将被设为0值。
	// the zero value. The entire allocation of memory, we reset that bit to 0.
	// 在整个内存分配中，我们将该位重置为0.（不占用内存空间吗？）
	// - Boolean false
	// - Integer 0
	// - Floating Point 0
	// - Complex 0i
	// - String "" (empty string)
	// - Pointer nil

	// Strings are a series of uint8 types.
	// 字符串是一系列unit8类型。
	// A string is a two word data structure: first word represents a pointer to a backing array, the
	// 字符串是一个由两个变量组成的数据结构：第一个变量表示指向返回数组的指针，
	// second word represents its length.
	// 第二个变量表示其长度。
	// If it is a zero value then the first word is nil, the second word is 0.
	// 如果它是零值，则第一个单词为nil，第二个单词为0。

	// ----------------------
	// 声明·初始化
	// ----------------------

	// var is the only guarantee to initialize a zero value for a type.
	// var是为类型初始化零值的唯一保证。
	var a int

	fmt.Printf("var a int \t %T [%v]\n", a, a)

}
