package code

import (
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"
	"gorm.io/gorm"
)

var (
	ErrNotFound         = gorm.ErrRecordNotFound
	FollowUserIdEmpty   = xcode.Add(40001, "关注用户id为空")
	FollowedUserIdEmpty = xcode.Add(40002, "被关注用户id为空")
	CannotFollowSelf    = xcode.Add(40003, "不能关注自己")
	UserIdEmpty         = xcode.Add(40004, "用户id为空")

	ErrUserAlreadyRegisterError = xcode.Add(40005, "user has been registered")
	ErrGenerateTokenError       = xcode.Add(40006, "generate token error")

	ErrUserNoExistsError = xcode.Add(40007, "user not exists")
	ErrUsernamePwdError  = xcode.Add(40008, "username or password error")
	ErrUserFollowSelf    = xcode.Add(40009, "can not follow yourself")

	ErrCommentNoExistsError     = xcode.Add(40010, "comment not exists")
	ErrCommentNoPermissionError = xcode.Add(40011, "no permission to delete comment")
)
