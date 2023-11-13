package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sort"
)

var (
	cacheUserIdPrefix    = "user:"       // 用户缓存主键前缀
	cacheUserPhonePrefix = "user:phone:" // 用户缓存电话前缀
	UserFavoritePrefix   = "user:favorite:"
)

type User struct {
	orm.Model

	NickName      string `gorm:"type:varchar(32)"` // 昵称
	AvatarUrl     string // 头像地址
	Age           int32  // 年龄
	Gender        int32  // 性别
	Password      string // 密码
	Phone         string `gorm:"type:varchar(100);index:idx_phone,unique"` // 手机号
	BackgroundUrl string // 背景图
	Slogan        string // 个人简介
	TikTokID      string `gorm:"type:varchar(100);index:idx_tiktok"` // 抖音ID

	FollowingCount int64
	FollowerCount  int64
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
	db         *orm.DB
	CacheRedis *redis.Redis
}

func NewUserModel(db *orm.DB, CacheRedis *redis.Redis) *UserModel {
	return &UserModel{
		db:         db,
		CacheRedis: CacheRedis,
	}
}
func (m *UserModel) FindByIDs(ctx context.Context, ids []int64) ([]*User, error) {
	var result []*User
	err := m.db.WithContext(ctx).Where("id in (?)", ids).Find(&result).Error
	val2idx := make(map[int64]int)
	for i, v := range ids {
		val2idx[v] = i
	}
	sort.Slice(result, func(i, j int) bool {
		return val2idx[result[i].ID] < val2idx[result[j].ID]
	})
	return result, err
}
func (m *UserModel) Insert(ctx context.Context, data *User) error {
	data.Model = orm.NewModel()
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *UserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var user *User
	key := fmt.Sprintf("%s%d", cacheUserIdPrefix, id)
	content, err := m.CacheRedis.GetCtx(ctx, key)
	if len(content) != 0 {
		if content == "*" {
			return nil, nil
		}
		// Warning: 可能雪花ID精度丢失
		user = &User{}
		json.Unmarshal([]byte(content), &user)
		return user, err
	} else { // 缓存不存在
		user, err = m.FindOneByDB(ctx, id)
		if user == nil { // 缓存穿透保护
			err = m.CacheRedis.SetexCtx(ctx, key, "*", 60)
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		raw, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}
		err = m.CacheRedis.SetexCtx(ctx, key, string(raw), 60)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}
func (m *UserModel) FindOneByDB(ctx context.Context, id int64) (*User, error) {
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

func (m *UserModel) Update(ctx context.Context, user *User, doerId int64) error {
	err := m.db.WithContext(ctx).
		Where("id = ?", doerId).
		Model(user).
		Updates(user).Error
	if err == nil {
		key := fmt.Sprintf("%s%d", cacheUserIdPrefix, doerId)
		_, err = m.CacheRedis.DelCtx(ctx, key)
	}
	return err
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
