package rand

import (
	"math/rand"
)

const (
	Number      = "1234567890"
	LowerLetter = "qwertyuiopasdfghjklzxcvbnm"
	UpperLetter = "QWERTYUIOPASDFGHJKLZXCVBNM"
	Symbol      = "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?"
)

func NonceStr(length int, bytes ...byte) (result string) {
	if length <= 0 {
		return
	}
	bl := len(bytes)
	if bl == 0 {
		bytes = []byte(Number)
		bl = len(bytes)
	}
	var bt byte
	for i := 0; i < length; i++ {
		bt = bytes[rand.Intn(bl)]
		result += string(bt)
	}
	return
}
