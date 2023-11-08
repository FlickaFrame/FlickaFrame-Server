package code

import "github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"

const (
	NoSupportCommentType = 43000 + iota
)

var (
	NoSupportCommentTypeErr = xcode.Add(NoSupportCommentType, "No support comment type(one of parent or child)")
)
