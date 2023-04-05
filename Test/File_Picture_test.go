package Test

import (
	"github.com/saber-404/GStone/File"
	"testing"
)

func TestPic(t *testing.T) {
	x := `D:\test.jpg`
	println(File.IsPicture(x))
	println(File.PictureType(x))
}
