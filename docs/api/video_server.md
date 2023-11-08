### 1. "Create Video"

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

### 2. "Delete Video"

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

### 3. "List video of following User"

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

### 4. "Get Video Info"

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

### 5. "Get Video Category List"

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

### 6. "Home Video Feed"

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

### 7. "Search Video By Keyword"

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

