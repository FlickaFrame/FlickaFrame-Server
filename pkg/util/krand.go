package util

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = iota // 纯数字
	KC_RAND_KIND_LOWER        // 小写字母
	KC_RAND_KIND_UPPER        // 大写字母
	KC_RAND_KIND_ALL          // 数字、大小写字母
)

// KRand 随机字符串
func KRand(size int, kind int) string {
	iKind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random iKind
			iKind = rand.Intn(3)
		}
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
