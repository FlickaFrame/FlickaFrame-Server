### 1. "Get a specific comment of an Video"

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

### 2. "List comments of an Video"

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

### 3. "Create a comment for an Video"

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

### 4. "Create a child comment for comment"

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

### 5. "Create a reply for an child Comment"

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

### 6. "Delete a comment of an Video"

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

### 7. "Edit comments of an Video"

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

