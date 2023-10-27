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

