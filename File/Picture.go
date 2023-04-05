package File

import (
	"github.com/saber-404/GStone/Strings"
	"os"
)

func isPicture(PicPath string) (Ok bool, PicType string) {
	file, err := os.Open(PicPath)
	defer file.Close()
	if err != nil {
		return false, ""
	}
	buf := make([]byte, 4)
	n, _ := file.Read(buf)
	fileCode := Strings.BytesToHexString(buf[:n])
	for k, _ := range picMap {
		//if strings.HasPrefix(fileCode, k) {
		//	return true,picMap[k]
		//}
		if fileCode == k {
			return true, picMap[k]
		}
	}
	return false, ""

}
func IsPicture(PicPath string) (Ok bool) {
	Ok, _ = isPicture(PicPath)
	return
}
func PictureType(PicPath string) (PicType string) {
	_, PicType = isPicture(PicPath)
	return
}
