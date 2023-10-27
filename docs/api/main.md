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

### 6. "Follow a user"

1. route definition

- Url: /api/v1/user/following/:user_id
- Method: PUT
- Request: `FollowReq`
- Response: `FollowResp`

2. request definition



```golang
type FollowReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"关注用户id" validate:"required"`
}
```


3. response definition



```golang
type FollowResp struct {
}
```

### 7. "Unfollow a user"

1. route definition

- Url: /api/v1/user/following/:user_id
- Method: DELETE
- Request: `UnFollowReq`
- Response: `UnFollowResp`

2. request definition



```golang
type UnFollowReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"关注用户id" validate:"required"`
}
```


3. response definition



```golang
type UnFollowResp struct {
}
```

### 8. "Check whether a user is followed by the authenticated user"

1. route definition

- Url: /api/v1/user/following/:user_id
- Method: GET
- Request: `CheckMyFollowingReq`
- Response: `CheckMyFollowingResp`

2. request definition



```golang
type CheckMyFollowingReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"关注用户id" validate:"required"`
}
```


3. response definition



```golang
type CheckMyFollowingResp struct {
	Status bool `json:"status" desc:"是否关注"`
}
```

### 9. "List the users that the authenticated user is following"

1. route definition

- Url: /api/v1/user/followers
- Method: GET
- Request: `ListMyFollowersReq`
- Response: `ListMyFollowersResp`

2. request definition



```golang
type ListMyFollowersReq struct {
	Page int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}
```


3. response definition



```golang
type ListMyFollowersResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}
```

### 10. "List the users that the authenticated user is following"

1. route definition

- Url: /api/v1/user/following
- Method: GET
- Request: `ListMyFollowingReq`
- Response: `ListMyFollowingResp`

2. request definition



```golang
type ListMyFollowingReq struct {
	Page int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}
```


3. response definition



```golang
type ListMyFollowingResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}
```

### 11. "ListFollowers list the given user&#39;s followers"

1. route definition

- Url: /api/v1/users/:user_id/followers
- Method: GET
- Request: `ListFollowersReq`
- Response: `ListFollowersResp`

2. request definition



```golang
type ListFollowersReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"用户id" validate:"required"`
	Page int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}
```


3. response definition



```golang
type ListFollowersResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}
```

### 12. "ListFollowing list the users that the given user is following"

1. route definition

- Url: /api/v1/users/:user_id/following
- Method: GET
- Request: `ListFollowingReq`
- Response: `ListFollowingResp`

2. request definition



```golang
type ListFollowingReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"用户id" validate:"required"`
	Page int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}
```


3. response definition



```golang
type ListFollowingResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}
```

### 13. "CheckFollowing check if one user is following another user"

1. route definition

- Url: /api/v1/users/:doer_id/following/:context_user_id
- Method: GET
- Request: `CheckFollowingReq`
- Response: `CheckFollowingResp`

2. request definition



```golang
type CheckFollowingReq struct {
	DoerUserId uint `json:"doer_user_id" path:"doer_user_id" desc:"用户id" validate:"required"`
	ContextUserId uint `json:"context_user_id" path:"doer_user_id" desc:"用户id" validate:"required"`
}
```


3. response definition



```golang
type CheckFollowingResp struct {
	Status bool `json:"status" desc:"是否关注"`
}
```

### 14. "Count the number of followers and following for the given user"

1. route definition

- Url: /api/v1/users/follow/:user_id/count
- Method: GET
- Request: `CountFollowReq`
- Response: `CountFollowResp`

2. request definition



```golang
type CountFollowReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"用户id" validate:"required"`
}
```


3. response definition



```golang
type CountFollowResp struct {
	FollowingCount int64 `json:"following_count" desc:"关注数量"`
	FollowersCount int64 `json:"follower_count" desc:"粉丝数量"`
}
```

