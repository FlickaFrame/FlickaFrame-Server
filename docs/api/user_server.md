### 1. "Follow a user"

1. route definition

- Url: /api/v1/user/follow_action/:user_id
- Method: PUT
- Request: `FollowReq`
- Response: `FollowResp`

2. request definition



```golang
type FollowReq struct {
	ContextUserId int64 `path:"user_id"`
}
```


3. response definition



```golang
type FollowResp struct {
}
```

### 2. "Unfollow a user"

1. route definition

- Url: /api/v1/user/follow_action/:user_id
- Method: DELETE
- Request: `UnFollowReq`
- Response: `UnFollowResp`

2. request definition



```golang
type UnFollowReq struct {
	ContextUserId int64 `path:"user_id"`
}
```


3. response definition



```golang
type UnFollowResp struct {
}
```

### 3. "ListMyFollowers list the followers user of the authenticated user"

1. route definition

- Url: /api/v1/user/me/followers
- Method: GET
- Request: `ListFollowReq`
- Response: `ListFollowUserResp`

2. request definition



```golang
type ListFollowReq struct {
	ContextUserId int64 `path:"user_id,optional"`
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type ListUserOption struct {
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}
```


3. response definition



```golang
type ListFollowUserResp struct {
	FollowUser []*FollowUser `json:"users"`
}
```

### 4. "ListMyFollowing list the following users of the authenticated user"

1. route definition

- Url: /api/v1/user/me/following
- Method: GET
- Request: `ListFollowReq`
- Response: `ListFollowUserResp`

2. request definition



```golang
type ListFollowReq struct {
	ContextUserId int64 `path:"user_id,optional"`
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type ListUserOption struct {
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}
```


3. response definition



```golang
type ListFollowUserResp struct {
	FollowUser []*FollowUser `json:"users"`
}
```

### 5. "ListFollowers list the given user&#39;s followers"

1. route definition

- Url: /api/v1/user/:user_id/followers
- Method: GET
- Request: `ListFollowReq`
- Response: `ListFollowUserResp`

2. request definition



```golang
type ListFollowReq struct {
	ContextUserId int64 `path:"user_id,optional"`
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type ListUserOption struct {
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}
```


3. response definition



```golang
type ListFollowUserResp struct {
	FollowUser []*FollowUser `json:"users"`
}
```

### 6. "ListFollowing list the users that the given user is following"

1. route definition

- Url: /api/v1/user/:user_id/following
- Method: GET
- Request: `ListFollowReq`
- Response: `ListFollowUserResp`

2. request definition



```golang
type ListFollowReq struct {
	ContextUserId int64 `path:"user_id,optional"`
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type ListUserOption struct {
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}
```


3. response definition



```golang
type ListFollowUserResp struct {
	FollowUser []*FollowUser `json:"users"`
}
```

### 7. "Get Current Logined User Detail Info"

1. route definition

- Url: /api/v1/user/detail
- Method: GET
- Request: `UserDetailInfoReq`
- Response: `UserDetailInfoResp`

2. request definition



```golang
type UserDetailInfoReq struct {
	ContextUserId int64 `path:"userId,optional"`
}
```


3. response definition



```golang
type UserDetailInfoResp struct {
	ID string `json:"userId" copier:"IDString"` // 用户ID
	NickName string `json:"nickName"` // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan string `json:"slogan"` // 个性签名
	Gender int64 `json:"gender"` // 性别
	Age int `json:"age"`
	BackgroundUrl string `json:"backgroundUrl,optional"` //用户主页背景图
	FollowingCount int `json:"followingCount"` // 关注数
	FollowerCount int `json:"followerCount"` // 粉丝数
	LikeCount int `json:"likeCount"` // 获赞数量
	PublishedVideoCount int `json:"publishVideoCount"` // 发布作品数量
	LikeVideoCount int `json:"likeVideoCount"` // 点赞作品数量
	CollectionsVideoCount int `json:"collectionsVideoCount"` // 收藏作品数量
	IsFollow bool `json:"isFollow"` // 是否关注
}

type UserBasicInfo struct {
	ID string `json:"userId" copier:"IDString"` // 用户ID
	NickName string `json:"nickName"` // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan string `json:"slogan"` // 个性签名
	Gender int64 `json:"gender"` // 性别
	Age int `json:"age"`
	BackgroundUrl string `json:"backgroundUrl,optional"` //用户主页背景图
}

type UserStatisticalInfo struct {
	FollowingCount int `json:"followingCount"` // 关注数
	FollowerCount int `json:"followerCount"` // 粉丝数
	LikeCount int `json:"likeCount"` // 获赞数量
	PublishedVideoCount int `json:"publishVideoCount"` // 发布作品数量
	LikeVideoCount int `json:"likeVideoCount"` // 点赞作品数量
	CollectionsVideoCount int `json:"collectionsVideoCount"` // 收藏作品数量
}

type UserInteractionInfo struct {
	IsFollow bool `json:"isFollow"` // 是否关注
}
```

### 8. "Update User Info"

1. route definition

- Url: /api/v1/user/info
- Method: PUT
- Request: `UpdateInfoReq`
- Response: `UpdateInfoReq`

2. request definition



```golang
type UpdateInfoReq struct {
	NickName string `json:"nickName"` // 用户名
	Slogan string `json:"slogan"` // 个性签名
	Gender int `json:"gender"` // 性别
	Age int `json:"age"` // 年龄
	AvatarUrl string `json:"avatarUrl"` // 头像
	BackgroundUrl string `json:"backgroundUrl"` // 用户主页背景图
}
```


3. response definition



```golang
type UpdateInfoReq struct {
	NickName string `json:"nickName"` // 用户名
	Slogan string `json:"slogan"` // 个性签名
	Gender int `json:"gender"` // 性别
	Age int `json:"age"` // 年龄
	AvatarUrl string `json:"avatarUrl"` // 头像
	BackgroundUrl string `json:"backgroundUrl"` // 用户主页背景图
}
```

### 9. "Update User Password"

1. route definition

- Url: /api/v1/user/updatepwd
- Method: POST
- Request: `UpdatePasswordReq`
- Response: `UpdatePasswordResp`

2. request definition



```golang
type UpdatePasswordReq struct {
}
```


3. response definition



```golang
type UpdatePasswordResp struct {
}
```

### 10. "Register User"

1. route definition

- Url: /api/v1/user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterResp`

2. request definition



```golang
type RegisterReq struct {
	Phone string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName,option"`
}
```


3. response definition



```golang
type RegisterResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 11. "Login User"

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
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 12. "Get User Detail Info"

1. route definition

- Url: /api/v1/user/detail/:userId
- Method: GET
- Request: `UserDetailInfoReq`
- Response: `UserDetailInfoResp`

2. request definition



```golang
type UserDetailInfoReq struct {
	ContextUserId int64 `path:"userId,optional"`
}
```


3. response definition



```golang
type UserDetailInfoResp struct {
	ID string `json:"userId" copier:"IDString"` // 用户ID
	NickName string `json:"nickName"` // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan string `json:"slogan"` // 个性签名
	Gender int64 `json:"gender"` // 性别
	Age int `json:"age"`
	BackgroundUrl string `json:"backgroundUrl,optional"` //用户主页背景图
	FollowingCount int `json:"followingCount"` // 关注数
	FollowerCount int `json:"followerCount"` // 粉丝数
	LikeCount int `json:"likeCount"` // 获赞数量
	PublishedVideoCount int `json:"publishVideoCount"` // 发布作品数量
	LikeVideoCount int `json:"likeVideoCount"` // 点赞作品数量
	CollectionsVideoCount int `json:"collectionsVideoCount"` // 收藏作品数量
	IsFollow bool `json:"isFollow"` // 是否关注
}

type UserBasicInfo struct {
	ID string `json:"userId" copier:"IDString"` // 用户ID
	NickName string `json:"nickName"` // 用户名
	AvatarUrl string `json:"avatarUrl"` // 头像
	Slogan string `json:"slogan"` // 个性签名
	Gender int64 `json:"gender"` // 性别
	Age int `json:"age"`
	BackgroundUrl string `json:"backgroundUrl,optional"` //用户主页背景图
}

type UserStatisticalInfo struct {
	FollowingCount int `json:"followingCount"` // 关注数
	FollowerCount int `json:"followerCount"` // 粉丝数
	LikeCount int `json:"likeCount"` // 获赞数量
	PublishedVideoCount int `json:"publishVideoCount"` // 发布作品数量
	LikeVideoCount int `json:"likeVideoCount"` // 点赞作品数量
	CollectionsVideoCount int `json:"collectionsVideoCount"` // 收藏作品数量
}

type UserInteractionInfo struct {
	IsFollow bool `json:"isFollow"` // 是否关注
}
```

### 13. "List User Ranking"

1. route definition

- Url: /api/v1/user/ranking
- Method: GET
- Request: `RankingReq`
- Response: `RankingResp`

2. request definition



```golang
type RankingReq struct {
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}

type ListUserOption struct {
	PageSize int `form:"pageSize,default=10"` // 分页大小,默认为 10
	Page int `form:"page,default=1"` // 当前页码,默认为 1
	ListAll bool `form:"listAll,default=false"` // 是否列出所有,默认为 false
}
```


3. response definition



```golang
type RankingResp struct {
	Users []*UserBasicInfo `json:"users"`
}
```

