/*
初始化工作,使用File包
*/

package File

var picMap = make(map[string]string)
var vidMap = make(map[string]string)
var audioMap = make(map[string]string)

func init() {
	//特征16进制都使用小写
	//pic
	picMap["ffd8ffe0"] = "jpg"
	picMap["ffd8ffe1"] = "jpg"
	picMap["ffd8ffe8"] = "jpg"
	picMap["89504e47"] = "png"
	//video
	vidMap["000000206674797069736f6d00000200"] = "mp4"
	vidMap["52494646"] = "avi"
	vidMap["4740"] = "ts"
	vidMap["2e524d46"] = "rmvb"
	vidMap["1a45dfa3934282886d6174726f736b61"] = "mkv"
	vidMap["1a45dfa39f4286810142f7810142f281"] = "mkv"
	vidMap["3026b2758e66cf11a6d900aa0062ce6c"] = "wmv"
	//audio
	audioMap["494433"] = "mp3"
	audioMap["fffb"] = "mp3"
	audioMap["0000001c6674797069736f6d00000200"] = "m4a"
	audioMap["00000020667479706d70343200000000"] = "m4a"
	audioMap["57415645666d74"] = "wav"
	audioMap["3026b2758e66cf11a6d900aa0062ce6c"] = "wma"
	audioMap["664c6143"] = "flac"
}
