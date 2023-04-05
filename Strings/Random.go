package Strings

import (
	"math/rand"
	"time"
)

func randString(MetaString string, length int) string {
	bytes := []byte(MetaString)
	var result []byte
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func RandomAlphaString(length int) string {
	MetaString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randString(MetaString, length)
}

func RandNumberString(length int) string {
	MetaString := "0123456789"
	return randString(MetaString, length)
}
