### 1. "通知查询"

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

