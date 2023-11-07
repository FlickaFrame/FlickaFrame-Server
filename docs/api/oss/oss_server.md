### 1. "Endpoint"

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

### 2. "Get Upload Token"

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

