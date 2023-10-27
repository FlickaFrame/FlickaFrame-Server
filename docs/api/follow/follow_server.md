### 1. "Follow a user"

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

### 2. "Unfollow a user"

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

### 3. "Check whether a user is followed by the authenticated user"

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

### 4. "List the users that the authenticated user is following"

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

### 5. "List the users that the authenticated user is following"

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

### 6. "ListFollowers list the given user&#39;s followers"

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

### 7. "ListFollowing list the users that the given user is following"

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

### 8. "CheckFollowing check if one user is following another user"

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

