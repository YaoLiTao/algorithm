package yami64

var alphabets [64]byte
var ascii [256]byte

func init() {
	alphabets = [64]byte{
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
		ascii[alphabets[i]] = byte(i)
	}
}

/**
编码原始字符串
 */
func Encode(raw string) string {
	rawByte := []byte(raw)
	var encodedByte []byte
	for i := 0; i < len(raw); i = i + 3 {
		encodedByte = append(encodedByte, alphabets[rawByte[i]&0x3F])
		encodedByte = append(encodedByte, alphabets[((rawByte[i]&0xC0)>>6)|((rawByte[i+1]&0x0F)<<2)])
		encodedByte = append(encodedByte, alphabets[((rawByte[i+1]&0xF0)>>4)|((rawByte[i+2]&0x03)<<4)])
		encodedByte = append(encodedByte, alphabets[(rawByte[i+2]&0xFC)>>2])
	}
	return string(encodedByte)
}

/**
解码yami64字符串
 */
func Decode(encoded string) string {
	encodedByte := []byte(encoded)
	var rawByte []byte
	for i := 0; i < len(encodedByte); i = i + 4 {
		rawByte = append(rawByte, ascii[encodedByte[i]]|(ascii[encodedByte[i+1]]&0x03)<<6)
		rawByte = append(rawByte, (ascii[encodedByte[i+1]]&0x3C)>>2|(ascii[encodedByte[i+2]]&0x0F)<<4)
		rawByte = append(rawByte, (ascii[encodedByte[i+2]]&0x30)>>4|(ascii[encodedByte[i+3]]&0x3F)<<2)
	}
	return string(rawByte)
}
