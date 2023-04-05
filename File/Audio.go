package File

import (
	"github.com/saber-404/GStone/Strings"
	"os"
	"strings"
)

func IsAudio(AudioPath string) (Ok bool) {
	Ok, _ = isAudio(AudioPath)
	return
}

func isAudio(AudioPath string) (Ok bool, AudioType string) {
	file, err := os.Open(AudioPath)
	defer file.Close()
	if err != nil {
		return false, ""
	}
	buf := make([]byte, 16)
	n, _ := file.Read(buf)
	fileCode := Strings.BytesToHexString(buf[:n])
	for k, _ := range audioMap {
		if strings.HasPrefix(fileCode, k) {
			return true, audioMap[k]
		}
	}
	return false, ""
}

func AudioType(AudioPath string) (AudioType string) {
	_, AudioType = isAudio(AudioPath)
	return
}
