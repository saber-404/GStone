package Test

import (
	"fmt"
	"github.com/saber-404/GStone/Strings"
	"testing"
)

func TestCheckWinFilePath(t *testing.T) {
	TestPath := []string{
		"C:\\pictures\\holiday\\photo.jpg",
		"D:",
		`C:\pictures\holiday\photo.jpg`,
		`C:\CON`,
		`C:\CON.txt`,
		`C:\CONF.txt*`,
		`C:\     \CONF.txt`,
		`C:/123/`,
	}
	for _, s := range TestPath {
		fmt.Println(s, Strings.CheckWinFilePath(s))
	}
}

func TestCheckWinFileName(t *testing.T) {
	TESTFILENAME := []string{
		`CON`,        //f
		`CON.123`,    //f
		`.CON...123`, //t
		`   `,        //f
		`123   444`,  //t
		`123*`,       //f
	}
	for _, s := range TESTFILENAME {
		fmt.Println(s, Strings.CheckWinFileName(s))
	}
	//SingleTest := `CON.123`
	//fmt.Println(SingleTest, CheckWinFileName(SingleTest))

}
