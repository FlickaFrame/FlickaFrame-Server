package code

import "github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"

const (
	DuplicateFavoriteCode = 44000 + iota // 不允许重复点赞
)

var (
	DuplicateFavoriteErr = xcode.Add(DuplicateFavoriteCode, "不允许重复点赞")
)
