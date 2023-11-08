### 1. "Favorite a Video"

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

### 2. "UnFavorite a Video"

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

### 3. "Check Video Favorite Status"

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

### 4. "Favorite a Comment"

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

### 5. "UnFavorite a Comment"

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

### 6. "Check comment Favorite Status"

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

