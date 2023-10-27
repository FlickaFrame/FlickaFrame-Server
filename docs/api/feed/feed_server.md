### 1. "feed"

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

### 2. "category"

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

