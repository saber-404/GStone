package Test

import (
	"github.com/saber-404/GStone/Strings"
	"testing"
	"time"
)

func TestRandNumberString(t *testing.T) {
	for i := 0; i < 10; i++ {
		x := Strings.RandNumberString(6)
		time.Sleep(1 * time.Nanosecond)
		println(x)
	}
}

func TestRandomAlphaString(t *testing.T) {
	for i := 0; i < 10; i++ {
		x := Strings.RandomAlphaString(6)
		time.Sleep(2 * time.Microsecond)
		println(x)
	}
}
