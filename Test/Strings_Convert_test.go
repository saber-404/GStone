package Test

import (
	"encoding/hex"
	"fmt"
	"github.com/saber-404/GStone/Strings"
	"testing"
)

func TestBytesToHexString(t *testing.T) {
	src := []byte{0x0f, 0x31, 0x32, 0x33}
	for _, b := range src {
		fmt.Printf("%08b ", b)
	}
	println()
	dst := make([]byte, len(src)*2)
	hex.Encode(dst, src)
	s := string(dst)
	println(s)
	println(Strings.BytesToHexString(src))
}
