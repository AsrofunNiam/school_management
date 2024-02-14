package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func InitSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + InitSeed().Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[InitSeed().Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUserId() string {
	return strconv.Itoa(int(RandomInt(1, 99))) + RandomString(3)
}

func RandomName() string {
	return RandomString(5)
}

func RandomRole() string {
	roles := []string{"teacher", "student"}
	n := len(roles)
	return roles[InitSeed().Intn(n)]
}
