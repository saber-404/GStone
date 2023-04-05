package File

import (
	"github.com/saber-404/GStone/Strings"
	"os"
	"strings"
)

func isVideo(VideoPath string) (Ok bool, VideoType string) {
	file, err := os.Open(VideoPath)
	defer file.Close()
	if err != nil {
		return false, ""
	}
	buf := make([]byte, 16)
	n, _ := file.Read(buf)
	fileCode := Strings.BytesToHexString(buf[:n])
	for k, _ := range vidMap {
		if strings.HasPrefix(fileCode, k) {
			return true, vidMap[k]
		}
	}
	return false, ""
}
func IsVideo(VideoPath string) (Ok bool) {
	Ok, _ = isVideo(VideoPath)
	return
}
func VideoType(VideoPath string) (VideoType string) {
	_, VideoType = isVideo(VideoPath)
	return
}
