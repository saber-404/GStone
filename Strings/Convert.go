package Strings

// BytesToHexString 获取16进制
func BytesToHexString(src []byte) string {
	l := len(src) * 2 //半个字节映射为一个字符
	if src == nil || l <= 0 {
		return ""
	}
	dst := make([]byte, l)
	hextable := "0123456789abcdef"
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]     //得到高半字节,右移操作
		dst[j+1] = hextable[v&0x0f] //得到底半字节
		j += 2
	}
	return string(dst)
	/*
		//以01字符串打印
		for _, b := range src {
			fmt.Printf("%08b ", b)
		}
	*/
}
