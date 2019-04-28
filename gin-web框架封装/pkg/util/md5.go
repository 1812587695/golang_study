package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum([]byte("hytx滴加密")))
}

func RandMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	r := rand.New(rand.NewSource(time.Now().Unix()))
	num := r.Intn(100000)
	t := time.Now().Unix()

	return hex.EncodeToString(m.Sum([]byte(string(num) + string(t))))
}
