### 1. "register"

1. route definition

- Url: /api/v1/user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterResp`

2. request definition



```golang
type RegisterReq struct {
	Phone string `json:"phone" validate:"e164,required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nick_name,option"`
}
```


3. response definition



```golang
type RegisterResp struct {
	AccessToken string `json:"access_token"`
	AccessExpire int64 `json:"access_expire"`
	RefreshAfter int64 `json:"refresh_after"`
}
```

### 2. "login"

1. route definition

- Url: /api/v1/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	AccessToken string `json:"access_token"`
	AccessExpire int64 `json:"access_expire"`
	RefreshAfter int64 `json:"refresh_after"`
}
```

### 3. "get user info"

1. route definition

- Url: /api/v1/user/detail
- Method: GET
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
}
```


3. response definition



```golang
type UserInfoResp struct {
	UserInfo User `json:"user_info"`
}

type User struct {
	ID int64 `json:"id"`
	Phone string `json:"phone"`
	NickName string `json:"nick_name"`
	Sex int64 `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Info string `json:"info"`
}
```

### 4. "feed"

1. route definition

- Url: /api/v1/feed
- Method: GET
- Request: `FeedReq`
- Response: `FeedResp`

2. request definition



```golang
type FeedReq struct {
	LatestTime int64 `json:"latest_time,optional" form:"latestTime,optional"` // 最新视频时间(毫秒时间戳)
	Limit int `json:"limit,optional" form:"limit,optional"` // 请求数量
	Token string `json:"token,optional" form:"token,optional"` // 登录token
	AuthorID uint `json:"author_id,optional" form:"authorID,optional"` // 作者ID
	Tag string `json:"tag,optional" form:"tag,optional"` // 标签
	CategoryID uint `json:"category_id,optional" form:"category_id,optional"` // 分类
}
```


3. response definition



```golang
type FeedResp struct {
	VideoList []*Video `json:"video_list"`
	NextTime int64 `json:"next_time"` // 下次请求时间(毫秒时间戳)
	Length int `json:"length"` // 视频列表长度
}
```

### 5. "category"

1. route definition

- Url: /api/v1/category
- Method: GET
- Request: `CategoryReq`
- Response: `CategoryResp`

2. request definition



```golang
type CategoryReq struct {
}
```


3. response definition



```golang
type CategoryResp struct {
	CategoryList []*Category `json:"category_list"`
}
```

