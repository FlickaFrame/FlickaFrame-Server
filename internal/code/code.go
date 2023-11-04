package code

import (
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"
	"gorm.io/gorm"
)

var (
	ErrNotFound = gorm.ErrRecordNotFound
)

const ( // 40000 - 40099
	FollowUserIdEmptyCode = 40000 + iota
	FollowedUserIdEmptyCode
	CannotFollowSelfCode
	CannotUnfollowSelfCode
	HadFollowedCode
	HadNotFollowedCode
)

var ( // 关注相关错误
	FollowUserIdEmpty   = xcode.Add(FollowUserIdEmptyCode, "关注用户id为空")
	FollowedUserIdEmpty = xcode.Add(FollowedUserIdEmptyCode, "被关注用户id为空")
	CannotFollowSelf    = xcode.Add(CannotFollowSelfCode, "Can not follow yourself")
	CannotUnfollowSelf  = xcode.Add(CannotUnfollowSelfCode, "Can not unfollow yourself")
	HadFollowed         = xcode.Add(HadFollowedCode, "Had followed")
	HadNotFollowed      = xcode.Add(HadNotFollowedCode, "you have not followed this user")
)

var (
	UserIdEmpty                 = xcode.Add(40011, "用户id为空")
	ErrUserAlreadyRegisterError = xcode.Add(40012, "user has been registered")
	ErrGenerateTokenError       = xcode.Add(40013, "generate token error")
	ErrUserNoExistsError        = xcode.Add(40021, "user not exists")
	ErrUsernamePwdError         = xcode.Add(40021, "username or password error")
	ErrCommentNoExistsError     = xcode.Add(40031, "comment not exists")
	ErrCommentNoPermissionError = xcode.Add(40032, "no permission to delete comment")
)
