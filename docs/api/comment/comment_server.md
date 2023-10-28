### 1. "Get a specific comment of an Video"

1. route definition

- Url: /api/v1/video/:video_id/comments/:comment_id
- Method: GET
- Request: `GetVideoCommentReq`
- Response: `GetVideoCommentResp`

2. request definition



```golang
type GetVideoCommentReq struct {
	VideoId uint `json:"videoId" path:"video_id"`
	CommentId uint `json:"commentId" path:"comment_id"`
}
```


3. response definition



```golang
type GetVideoCommentResp struct {
	Commnent *Commnent `json:"comment"`
}
```

### 2. "List comments of an Video"

1. route definition

- Url: /api/v1/video/:video_id/comments
- Method: GET
- Request: `ListVideoCommentsReq`
- Response: `ListVideoCommentsResp`

2. request definition



```golang
type ListVideoCommentsReq struct {
	VideoId uint `json:"videoId" path:"video_id"`
}
```


3. response definition



```golang
type ListVideoCommentsResp struct {
	CreatedAt string `json:"createdAt"`
	UserID string `json:"userId"`
	Comments []Commnent `json:"comments"`
}
```

### 3. "Create a comment for an Video"

1. route definition

- Url: /api/v1/video/:video_id/comments
- Method: PUT
- Request: `CreateVideoCommentReq`
- Response: `CreateVideoCommentResp`

2. request definition



```golang
type CreateVideoCommentReq struct {
	Content string `json:"content"`
	VideoId uint `json:"videoId" path:"video_id"`
}
```


3. response definition



```golang
type CreateVideoCommentResp struct {
	Commnent *Commnent `json:"comment"`
}
```

### 4. "Delete a comment of an Video"

1. route definition

- Url: /api/v1/comments/:comment_id
- Method: DELETE
- Request: `DeleteVideoCommentReq`
- Response: `DeleteVideoCommentResp`

2. request definition



```golang
type DeleteVideoCommentReq struct {
	VideoId uint `json:"videoId" path:"video_id"`
	CommentId uint `json:"commentId" path:"comment_id"`
}
```


3. response definition



```golang
type DeleteVideoCommentResp struct {
}
```

### 5. "Edit comments of an Video"

1. route definition

- Url: /api/v1/comments/:comment_id
- Method: POST
- Request: `EditVideoCommentReq`
- Response: `EditVideoCommentResp`

2. request definition



```golang
type EditVideoCommentReq struct {
	VideoId uint `json:"videoId" path:"video_id"`
	CommentId uint `json:"commentId" path:"comment_id"`
	Content string `json:"content"`
}
```


3. response definition



```golang
type EditVideoCommentResp struct {
}
```

### 6. "Create a reply for an Comment"

1. route definition

- Url: /api/v1/video/:video_id/reply
- Method: PUT
- Request: `CreateReplyCommentReq`
- Response: `CreateReplyCommentResp`

2. request definition



```golang
type CreateReplyCommentReq struct {
	Content string `json:"content"`
	VideoId uint `json:"videoId" path:"video_id"`
	ParentId uint `json:"parentId"`
	TargetId uint `json:"targetId"`
}
```


3. response definition



```golang
type CreateReplyCommentResp struct {
	Commnent *Commnent `json:"comment"`
}
```

