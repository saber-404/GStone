package Test

import (
	"github.com/saber-404/GStone/File"
	"testing"
)

func TestVid(t *testing.T) {
	x := `D:\test.ts`
	println(File.IsVideo(x))
	println(File.VideoType(x))
}
