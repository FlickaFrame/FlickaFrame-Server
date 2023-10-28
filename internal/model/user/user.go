package user

import (
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"gorm.io/gorm"
)

const (
	SexUnknown = iota
	SexMale
	SexFemale
)

const (
	SaltByteLength        = 16
	PasswordHashAlgorithm = "argon2"
)

type User struct {
	gorm.Model
	NickName  string `gorm:"type:varchar(32)"` // 昵称
	AvatarUrl string // 头像地址
	Age       int    // 年龄
	Gender    int    // 性别
	Password  string // 密码
	Phone     string `gorm:"type:varchar(100);index:idx_phone,unique"` // 手机号
	Slogan    string // 个人简介
	TikTokID  string `gorm:"type:varchar(100);index:idx_tiktok"` // 抖音ID

	FollowingCount int
	FollowerCount  int
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) SetPassword(passwd string) (err error) {
	//TODO: 密钥加盐
	if len(passwd) == 0 {
		u.Password = ""
		return nil
	}
	u.Password = util.Md5ByString(u.Password)
	return nil
}

// ValidatePassword checks if the given password matches the one belonging to the user.
func (u *User) ValidatePassword(password string) bool {
	return u.Password == util.Md5ByString(password)
}

// IsPasswordSet checks if the password is set or left empty
func (u *User) IsPasswordSet() bool {
	return len(u.Password) != 0
}
