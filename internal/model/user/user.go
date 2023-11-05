package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"time"
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
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

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

type UserModel struct {
	db *orm.DB
}

func NewUserModel(db *orm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) Insert(ctx context.Context, data *User) error {
	data.ID = snowflake.UserIDNode.Generate().Int64()
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *UserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *UserModel) MustFindOne(ctx context.Context, id int64) *User {
	user, _ := m.FindOne(ctx, id)
	return user
}

func (m *UserModel) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("phone = ?", phone).First(&result).Error
	return &result, err
}

func (m *UserModel) Update(ctx context.Context, user *User) error {
	return m.db.WithContext(ctx).Model(user).Updates(user).Error
}

type ListOptions struct {
	orm.ListOptions
	UserIds []uint
}

func (m *UserModel) List(ctx context.Context, listOptions ListOptions) ([]*User, error) {
	var result []*User
	sess := m.db.WithContext(ctx)
	sess = orm.SetSessionPagination(sess, &listOptions) // set pagination
	err := sess.Find(&result).Error
	return result, err
}
