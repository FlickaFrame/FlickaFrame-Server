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

### 14. "Create Video"

1. route definition

- Url: /api/v1/video/create
- Method: POST
- Request: `CreateVideoReq`
- Response: `CreateVideoResp`

2. request definition



```golang
type CreateVideoReq struct {
	Title string `json:"title"` // 视频标题
	PlayUrl string `json:"playUrl"` // 视频播放地址
	ThumbUrl string `json:"thumbUrl"` // 视频封面地址
	Description string `json:"description,optional"` // 视频描述
	CategoryID string `json:"category"` // 视频分类
	Tags []string `json:"tags"` // 视频标签
	PublishTime int64 `json:"publishTime,optional"` // 视频发布时间(毫秒时间戳)
	VideoKey string `json:"videoKey,optional"` // 视频上传key
	Visibility int `json:"visibility"` // 视频可见性(1:公开,2:私密)
	VideoDuration float32 `json:"videoDuration"` // 视频时长(秒)
	VideoHeight uint `json:"videoHeight"` // 视频高度(像素)
	VideoWidth uint `json:"videoWidth"` // 视频宽度(像素)
}
```


3. response definition



```golang
type CreateVideoResp struct {
}
```

### 15. "Delete Video"

1. route definition

- Url: /api/v1/video/:videoId
- Method: DELETE
- Request: `DeleteVideoReq`
- Response: `DeleteVideoResp`

2. request definition



```golang
type DeleteVideoReq struct {
	VideoID int64 `path:"videoId"`
}
```


3. response definition



```golang
type DeleteVideoResp struct {
}
```

### 16. "List video of following User"

1. route definition

- Url: /api/v1/video/following
- Method: GET
- Request: `FeedReq`
- Response: `FeedResp`

2. request definition



```golang
type FeedReq struct {
	Cursor int64 `form:"cursor,default=0"` // 最新视频时间(毫秒时间戳)
	Limit int `form:"limit,default=10"` // 请求数量
	AuthorID string `form:"authorID,default=0"` // 作者ID(是否根据用户ID过滤)
	Tag string `form:"tag,optional"` // 标签(是否根据标签过滤)
	CategoryID string `form:"categoryId,optional"` // 分类(是否根据分类过滤)
}
```


3. response definition



```golang
type FeedResp struct {
	Next string `json:"next"` // 请求游标
	List []*FeedVideoItem `json:"list"` // 视频列表
	IsEnd bool `json:"isEnd"` // 是否已到最后一页
}
```

### 17. "Get Video Info"

1. route definition

- Url: /api/v1/video/detail/:videoId
- Method: GET
- Request: `GetVideoInfoReq`
- Response: `GetVideoInfoResp`

2. request definition



```golang
type GetVideoInfoReq struct {
	VideoId int64 `path:"videoId"`
}
```


3. response definition



```golang
type GetVideoInfoResp struct {
	Video *VideoBasicInfo `json:"video"`
}
```

### 18. "Get Video Category List"

1. route definition

- Url: /api/v1/video/category
- Method: GET
- Request: `-`
- Response: `CategoryResp`

2. request definition



3. response definition



```golang
type CategoryResp struct {
	CategoryList []*Category `json:"categoryList"`
}
```

### 19. "Home Video Feed"

1. route definition

- Url: /api/v1/video/feed
- Method: GET
- Request: `FeedReq`
- Response: `FeedResp`

2. request definition



```golang
type FeedReq struct {
	Cursor int64 `form:"cursor,default=0"` // 最新视频时间(毫秒时间戳)
	Limit int `form:"limit,default=10"` // 请求数量
	AuthorID string `form:"authorID,default=0"` // 作者ID(是否根据用户ID过滤)
	Tag string `form:"tag,optional"` // 标签(是否根据标签过滤)
	CategoryID string `form:"categoryId,optional"` // 分类(是否根据分类过滤)
}
```


3. response definition



```golang
type FeedResp struct {
	Next string `json:"next"` // 请求游标
	List []*FeedVideoItem `json:"list"` // 视频列表
	IsEnd bool `json:"isEnd"` // 是否已到最后一页
}
```

### 20. "Search Video By Keyword"

1. route definition

- Url: /api/v1/video/search
- Method: POST
- Request: `VideoSearchReq`
- Response: `VideoSearchResp`

2. request definition



```golang
type VideoSearchReq struct {
	Keyword string `json:"keyword"` // 搜索关键字
	Offset int64 `json:"offset,optional"` // 偏移量
	Limit int64 `json:"limit,optional"` // 请求数量
}
```


3. response definition



```golang
type VideoSearchResp struct {
	Videos []*FeedVideoItem `json:"list"` // 视频列表
	Query string `json:"query"` // 搜索关键字
	ProcessingTimeMs int64 `json:"processingTimeMs"` // 搜索耗时(毫秒)
	Offset int64 `json:"offset"` // 偏移量
	Limit int64 `json:"limit"` // 请求数量
	EstimatedTotalHits int64 `json:"estimatedTotalHits"` // 搜索结果总数
}
```

### 21. "Get a specific comment of an Video"

1. route definition

- Url: /api/v1/video/:video_id/comments/:comment_id
- Method: GET
- Request: `GetVideoCommentReq`
- Response: `GetVideoCommentResp`

2. request definition



```golang
type GetVideoCommentReq struct {
	CommentId string `json:"commentId" path:"comment_id"`
}
```


3. response definition



```golang
type GetVideoCommentResp struct {
	Commnent *ParentComment `json:"comment"`
}
```

### 22. "List comments of an Video"

1. route definition

- Url: /api/v1/video/:video_id/comments
- Method: GET
- Request: `ListVideoCommentsReq`
- Response: `ListVideoCommentsResp`

2. request definition



```golang
type ListVideoCommentsReq struct {
	VideoId string `path:"video_id,optional"`
}
```


3. response definition



```golang
type ListVideoCommentsResp struct {
	Comments []*ParentComment `json:"comments"`
}
```

### 23. "Create a comment for an Video"

1. route definition

- Url: /api/v1/comment/video
- Method: POST
- Request: `CreateVideoCommentReq`
- Response: `CreateVideoCommentResp`

2. request definition



```golang
type CreateVideoCommentReq struct {
	VideoId string `json:"videoId"`
	Content string `json:"content"`
	AtUsersId []string `json:"atUsersId,optional"`
}
```


3. response definition



```golang
type CreateVideoCommentResp struct {
	Comment *ParentComment `json:"comment"`
}
```

### 24. "Create a child comment for comment"

1. route definition

- Url: /api/v1/comment/parent
- Method: POST
- Request: `CreateChildCommentReq`
- Response: `CreateChildCommentResp`

2. request definition



```golang
type CreateChildCommentReq struct {
	VideoId string `json:"videoId"`
	Content string `json:"content"`
	AtUsersId []string `json:"atUsersId,optional"`
	ParentCommentId string `json:"parentCommentId"`
	TargetCommentId string `json:"targetCommentId,optional"`
}
```


3. response definition



```golang
type CreateChildCommentResp struct {
	Comment *ChildComment `json:"comment"`
}
```

### 25. "Create a reply for an child Comment"

1. route definition

- Url: /api/v1/comment/child
- Method: PUT
- Request: `CreateReplyCommentReq`
- Response: `CreateReplyCommentResp`

2. request definition



```golang
type CreateReplyCommentReq struct {
	VideoId string `json:"videoId"`
	Content string `json:"content"`
	AtUsersId []string `json:"atUsersId,optional"`
	ParentCommentId string `json:"parentCommentId,optional"`
	TargetCommentId string `json:"targetCommentId,optional"`
}
```


3. response definition



```golang
type CreateReplyCommentResp struct {
	Comment *ChildComment `json:"comment"`
}
```

### 26. "Delete a comment of an Video"

1. route definition

- Url: /api/v1/comment/:comment_id
- Method: DELETE
- Request: `DeleteVideoCommentReq`
- Response: `DeleteVideoCommentResp`

2. request definition



```golang
type DeleteVideoCommentReq struct {
	CommentId string `path:"comment_id"`
	Type string `form:"type"`
}
```


3. response definition



```golang
type DeleteVideoCommentResp struct {
}
```

### 27. "Edit comments of an Video"

1. route definition

- Url: /api/v1/comment/:comment_id
- Method: POST
- Request: `EditVideoCommentReq`
- Response: `EditVideoCommentResp`

2. request definition



```golang
type EditVideoCommentReq struct {
	VideoId string `json:"videoId" path:"video_id"`
	CommentId string `json:"commentId" path:"comment_id"`
	Content string `json:"content"`
}
```


3. response definition



```golang
type EditVideoCommentResp struct {
}
```

### 28. "Favorite a Video"

1. route definition

- Url: /api/v1/favorite/video/:targetId
- Method: PUT
- Request: `FavoriteReq`
- Response: `FavoriteResp`

2. request definition



```golang
type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}
```


3. response definition



```golang
type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}
```

### 29. "UnFavorite a Video"

1. route definition

- Url: /api/v1/favorite/video/:targetId
- Method: DELETE
- Request: `FavoriteReq`
- Response: `FavoriteResp`

2. request definition



```golang
type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}
```


3. response definition



```golang
type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}
```

### 30. "Check Video Favorite Status"

1. route definition

- Url: /api/v1/favorite/video/:targetId
- Method: GET
- Request: `FavoriteReq`
- Response: `FavoriteResp`

2. request definition



```golang
type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}
```


3. response definition



```golang
type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}
```

### 31. "Favorite a Comment"

1. route definition

- Url: /api/v1/favorite/comment/:targetId
- Method: PUT
- Request: `FavoriteReq`
- Response: `FavoriteResp`

2. request definition



```golang
type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}
```


3. response definition



```golang
type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}
```

### 32. "UnFavorite a Comment"

1. route definition

- Url: /api/v1/favorite/comment/:targetId
- Method: DELETE
- Request: `FavoriteReq`
- Response: `FavoriteResp`

2. request definition



```golang
type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}
```


3. response definition



```golang
type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}
```

### 33. "Check comment Favorite Status"

1. route definition

- Url: /api/v1/favorite/comment/:targetId
- Method: GET
- Request: `FavoriteReq`
- Response: `FavoriteResp`

2. request definition



```golang
type FavoriteReq struct {
	TargetId string `path:"targetId" copier:"IDString"`
}
```


3. response definition



```golang
type FavoriteResp struct {
	IsFavorite bool `json:"isFavorite"`
}
```

### 34. "Endpoint"

1. route definition

- Url: /api/v1/oss/endpoint
- Method: GET
- Request: `-`
- Response: `ossEndpointResponse`

2. request definition



3. response definition



```golang
type OssEndpointResponse struct {
	EndPoint string `json:"endpoint"`
}
```

### 35. "Get Upload Token"

1. route definition

- Url: /api/v1/oss/uptoken
- Method: GET
- Request: `CreateUpTokenReq`
- Response: `CreateUpTokenResp`

2. request definition



```golang
type CreateUpTokenReq struct {
	UploadType string `form:"uploadType"` // 上传类型(video:视频,cover:封面,avatar:头像)
}
```


3. response definition



```golang
type CreateUpTokenResp struct {
	UpToken string `json:"upToken"` // 上传凭证
	Expires int64 `json:"expires"` // 上传凭证过期时间(秒)
}
```

### 36. "通知查询"

1. route definition

- Url: /api/v1/notice/type
- Method: GET
- Request: `FollowNoticeReq`
- Response: `FollowNoticeResp`

2. request definition



```golang
type FollowNoticeReq struct {
	Cursor int64 `form:"cursor,default=0"` // 最新通知时间(毫秒时间戳)
	Limit int `form:"limit,default=10"` // 请求数量
	NoticeType string `form:"noticeType"` // 通知类型
}
```


3. response definition



```golang
type FollowNoticeResp struct {
	Next string `json:"next"` // 请求游标
	List []*NoticeItem `json:"list"` // 通知列表
	IsEnd bool `json:"isEnd"` // 是否已到最后一页
}
```

