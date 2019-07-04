package yami64

import "fmt"

var alphabets [64]int
var ascii [256]int

func init() {
	alphabets = [64]int{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '_',
	}

	for i := 0; i < 64; i++ {
		ascii[alphabets[i]] = i
	}

	fmt.Println(alphabets)
	fmt.Println(ascii)
	fmt.Println(ascii['-'])
}

/**
编码原始字符串
 */
func Encode(raw string) string {
	for i := 0; i < len(raw); i++ {
	}
	return ""
}

/**
解码yami64字符串
 */
func Decode(encoded string) string {
	return ""
}
