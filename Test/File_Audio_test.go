package Test

import (
	"github.com/saber-404/GStone/File"
	"testing"
)

func TestAudio(t *testing.T) {
	x := `C:\example.wma`
	println(File.IsAudio(x))
	println(File.AudioType(x))
}
