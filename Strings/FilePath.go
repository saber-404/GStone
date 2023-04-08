package Strings

import (
	"regexp"
	"strings"
)

// CheckWinDriverName 判断驱动器合法性
func CheckWinDriverName(DriverName string) (Ok bool) {
	Pattern := `^[a-zA-Z]:$`
	re := regexp.MustCompile(Pattern)
	Ok = re.MatchString(DriverName)
	return
}

// CheckWinFileName 判断windows系统文件名合法性
func CheckWinFileName(FileName string) (Ok bool) {
	PatternOne := `^(PRN|AUX|CLOCK|NUL|CON|COM[1-9]|LPT[1-9])$` //是否有保留字
	PatternTwo := `^[^*\"<>:\/\\\\|?]*$`                        //是否没有非法字符
	re1 := regexp.MustCompile(PatternOne)
	re2 := regexp.MustCompile(PatternTwo)
	//SplitN将字符串s按照分隔符号sep分割成n段,如果sep在s中没有就返回s
	x := strings.SplitN(FileName, ".", 2)
	xl := len(x)
	//xl只能是1或2
	//先检测x[0]
	//对x[0]检测保留字和非法字符 检测到就为假
	t1 := re1.MatchString(x[0])
	if t1 {
		return false
	}
	t2 := re2.MatchString(x[0])
	if !t2 {
		return false
	}
	//xl == 2的情况下，再检测x[1]
	//对x[1]检测非法字符就行 检测到就为假
	if xl == 2 {
		t2 = re2.MatchString(x[1])
		if !t2 {
			return false
		}
	}
	return true
}
func checkWinFilePath(AbsFilePath, sep string) (Ok bool) {
	l := len(AbsFilePath)
	if l < 2 {
		return false
	} else {
		PathParts := strings.Split(AbsFilePath, sep)
		DriverNamePart := PathParts[0]
		if !CheckWinDriverName(DriverNamePart) {
			return false
		}
		PathParts = PathParts[1:]
		pl := len(PathParts)
		if pl == 0 {
			return true
		}
		for _, filename := range PathParts {
			if !CheckWinFileName(filename) {
				return false
			}
		}
		return true
	}
}

// CheckWinFilePath 判断windows系统绝对路径字符串合法性
func CheckWinFilePath(AbsFilePath string) (Ok bool) {
	return checkWinFilePath(AbsFilePath, `\`) || checkWinFilePath(AbsFilePath, `/`)
}
