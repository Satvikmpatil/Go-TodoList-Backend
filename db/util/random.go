package util

import (
	"math/rand"
	"strings"
	"time"
)


const (
	abc = "abcdefghijklmnopqrstuvwxyz"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min,max int64) int64{
	return min+rand.Int63n(max-min+1)
}

func Randomstring(n int)string{
	var sb strings.Builder
	k := len(abc)
	for i := 0;i<n;i++{
		c := abc[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUser() string{
	return Randomstring(6)
}

func RandomTask() string{
	return Randomstring(10)
}

func Randomstatus() bool{
	if RandomInt(2,10) % 2 == 0 {
		return true
	} else{
		return false
	}
}