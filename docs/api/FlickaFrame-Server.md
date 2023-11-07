---
title: FlickaFrame-Server copy v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# FlickaFrame-Serve

> v1.0.0

Base URLs:

* <a href="http://localhost:8080">开发环境: http://localhost:8080</a>

# Authentication

- HTTP Authentication, scheme: bearer

# oss

<a id="opIdEndpoint"></a>

## GET Get Oss Server Endpoint

GET /api/v1/oss/endpoint

返回oss服务的域名地址

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "ok",
  "success": true,
  "data": {
    "endpoint": "http://s2i8a2ssf.hn-bkt.clouddn.com"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» endpoint|string|true|none||none|

<a id="opIdCreateUpToken"></a>

## GET Get Upload Token

GET /api/v1/oss/uptoken?uploadType=video

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uploadType|query|string| 是 | 上传类型(video:视频,cover:封面,avatar:头像)|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "ok",
  "success": true,
  "data": {
    "upToken": "4gzQF7x9ZTqloWqIjSzfknoC0L0xONQH39gp8ZFm:ckRqLtV08gQHFItIpcxxgAPj4TE=:eyJzY29wZSI6ImNhaWNhbmRvbmc6dmlkZW8vIiwiaXNQcmVmaXhhbFNjb3BlIjoxLCJkZWFkbGluZSI6MTY5OTE5OTY0M30=",
    "expires": 3600
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» upToken|string|true|none||none|
|»» expires|integer|true|none||none|

## POST Upload Video To OSS

POST /

> Body 请求参数

```yaml
file: string
token: 4gzQF7x9ZTqloWqIjSzfknoC0L0xONQH39gp8ZFm:492x5Kn8Z41M4QXpvNpFDeDURts=:eyJzY29wZSI6ImNhaWNhbmRvbmc6dmlkZW8vIiwiaXNQcmVmaXhhbFNjb3BlIjoxLCJkZWFkbGluZSI6MTY5OTIwMjc1N30=
key: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|
|» token|body|string| 否 |none|
|» key|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# user

<a id="opIdCurrentUserInfo"></a>

## GET Get Current User Detail Info

GET /api/v1/user/detail

获取当前用户的详细个人信息

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "success": true,
  "data": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "age": 0,
    "followingCount": 0,
    "followerCount": 0,
    "likeCount": 0,
    "publishVideoCount": 0,
    "likeVideoCount": 0,
    "collectionsVideoCount": 0,
    "isFollow": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» userId|string|true|none||none|
|»» nickName|string|true|none||none|
|»» avatarUrl|string|true|none||none|
|»» slogan|string|true|none||none|
|»» gender|integer|true|none||none|
|»» age|integer|true|none||none|
|»» followingCount|integer|true|none||none|
|»» followerCount|integer|true|none||none|
|»» likeCount|integer|true|none||none|
|»» publishVideoCount|integer|true|none||none|
|»» likeVideoCount|integer|true|none||none|
|»» collectionsVideoCount|integer|true|none||none|
|»» isFollow|boolean|true|none||none|

<a id="opIdGetUserDetailInfo"></a>

## GET Get User Detail Info

GET /api/v1/user/detail/{userId}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|userId|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "success": true,
  "data": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "age": 0,
    "followingCount": 0,
    "followerCount": 0,
    "likeCount": 0,
    "publishVideoCount": 0,
    "likeVideoCount": 0,
    "collectionsVideoCount": 0,
    "isFollow": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» userId|string|true|none||none|
|»» nickName|string|true|none||none|
|»» avatarUrl|string|true|none||none|
|»» slogan|string|true|none||none|
|»» gender|integer|true|none||none|
|»» age|integer|true|none||none|
|»» followingCount|integer|true|none||none|
|»» followerCount|integer|true|none||none|
|»» likeCount|integer|true|none||none|
|»» publishVideoCount|integer|true|none||none|
|»» likeVideoCount|integer|true|none||none|
|»» collectionsVideoCount|integer|true|none||none|
|»» isFollow|boolean|true|none||none|

<a id="opIdUpdateInfo"></a>

## POST Update User Info

POST /api/v1/user/info

> Body 请求参数

```json
{
  "slogan": "velit occaecat laborum ut dolore",
  "age": 99,
  "gender": 53,
  "nickName": "程敏"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» slogan|body|string| 是 |none|
|» age|body|integer| 是 |none|
|» gender|body|integer| 是 |none|
|» nickName|body|string| 是 |none|
|» avatarUrl|body|string| 是 |用户头像url|

> 返回示例

> 200 Response

```json
{
  "nickName": "string",
  "slogan": "string",
  "gender": 0,
  "age": 0,
  "password": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[UpdateInfoReq](#schemaupdateinforeq)|

<a id="opIdlogin"></a>

## POST Login User

POST /api/v1/user/login

> Body 请求参数

```json
{
  "phone": "13280290697",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» phone|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "ok",
  "success": true,
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTk4MDEzMzYsImlhdCI6MTY5OTE5NjUzNiwiand0VXNlcklkIjozOX0.Onu56d6u0KCqW8fylYup1fDDVVQnUW5yCa0bL3iumO8",
    "accessExpire": 1699801336,
    "refreshAfter": 1699498936
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» accessToken|string|true|none||none|
|»» accessExpire|integer|true|none||none|
|»» refreshAfter|integer|true|none||none|

<a id="opIdRanking"></a>

## GET List User Ranking

GET /api/v1/user/ranking

> 返回示例

> 200 Response

```json
{
  "users": [
    {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string",
      "slogan": "string",
      "gender": 0,
      "age": 0
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[RankingResp](#schemarankingresp)|

<a id="opIdregister"></a>

## POST Register User

POST /api/v1/user/register

> Body 请求参数

```json
{
  "phone": "string",
  "password": "string",
  "nickName": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|[RegisterReq](#schemaregisterreq)| 否 | RegisterReq|none|

> 返回示例

> 200 Response

```json
{
  "accessToken": "string",
  "accessExpire": 0,
  "refreshAfter": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[RegisterResp](#schemaregisterresp)|

<a id="opIdUpdatePassword"></a>

## POST Update User Password

POST /api/v1/user/updatepwd

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|[UpdatePasswordReq](#schemaupdatepasswordreq)| 否 | UpdatePasswordReq|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[UpdatePasswordResp](#schemaupdatepasswordresp)|

# video

<a id="opIdCategory"></a>

## GET Get Video Category List

GET /api/v1/video/category

> 返回示例

> 200 Response

```json
{
  "categoryList": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[CategoryResp](#schemacategoryresp)|

<a id="opIdCreateVideo"></a>

## POST Create Video

POST /api/v1/video/create

> Body 请求参数

```json
{
  "title": "军只提代",
  "playUrl": "http://hur.cm/hujk",
  "thumbUrl": "http://mclwgjhwc.kn/gieo",
  "description": "做问县目政济种是八对历国第较。",
  "category": "1",
  "tags": [
    "#Tag1",
    "#Tag2"
  ],
  "visibility": 1,
  "videoKey": "123123123",
  "publishTime": 548514652930,
  "videoDuration": 13.111
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» playUrl|body|string| 是 | 视频地址|none|
|» thumbUrl|body|string| 是 | 封面地址|none|
|» title|body|string| 是 | 标题|none|
|» description|body|string| 是 | 描述|none|
|» category|body|string| 是 | 分类|none|
|» tags|body|[string]| 是 | 标签|none|
|» visibility|body|integer| 是 | 视频可见性|(1:公开,2:私密)|
|» videoDuration|body|number| 是 | 视频时长|秒|
|» videoHeight|body|integer| 是 | 视频高度|像素|
|» videoWidth|body|integer| 是 | 视频宽度|像素|
|» publishTime|body|integer| 是 | 发布时间(定时发布)|毫秒时间戳|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[CreateVideoResp](#schemacreatevideoresp)|

<a id="opIdDeleteVideo"></a>

## DELETE Delete Video

DELETE /api/v1/video/{videoId}

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|videoId|path|string| 是 ||none|
|body|body|[DeleteVideoReq](#schemadeletevideoreq)| 否 | DeleteVideoReq|none|

> 返回示例

> A successful response.

```json
{
  "code": 500,
  "msg": "视频删除失败,请检查权限或视频是否存在",
  "success": false
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[DeleteVideoResp](#schemadeletevideoresp)|

<a id="opIdGetVideoInfo"></a>

## GET Get Video Info

GET /api/v1/video/detail/{videoId}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|videoId|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "video": {
    "id": "string",
    "title": "string",
    "description": "string",
    "playUrl": "string",
    "thumbUrl": "string",
    "createdAt": 0,
    "category": {
      "id": "string",
      "name": "string"
    },
    "tags": [
      {
        "id": "string",
        "name": "string"
      }
    ],
    "author": {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string",
      "slogan": "string",
      "gender": 0,
      "isFollow": true
    }
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[GetVideoInfoResp](#schemagetvideoinforesp)|

<a id="opIdFeed"></a>

## GET Home Video Feed

GET /api/v1/video/feed

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|cursor|query|integer| 否 || 最新视频时间(毫秒时间戳)|
|limit|query|integer| 否 || 请求数量|
|authorID|query|integer| 否 || 作者ID(是否根据用户ID过滤)|
|tag|query|string| 否 || 标签(是否根据标签过滤)|
|categoryId|query|integer| 否 || 分类(是否根据分类过滤)|

> 返回示例

> 200 Response

```json
{
  "next": "string",
  "list": [
    {
      "": {
        "isFavorite": true
      }
    }
  ],
  "isEnd": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[FeedResp](#schemafeedresp)|

<a id="opIdFollowing"></a>

## GET List video of following User

GET /api/v1/video/following

关注用户的视频Feed流

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|pageSize|query|integer| 是 || 分页大小,默认为 10|
|page|query|integer| 是 || 当前页码,默认为 1|
|listAll|query|string| 是 || 是否列出所有,默认为 false|

> 返回示例

> 200 Response

```json
{
  "videoList": [
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "playUrl": "string",
      "thumbUrl": "string",
      "createdAt": 0,
      "category": {
        "id": "string",
        "name": "string"
      },
      "tags": [
        {
          "id": "string",
          "name": "string"
        }
      ],
      "author": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string",
        "slogan": "string",
        "gender": 0,
        "isFollow": true
      }
    }
  ],
  "nextTime": 0,
  "cursorScore": "string",
  "length": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[FollowingResp](#schemafollowingresp)|

<a id="opIdSearch"></a>

## POST Search Video By Keyword

POST /api/v1/video/search

> Body 请求参数

```json
{
  "keyword": "string",
  "offset": 0,
  "limit": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|[SearchReq](#schemasearchreq)| 否 | SearchReq|none|

> 返回示例

> 200 Response

```json
{
  "hits": {},
  "query": "string",
  "processingTimeMs": 0,
  "offset": 0,
  "limit": 0,
  "estimatedTotalHits": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[SearchResp](#schemasearchresp)|

# comment

<a id="opIdDeleteVideoComment"></a>

## DELETE Delete a comment of an Video

DELETE /api/v1/comment/{comment_id}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|comment_id|path|string| 是 ||none|
|type|query|string| 是 ||选择删除一级评论/二级评论(parent/child)|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[DeleteVideoCommentResp](#schemadeletevideocommentresp)|

<a id="opIdEditVideoComment"></a>

## POST Edit comments of an Video

POST /api/v1/comments/{comment_id}

> Body 请求参数

```json
{
  "videoId": 0,
  "commentId": 0,
  "content": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|comment_id|path|string| 是 ||none|
|body|body|[EditVideoCommentReq](#schemaeditvideocommentreq)| 否 | EditVideoCommentReq|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[EditVideoCommentResp](#schemaeditvideocommentresp)|

<a id="opIdListVideoComments"></a>

## GET List comments of an Video

GET /api/v1/video/{video_id}/comments

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|video_id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "success": true,
  "data": {
    "comments": [
      {
        "id": "string",
        "content": "string",
        "atUsers": null,
        "userInfo": {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string",
          "slogan": "string",
          "gender": 0,
          "age": 0
        },
        "showTags": [
          "string"
        ],
        "likedCount": 0,
        "liked": true,
        "createTime": 0,
        "status": 0,
        "videoId": "string",
        "childComments": [
          {
            "id": null,
            "content": null,
            "atUsers": null,
            "userInfo": null,
            "showTags": null,
            "likedCount": null,
            "liked": null,
            "createTime": null,
            "status": null,
            "targetComment": null
          }
        ],
        "childCount": 0,
        "childHasMore": true
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» comments|[object]|true|none||none|
|»»» id|string|false|none||none|
|»»» content|string|false|none||none|
|»»» atUsers|null|false|none||none|
|»»» userInfo|object|false|none||none|
|»»»» userId|string|true|none||none|
|»»»» nickName|string|true|none||none|
|»»»» avatarUrl|string|true|none||none|
|»»»» slogan|string|true|none||none|
|»»»» gender|integer|true|none||none|
|»»»» age|integer|true|none||none|
|»»» showTags|[string]|false|none||none|
|»»» likedCount|integer|false|none||none|
|»»» liked|boolean|false|none||none|
|»»» createTime|integer|false|none||none|
|»»» status|integer|false|none||none|
|»»» videoId|string|false|none||none|
|»»» childComments|[object]|false|none||none|
|»»»» id|string|true|none||none|
|»»»» content|string|true|none||none|
|»»»» atUsers|null|true|none||none|
|»»»» userInfo|object|true|none||none|
|»»»»» userId|string|true|none||none|
|»»»»» nickName|string|true|none||none|
|»»»»» avatarUrl|string|true|none||none|
|»»»»» slogan|string|true|none||none|
|»»»»» gender|integer|true|none||none|
|»»»»» age|integer|true|none||none|
|»»»» showTags|[string]|true|none||none|
|»»»» likedCount|integer|true|none||none|
|»»»» liked|boolean|true|none||none|
|»»»» createTime|integer|true|none||none|
|»»»» status|integer|true|none||none|
|»»»» targetComment|object¦null|true|none||none|
|»»»»» id|string|true|none||none|
|»»»»» userInfo|object|true|none||none|
|»»»»»» userId|string|true|none||none|
|»»»»»» nickName|string|true|none||none|
|»»»»»» avatarUrl|string|true|none||none|
|»»»»»» slogan|string|true|none||none|
|»»»»»» gender|integer|true|none||none|
|»»»»»» age|integer|true|none||none|
|»»» childCount|integer|false|none||none|
|»»» childHasMore|boolean|false|none||none|

<a id="opIdCreateVideoComment"></a>

## POST Create a comment for an Video

POST /api/v1/comment/video

> Body 请求参数

```json
{
  "videoId": 3,
  "content": "这个视频真看好",
  "atUsersId": [
    1,
    2,
    3
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» videoId|body|integer| 是 ||none|
|» content|body|string| 是 ||none|
|» atUsersId|body|[integer]| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "success": true,
  "data": {
    "comment": {
      "id": "string",
      "content": "string",
      "atUsers": null,
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string",
        "slogan": "string",
        "gender": 0,
        "age": 0
      },
      "showTags": null,
      "likedCount": 0,
      "liked": true,
      "createTime": 0,
      "status": 0,
      "videoId": "string",
      "childComments": null,
      "childCount": "string",
      "childHasMore": true
    }
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» success|boolean|true|none||none|
|» data|object|true|none||none|
|»» comment|object|true|none||none|
|»»» id|string|true|none||none|
|»»» content|string|true|none||none|
|»»» atUsers|null|true|none||none|
|»»» userInfo|object|true|none||none|
|»»»» userId|string|true|none||none|
|»»»» nickName|string|true|none||none|
|»»»» avatarUrl|string|true|none||none|
|»»»» slogan|string|true|none||none|
|»»»» gender|integer|true|none||none|
|»»»» age|integer|true|none||none|
|»»» showTags|null|true|none||none|
|»»» likedCount|integer|true|none||none|
|»»» liked|boolean|true|none||none|
|»»» createTime|integer|true|none||none|
|»»» status|integer|true|none||none|
|»»» videoId|string|true|none||none|
|»»» childComments|null|true|none||none|
|»»» childCount|string|true|none||none|
|»»» childHasMore|boolean|true|none||none|

<a id="opIdGetVideoComment"></a>

## GET Get a specific comment of an Video

GET /api/v1/video/{video_id}/comments/{comment_id}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|video_id|path|string| 是 ||none|
|comment_id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "comment": {
    "": {
      "id": "string",
      "content": "string",
      "atUsers": [
        {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      ],
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      },
      "showTags": [
        {
          "id": 0,
          "name": "string"
        }
      ],
      "likedCount": 0,
      "liked": true,
      "createTime": 0,
      "status": 0
    },
    "videoId": "string",
    "childComments": [
      {
        "": {
          "id": "string",
          "content": "string",
          "atUsers": [
            null
          ],
          "userInfo": {},
          "showTags": [
            null
          ],
          "likedCount": 0,
          "liked": true,
          "createTime": 0,
          "status": 0
        },
        "targetComment": {
          "id": "string",
          "userInfo": {}
        }
      }
    ],
    "childCount": "string",
    "childHasMore": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[GetVideoCommentResp](#schemagetvideocommentresp)|

<a id="opIdCreateReplyComment"></a>

## PUT Create a reply for an Comment

PUT /api/v1/video/{video_id}/reply

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|video_id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "comment": {
    "": {
      "id": "string",
      "content": "string",
      "atUsers": [
        {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      ],
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      },
      "showTags": [
        {
          "id": 0,
          "name": "string"
        }
      ],
      "likedCount": 0,
      "liked": true,
      "createTime": 0,
      "status": 0
    },
    "videoId": "string",
    "childComments": [
      {
        "": {
          "id": "string",
          "content": "string",
          "atUsers": [
            null
          ],
          "userInfo": {},
          "showTags": [
            null
          ],
          "likedCount": 0,
          "liked": true,
          "createTime": 0,
          "status": 0
        },
        "targetComment": {
          "id": "string",
          "userInfo": {}
        }
      }
    ],
    "childCount": "string",
    "childHasMore": true
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[CreateReplyCommentResp](#schemacreatereplycommentresp)|

## POST Create a child comment for a parent comment

POST /api/v1/comment/parent

给一级评论增加二级评论

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# favorite

<a id="opIdFavoriteComment"></a>

## PUT Favorite a Comment

PUT /api/v1/favorite/comment/{targetId}

> Body 请求参数

```json
{
  "targetId": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|targetId|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» targetId|body|integer| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "favorite": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[FavoriteCommentResp](#schemafavoritecommentresp)|

<a id="opIdCheckCommentFavorite"></a>

## GET Check Video Favorite Status

GET /api/v1/favorite/video/{targetId}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|targetId|path|string| 是 ||none|
|videoId|query|integer| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "isFavorite": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[CheckCommentFavoriteResp](#schemacheckcommentfavoriteresp)|

## DELETE UnFavorite a Video

DELETE /api/v1/favorite/video/{targetId}

视频取消喜欢

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|targetId|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

<a id="opIdCheckVideoFavorite"></a>

## GET Check comment Favorite Status

GET /api/v1/favorite/comment/{targetId}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|targetId|path|string| 是 ||none|
|videoId|query|integer| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "isFavorite": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[CheckVideoFavoriteResp](#schemacheckvideofavoriteresp)|

## DELETE UnFavorite a Comment

DELETE /api/v1/favorite/comment/{targetId}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|targetId|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# follow

<a id="opIdUnfollow"></a>

## DELETE Unfollow a user

DELETE /api/v1/user/follow_action/{user_id}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[UnFollowResp](#schemaunfollowresp)|

<a id="opIdFollow"></a>

## PUT Follow a user

PUT /api/v1/user/follow_action/{user_id}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[FollowResp](#schemafollowresp)|

<a id="opIdListMyFollowers"></a>

## GET List my followers 

GET /api/v1/user/me/followers

列出当前登录用户的粉丝列表

> 返回示例

> 200 Response

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[ListFollowUserResp](#schemalistfollowuserresp)|

<a id="opIdListMyFollowing"></a>

## GET List my following

GET /api/v1/user/me/following

> 返回示例

> 200 Response

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[ListFollowUserResp](#schemalistfollowuserresp)|

<a id="opIdListFollowers"></a>

## GET List followers list for given user

GET /api/v1/user/{user_id}/followers

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[ListFollowUserResp](#schemalistfollowuserresp)|

# notice

## GET 未命名接口

GET /api/v1/notice/type

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_Response">Response</h2>

<a id="schemaresponse"></a>
<a id="schema_Response"></a>
<a id="tocSresponse"></a>
<a id="tocsresponse"></a>

```json
{
  "code": 0,
  "msg": "string",
  "success": true,
  "data": {}
}

```

响应体

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|code|number|true|none|返回错误码|none|
|msg|string|true|none|响应错误信息|none|
|success|boolean|true|none|响应成功标志|none|
|data|object|true|none|响应数据体|none|

<h2 id="tocS_UserStatisticalInfo">UserStatisticalInfo</h2>

<a id="schemauserstatisticalinfo"></a>
<a id="schema_UserStatisticalInfo"></a>
<a id="tocSuserstatisticalinfo"></a>
<a id="tocsuserstatisticalinfo"></a>

```json
{
  "followingCount": 0,
  "followerCount": 0,
  "likeCount": 0,
  "publishVideoCount": 0,
  "likeVideoCount": 0,
  "collectionsVideoCount": 0
}

```

UserStatisticalInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|followingCount|integer(int32)|true|none||关注数|
|followerCount|integer(int32)|true|none||粉丝数|
|likeCount|integer(int32)|true|none||获赞数量|
|publishVideoCount|integer(int32)|true|none||发布作品数量|
|likeVideoCount|integer(int32)|true|none||点赞作品数量|
|collectionsVideoCount|integer(int32)|true|none||收藏作品数量|

<h2 id="tocS_ossEndpointResponse">ossEndpointResponse</h2>

<a id="schemaossendpointresponse"></a>
<a id="schema_ossEndpointResponse"></a>
<a id="tocSossendpointresponse"></a>
<a id="tocsossendpointresponse"></a>

```json
{
  "endpoint": "string"
}

```

ossEndpointResponse

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|endpoint|string|true|none||none|

<h2 id="tocS_UserInteractionInfo">UserInteractionInfo</h2>

<a id="schemauserinteractioninfo"></a>
<a id="schema_UserInteractionInfo"></a>
<a id="tocSuserinteractioninfo"></a>
<a id="tocsuserinteractioninfo"></a>

```json
{
  "isFollow": true
}

```

UserInteractionInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|isFollow|boolean(boolean)|true|none||是否关注|

<h2 id="tocS_UserHomeDetail">UserHomeDetail</h2>

<a id="schemauserhomedetail"></a>
<a id="schema_UserHomeDetail"></a>
<a id="tocSuserhomedetail"></a>
<a id="tocsuserhomedetail"></a>

```json
{}

```

UserHomeDetail

### 属性

*None*

<h2 id="tocS_UserDetailInfoResp">UserDetailInfoResp</h2>

<a id="schemauserdetailinforesp"></a>
<a id="schema_UserDetailInfoResp"></a>
<a id="tocSuserdetailinforesp"></a>
<a id="tocsuserdetailinforesp"></a>

```json
{
  "": {
    "isFollow": true
  }
}

```

UserDetailInfoResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[UserInteractionInfo](#schemauserinteractioninfo)|false|none||none|

<h2 id="tocS_VideoStatisticalInfo">VideoStatisticalInfo</h2>

<a id="schemavideostatisticalinfo"></a>
<a id="schema_VideoStatisticalInfo"></a>
<a id="tocSvideostatisticalinfo"></a>
<a id="tocsvideostatisticalinfo"></a>

```json
{
  "favoriteCount": 0,
  "commentNum": 0,
  "shareNum": 0
}

```

VideoStatisticalInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|favoriteCount|integer(int64)|true|none||点赞数|
|commentNum|integer(int64)|true|none||评论数|
|shareNum|integer(int64)|true|none||分享数|

<h2 id="tocS_UserDetailInfoReq">UserDetailInfoReq</h2>

<a id="schemauserdetailinforeq"></a>
<a id="schema_UserDetailInfoReq"></a>
<a id="tocSuserdetailinforeq"></a>
<a id="tocsuserdetailinforeq"></a>

```json
{
  "userId": 0
}

```

UserDetailInfoReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|integer(int64)|false|none||none|

<h2 id="tocS_VideoOwnerInfo">VideoOwnerInfo</h2>

<a id="schemavideoownerinfo"></a>
<a id="schema_VideoOwnerInfo"></a>
<a id="tocSvideoownerinfo"></a>
<a id="tocsvideoownerinfo"></a>

```json
{
  "userId": {},
  "nickName": "string",
  "avatarUrl": "string",
  "slogan": "string",
  "gender": 0,
  "isFollow": true
}

```

VideoOwnerInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|
|nickName|string|true|none||用户名|
|avatarUrl|string|true|none||头像|
|slogan|string|true|none||个性签名|
|gender|integer(int64)|true|none||性别|
|isFollow|boolean(boolean)|true|none||是否关注|

<h2 id="tocS_UserDetailInfo">UserDetailInfo</h2>

<a id="schemauserdetailinfo"></a>
<a id="schema_UserDetailInfo"></a>
<a id="tocSuserdetailinfo"></a>
<a id="tocsuserdetailinfo"></a>

```json
{
  "": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "age": 0
  },
  "followingCount": 0,
  "followerCount": 0,
  "likeCount": 0,
  "publishVideoCount": 0,
  "likeVideoCount": 0,
  "collectionsVideoCount": 0
}

```

UserDetailInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[UserBasicInfo](#schemauserbasicinfo)|false|none||none|
|followingCount|integer(int32)|true|none||关注数|
|followerCount|integer(int32)|true|none||粉丝数|
|likeCount|integer(int32)|true|none||获赞数量|
|publishVideoCount|integer(int32)|true|none||发布作品数量|
|likeVideoCount|integer(int32)|true|none||点赞作品数量|
|collectionsVideoCount|integer(int32)|true|none||收藏作品数量|

<h2 id="tocS_VideoManageInfo">VideoManageInfo</h2>

<a id="schemavideomanageinfo"></a>
<a id="schema_VideoManageInfo"></a>
<a id="tocSvideomanageinfo"></a>
<a id="tocsvideomanageinfo"></a>

```json
{
  "publishTime": "string",
  "publishStatus": 0,
  "visibility": 0
}

```

VideoManageInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|publishTime|string|true|none||视频发布时间|
|publishStatus|integer(int32)|true|none||视频发布状态|
|visibility|integer(int32)|true|none||视频可见性|

<h2 id="tocS_UserBasicInfo">UserBasicInfo</h2>

<a id="schemauserbasicinfo"></a>
<a id="schema_UserBasicInfo"></a>
<a id="tocSuserbasicinfo"></a>
<a id="tocsuserbasicinfo"></a>

```json
{
  "userId": "string",
  "nickName": "string",
  "avatarUrl": "string",
  "slogan": "string",
  "gender": 0,
  "age": 0
}

```

UserBasicInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|string|true|none||用户ID|
|nickName|string|true|none||用户名|
|avatarUrl|string|true|none||头像|
|slogan|string|true|none||个性签名|
|gender|integer(int64)|true|none||性别|
|age|integer(int32)|true|none||none|

<h2 id="tocS_VideoBasicInfo">VideoBasicInfo</h2>

<a id="schemavideobasicinfo"></a>
<a id="schema_VideoBasicInfo"></a>
<a id="tocSvideobasicinfo"></a>
<a id="tocsvideobasicinfo"></a>

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "playUrl": "string",
  "thumbUrl": "string",
  "createdAt": 0,
  "category": {
    "id": "string",
    "name": "string"
  },
  "tags": [
    {
      "id": "string",
      "name": "string"
    }
  ],
  "author": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "isFollow": true
  }
}

```

VideoBasicInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||视频ID|
|title|string|true|none||视频标题|
|description|string|true|none||视频描述|
|playUrl|string|true|none||视频播放地址|
|thumbUrl|string|true|none||视频封面地址|
|createdAt|integer(int64)|true|none||视频创建时间(毫秒时间戳)|
|category|[Category](#schemacategory)|true|none||none|
|tags|[[Tag](#schematag)]|true|none||视频标签|
|author|[VideoUserInfo](#schemavideouserinfo)|true|none||none|

<h2 id="tocS_UpdateInfoResp">UpdateInfoResp</h2>

<a id="schemaupdateinforesp"></a>
<a id="schema_UpdateInfoResp"></a>
<a id="tocSupdateinforesp"></a>
<a id="tocsupdateinforesp"></a>

```json
{}

```

UpdateInfoResp

### 属性

*None*

<h2 id="tocS_UpdateInfoReq">UpdateInfoReq</h2>

<a id="schemaupdateinforeq"></a>
<a id="schema_UpdateInfoReq"></a>
<a id="tocSupdateinforeq"></a>
<a id="tocsupdateinforeq"></a>

```json
{
  "nickName": "string",
  "slogan": "string",
  "gender": 0,
  "age": 0,
  "password": "string"
}

```

UpdateInfoReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|nickName|string|false|none||用户名|
|slogan|string|false|none||个性签名|
|gender|integer(int32)|false|none||性别|
|age|integer(int32)|false|none||none|
|password|string|false|none||none|

<h2 id="tocS_UpdatePasswordResp">UpdatePasswordResp</h2>

<a id="schemaupdatepasswordresp"></a>
<a id="schema_UpdatePasswordResp"></a>
<a id="tocSupdatepasswordresp"></a>
<a id="tocsupdatepasswordresp"></a>

```json
{}

```

UpdatePasswordResp

### 属性

*None*

<h2 id="tocS_UpdatePasswordReq">UpdatePasswordReq</h2>

<a id="schemaupdatepasswordreq"></a>
<a id="schema_UpdatePasswordReq"></a>
<a id="tocSupdatepasswordreq"></a>
<a id="tocsupdatepasswordreq"></a>

```json
{}

```

UpdatePasswordReq

### 属性

*None*

<h2 id="tocS_TagRsp">TagRsp</h2>

<a id="schematagrsp"></a>
<a id="schema_TagRsp"></a>
<a id="tocStagrsp"></a>
<a id="tocstagrsp"></a>

```json
{
  "tag_list": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}

```

TagRsp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|tag_list|[[Tag](#schematag)]|true|none||标签列表|

<h2 id="tocS_TagReq">TagReq</h2>

<a id="schematagreq"></a>
<a id="schema_TagReq"></a>
<a id="tocStagreq"></a>
<a id="tocstagreq"></a>

```json
{
  "category": {}
}

```

TagReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|category|object|false|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|

<h2 id="tocS_Tag">Tag</h2>

<a id="schematag"></a>
<a id="schema_Tag"></a>
<a id="tocStag"></a>
<a id="tocstag"></a>

```json
{
  "id": "string",
  "name": "string"
}

```

Tag

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||标签id|
|name|string|true|none||标签名称|

<h2 id="tocS_TargetComment">TargetComment</h2>

<a id="schematargetcomment"></a>
<a id="schema_TargetComment"></a>
<a id="tocStargetcomment"></a>
<a id="tocstargetcomment"></a>

```json
{
  "id": "string",
  "userInfo": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string"
  }
}

```

TargetComment

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||回复的目标评论ID|
|userInfo|[CommentUserInfo](#schemacommentuserinfo)|true|none||none|

<h2 id="tocS_RankingResp">RankingResp</h2>

<a id="schemarankingresp"></a>
<a id="schema_RankingResp"></a>
<a id="tocSrankingresp"></a>
<a id="tocsrankingresp"></a>

```json
{
  "users": [
    {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string",
      "slogan": "string",
      "gender": 0,
      "age": 0
    }
  ]
}

```

RankingResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|users|[[UserBasicInfo](#schemauserbasicinfo)]|true|none||none|

<h2 id="tocS_RankingReq">RankingReq</h2>

<a id="schemarankingreq"></a>
<a id="schema_RankingReq"></a>
<a id="tocSrankingreq"></a>
<a id="tocsrankingreq"></a>

```json
{
  "": {
    "pageSize": "10",
    "page": "1",
    "listAll": "false"
  }
}

```

RankingReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[ListUserOption](#schemalistuseroption)|false|none||none|

<h2 id="tocS_ParentComment">ParentComment</h2>

<a id="schemaparentcomment"></a>
<a id="schema_ParentComment"></a>
<a id="tocSparentcomment"></a>
<a id="tocsparentcomment"></a>

```json
{
  "": {
    "id": "string",
    "content": "string",
    "atUsers": [
      {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      }
    ],
    "userInfo": {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string"
    },
    "showTags": [
      {
        "id": 0,
        "name": "string"
      }
    ],
    "likedCount": 0,
    "liked": true,
    "createTime": 0,
    "status": 0
  },
  "videoId": "string",
  "childComments": [
    {
      "": {
        "id": "string",
        "content": "string",
        "atUsers": [
          {
            "userId": null,
            "nickName": null,
            "avatarUrl": null
          }
        ],
        "userInfo": {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        },
        "showTags": [
          {
            "id": null,
            "name": null
          }
        ],
        "likedCount": 0,
        "liked": true,
        "createTime": 0,
        "status": 0
      },
      "targetComment": {
        "id": "string",
        "userInfo": {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      }
    }
  ],
  "childCount": "string",
  "childHasMore": true
}

```

ParentComment

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[CommentBasicInfo](#schemacommentbasicinfo)|false|none||none|
|videoId|string|true|none||视频ID|
|childComments|[[ChildComment](#schemachildcomment)]|true|none||二级评论列表|
|childCount|string|true|none||二级评论数|
|childHasMore|boolean(boolean)|true|none||是否还有更多二级评论|

<h2 id="tocS_VideoItem">VideoItem</h2>

<a id="schemavideoitem"></a>
<a id="schema_VideoItem"></a>
<a id="tocSvideoitem"></a>
<a id="tocsvideoitem"></a>

```json
{
  "id": 0,
  "title": "string",
  "description": "string",
  "tags": [
    "string"
  ],
  "videoUrl": "string",
  "height": {},
  "width": {},
  "coverUrl": "string",
  "interaction": {
    "isFavorite": true
  },
  "user": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "isFollow": true
  },
  "publishTime": "string"
}

```

VideoItem

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||视频ID|
|title|string|true|none||视频标题|
|description|string|true|none||视频描述|
|tags|[string]|true|none||视频标签|
|videoUrl|string|true|none||视频播放地址|
|height|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|
|width|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|
|coverUrl|string|true|none||视频封面地址|
|interaction|[VideoInteractInfo](#schemavideointeractinfo)|true|none||none|
|user|[VideoUserInfo](#schemavideouserinfo)|true|none||none|
|publishTime|string|true|none||视频创建时间(毫秒时间戳)|

<h2 id="tocS_VideoInteractInfo">VideoInteractInfo</h2>

<a id="schemavideointeractinfo"></a>
<a id="schema_VideoInteractInfo"></a>
<a id="tocSvideointeractinfo"></a>
<a id="tocsvideointeractinfo"></a>

```json
{
  "isFavorite": true
}

```

VideoInteractInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|isFavorite|boolean(boolean)|true|none||当前用户是否已点赞|

<h2 id="tocS_VideoCoverInfo">VideoCoverInfo</h2>

<a id="schemavideocoverinfo"></a>
<a id="schema_VideoCoverInfo"></a>
<a id="tocSvideocoverinfo"></a>
<a id="tocsvideocoverinfo"></a>

```json
{
  "imageScene": "string",
  "url": "string"
}

```

VideoCoverInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|imageScene|string|true|none||视频封面类型|
|url|string|true|none||视频封面地址|

<h2 id="tocS_ListUserOption">ListUserOption</h2>

<a id="schemalistuseroption"></a>
<a id="schema_ListUserOption"></a>
<a id="tocSlistuseroption"></a>
<a id="tocslistuseroption"></a>

```json
{
  "pageSize": "10",
  "page": "1",
  "listAll": "false"
}

```

ListUserOption

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|pageSize|integer(int32)|true|none||分页大小,默认为 10|
|page|integer(int32)|true|none||当前页码,默认为 1|
|listAll|boolean(boolean)|true|none||是否列出所有,默认为 false|

<h2 id="tocS_ListVideoFavoriteResp">ListVideoFavoriteResp</h2>

<a id="schemalistvideofavoriteresp"></a>
<a id="schema_ListVideoFavoriteResp"></a>
<a id="tocSlistvideofavoriteresp"></a>
<a id="tocslistvideofavoriteresp"></a>

```json
{}

```

ListVideoFavoriteResp

### 属性

*None*

<h2 id="tocS_VideoCover">VideoCover</h2>

<a id="schemavideocover"></a>
<a id="schema_VideoCover"></a>
<a id="tocSvideocover"></a>
<a id="tocsvideocover"></a>

```json
{
  "traceId": "string",
  "infoList": [
    {
      "imageScene": "string",
      "url": "string"
    }
  ],
  "fileId": "string",
  "height": 0,
  "width": 0,
  "url": "string"
}

```

VideoCover

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|traceId|string|true|none||none|
|infoList|[[VideoCoverInfo](#schemavideocoverinfo)]|true|none||none|
|fileId|string|true|none||none|
|height|integer(int32)|true|none||none|
|width|integer(int32)|true|none||none|
|url|string|true|none||none|

<h2 id="tocS_ListVideoFavoriteReq">ListVideoFavoriteReq</h2>

<a id="schemalistvideofavoritereq"></a>
<a id="schema_ListVideoFavoriteReq"></a>
<a id="tocSlistvideofavoritereq"></a>
<a id="tocslistvideofavoritereq"></a>

```json
{}

```

ListVideoFavoriteReq

### 属性

*None*

<h2 id="tocS_VideoCard">VideoCard</h2>

<a id="schemavideocard"></a>
<a id="schema_VideoCard"></a>
<a id="tocSvideocard"></a>
<a id="tocsvideocard"></a>

```json
{
  "cover": {
    "traceId": "string",
    "infoList": [
      {
        "imageScene": "string",
        "url": "string"
      }
    ],
    "fileId": "string",
    "height": 0,
    "width": 0,
    "url": "string"
  },
  "type": "string",
  "displayTitle": "string",
  "user": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "isFollow": true
  },
  "interactInfo": {
    "isFavorite": true
  }
}

```

VideoCard

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|cover|[VideoCover](#schemavideocover)|true|none||none|
|type|string|true|none||视频分类|
|displayTitle|string|true|none||视频标题|
|user|[VideoUserInfo](#schemavideouserinfo)|true|none||none|
|interactInfo|[VideoInteractInfo](#schemavideointeractinfo)|true|none||none|

<h2 id="tocS_ListFollowUserResp">ListFollowUserResp</h2>

<a id="schemalistfollowuserresp"></a>
<a id="schema_ListFollowUserResp"></a>
<a id="tocSlistfollowuserresp"></a>
<a id="tocslistfollowuserresp"></a>

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}

```

ListFollowUserResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|users|[[FollowUser](#schemafollowuser)]|true|none||none|

<h2 id="tocS_SearchResp">SearchResp</h2>

<a id="schemasearchresp"></a>
<a id="schema_SearchResp"></a>
<a id="tocSsearchresp"></a>
<a id="tocssearchresp"></a>

```json
{
  "hits": {},
  "query": "string",
  "processingTimeMs": 0,
  "offset": 0,
  "limit": 0,
  "estimatedTotalHits": 0
}

```

SearchResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|hits|object|true|none||搜索结果|
|query|string|true|none||搜索关键字|
|processingTimeMs|integer(int64)|true|none||搜索耗时(毫秒)|
|offset|integer(int64)|true|none||偏移量|
|limit|integer(int64)|true|none||请求数量|
|estimatedTotalHits|integer(int64)|true|none||搜索结果总数|

<h2 id="tocS_ListFollowReq">ListFollowReq</h2>

<a id="schemalistfollowreq"></a>
<a id="schema_ListFollowReq"></a>
<a id="tocSlistfollowreq"></a>
<a id="tocslistfollowreq"></a>

```json
{
  "user_id": 0,
  "": {
    "pageSize": "10",
    "page": "1",
    "listAll": "false"
  }
}

```

ListFollowReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|user_id|integer(int64)|false|none||none|
|*anonymous*|[ListUserOption](#schemalistuseroption)|false|none||none|

<h2 id="tocS_SearchReq">SearchReq</h2>

<a id="schemasearchreq"></a>
<a id="schema_SearchReq"></a>
<a id="tocSsearchreq"></a>
<a id="tocssearchreq"></a>

```json
{
  "keyword": "string",
  "offset": 0,
  "limit": 0
}

```

SearchReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|keyword|string|true|none||搜索关键字|
|offset|integer(int64)|true|none||偏移量|
|limit|integer(int64)|true|none||请求数量|

<h2 id="tocS_FollowingResp">FollowingResp</h2>

<a id="schemafollowingresp"></a>
<a id="schema_FollowingResp"></a>
<a id="tocSfollowingresp"></a>
<a id="tocsfollowingresp"></a>

```json
{
  "videoList": [
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "playUrl": "string",
      "thumbUrl": "string",
      "createdAt": 0,
      "category": {
        "id": "string",
        "name": "string"
      },
      "tags": [
        {
          "id": "string",
          "name": "string"
        }
      ],
      "author": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string",
        "slogan": "string",
        "gender": 0,
        "isFollow": true
      }
    }
  ],
  "nextTime": 0,
  "cursorScore": "string",
  "length": 0
}

```

FollowingResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoList|[[VideoBasicInfo](#schemavideobasicinfo)]|true|none||none|
|nextTime|integer(int64)|true|none||下次请求时间(毫秒时间戳)|
|cursorScore|string|true|none||下次请求时间(毫秒时间戳)|
|length|integer(int32)|true|none||视频列表长度|

<h2 id="tocS_TargetCommnet">TargetCommnet</h2>

<a id="schematargetcommnet"></a>
<a id="schema_TargetCommnet"></a>
<a id="tocStargetcommnet"></a>
<a id="tocstargetcommnet"></a>

```json
{
  "id": "string",
  "userInfo": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string"
  }
}

```

TargetCommnet

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|userInfo|[CommentUserInfo](#schemacommentuserinfo)|true|none||none|

<h2 id="tocS_GetVideoInfoResp">GetVideoInfoResp</h2>

<a id="schemagetvideoinforesp"></a>
<a id="schema_GetVideoInfoResp"></a>
<a id="tocSgetvideoinforesp"></a>
<a id="tocsgetvideoinforesp"></a>

```json
{
  "video": {
    "id": "string",
    "title": "string",
    "description": "string",
    "playUrl": "string",
    "thumbUrl": "string",
    "createdAt": 0,
    "category": {
      "id": "string",
      "name": "string"
    },
    "tags": [
      {
        "id": "string",
        "name": "string"
      }
    ],
    "author": {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string",
      "slogan": "string",
      "gender": 0,
      "isFollow": true
    }
  }
}

```

GetVideoInfoResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|video|[VideoBasicInfo](#schemavideobasicinfo)|true|none||none|

<h2 id="tocS_FollowingReq">FollowingReq</h2>

<a id="schemafollowingreq"></a>
<a id="schema_FollowingReq"></a>
<a id="tocSfollowingreq"></a>
<a id="tocsfollowingreq"></a>

```json
{
  "pageSize": 0,
  "page": 0,
  "listAll": true
}

```

FollowingReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|pageSize|integer(int32)|true|none||分页大小,默认为 10|
|page|integer(int32)|true|none||当前页码,默认为 1|
|listAll|boolean(boolean)|true|none||是否列出所有,默认为 false|

<h2 id="tocS_SubComments">SubComments</h2>

<a id="schemasubcomments"></a>
<a id="schema_SubComments"></a>
<a id="tocSsubcomments"></a>
<a id="tocssubcomments"></a>

```json
{
  "id": "string",
  "videoId": "string",
  "content": "string",
  "createTime": 0,
  "status": 0,
  "atUsers": [
    {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string"
    }
  ],
  "userInfo": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string"
  },
  "commentInteractInfo": {
    "liked": true,
    "likedCount": "string"
  },
  "showTags": [
    "string"
  ]
}

```

SubComments

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|videoId|string|true|none||none|
|content|string|true|none||none|
|createTime|integer(int64)|true|none||none|
|status|integer(int32)|true|none||none|
|atUsers|[[CommentUserInfo](#schemacommentuserinfo)]|true|none||none|
|userInfo|[CommentUserInfo](#schemacommentuserinfo)|true|none||none|
|commentInteractInfo|[CommentInteractInfo](#schemacommentinteractinfo)|true|none||none|
|showTags|[string]|true|none||none|

<h2 id="tocS_GetVideoInfoReq">GetVideoInfoReq</h2>

<a id="schemagetvideoinforeq"></a>
<a id="schema_GetVideoInfoReq"></a>
<a id="tocSgetvideoinforeq"></a>
<a id="tocsgetvideoinforeq"></a>

```json
{
  "videoId": 0
}

```

GetVideoInfoReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|

<h2 id="tocS_ListCommentFavoriteResp">ListCommentFavoriteResp</h2>

<a id="schemalistcommentfavoriteresp"></a>
<a id="schema_ListCommentFavoriteResp"></a>
<a id="tocSlistcommentfavoriteresp"></a>
<a id="tocslistcommentfavoriteresp"></a>

```json
{}

```

ListCommentFavoriteResp

### 属性

*None*

<h2 id="tocS_ListCommentFavoriteReq">ListCommentFavoriteReq</h2>

<a id="schemalistcommentfavoritereq"></a>
<a id="schema_ListCommentFavoriteReq"></a>
<a id="tocSlistcommentfavoritereq"></a>
<a id="tocslistcommentfavoritereq"></a>

```json
{}

```

ListCommentFavoriteReq

### 属性

*None*

<h2 id="tocS_VideoUserInfo">VideoUserInfo</h2>

<a id="schemavideouserinfo"></a>
<a id="schema_VideoUserInfo"></a>
<a id="tocSvideouserinfo"></a>
<a id="tocsvideouserinfo"></a>

```json
{
  "userId": "string",
  "nickName": "string",
  "avatarUrl": "string",
  "slogan": "string",
  "gender": 0,
  "isFollow": true
}

```

VideoUserInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|string|true|none||用户ID|
|nickName|string|true|none||用户名|
|avatarUrl|string|true|none||头像|
|slogan|string|true|none||个性签名|
|gender|integer(int64)|true|none||性别|
|isFollow|boolean(boolean)|true|none||是否关注|

<h2 id="tocS_Video">Video</h2>

<a id="schemavideo"></a>
<a id="schema_Video"></a>
<a id="tocSvideo"></a>
<a id="tocsvideo"></a>

```json
{
  "id": 0,
  "title": "string",
  "playUrl": "string",
  "thumbUrl": "string",
  "favNum": 0,
  "commentNum": 0,
  "shareNum": 0,
  "createdAt": "string",
  "isFav": true,
  "isFollow": true,
  "tags": [
    "string"
  ],
  "author": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string",
    "slogan": "string",
    "gender": 0,
    "isFollow": true
  },
  "description": "string",
  "publishTime": "string",
  "publishStatus": 0,
  "visibility": 0
}

```

Video

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||视频ID|
|title|string|true|none||视频标题|
|playUrl|string|true|none||视频播放地址|
|thumbUrl|string|true|none||视频封面地址|
|favNum|integer(int64)|true|none||点赞数|
|commentNum|integer(int64)|true|none||评论数|
|shareNum|integer(int64)|true|none||分享数|
|createdAt|string|true|none||视频创建时间(毫秒时间戳)|
|isFav|boolean(boolean)|true|none||当前用户是否已点赞|
|isFollow|boolean(boolean)|true|none||当前用户是否已关注该用户|
|tags|[string]|true|none||视频标签|
|author|[VideoUserInfo](#schemavideouserinfo)|true|none||none|
|description|string|true|none||视频描述|
|publishTime|string|true|none||视频发布时间|
|publishStatus|integer(int32)|true|none||视频发布状态|
|visibility|integer(int32)|true|none||视频可见性|

<h2 id="tocS_FeedVideoItem">FeedVideoItem</h2>

<a id="schemafeedvideoitem"></a>
<a id="schema_FeedVideoItem"></a>
<a id="tocSfeedvideoitem"></a>
<a id="tocsfeedvideoitem"></a>

```json
{
  "": {
    "isFavorite": true
  }
}

```

FeedVideoItem

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[VideoInteractInfo](#schemavideointeractinfo)|false|none||none|

<h2 id="tocS_UserInfoResp">UserInfoResp</h2>

<a id="schemauserinforesp"></a>
<a id="schema_UserInfoResp"></a>
<a id="tocSuserinforesp"></a>
<a id="tocsuserinforesp"></a>

```json
{
  "userInfo": {
    "id": 0,
    "phone": "string",
    "nickName": "string",
    "sex": 0,
    "avatarUrl": "string",
    "info": "string"
  }
}

```

UserInfoResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userInfo|[User](#schemauser)|true|none||none|

<h2 id="tocS_ListVideoCommentsResp">ListVideoCommentsResp</h2>

<a id="schemalistvideocommentsresp"></a>
<a id="schema_ListVideoCommentsResp"></a>
<a id="tocSlistvideocommentsresp"></a>
<a id="tocslistvideocommentsresp"></a>

```json
{
  "createdAt": "string",
  "userId": "string",
  "comments": [
    {
      "": {
        "id": "string",
        "content": "string",
        "atUsers": [
          {
            "userId": null,
            "nickName": null,
            "avatarUrl": null
          }
        ],
        "userInfo": {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        },
        "showTags": [
          {
            "id": null,
            "name": null
          }
        ],
        "likedCount": 0,
        "liked": true,
        "createTime": 0,
        "status": 0
      },
      "videoId": "string",
      "childComments": [
        {
          "": {
            "id": null,
            "content": null,
            "atUsers": null,
            "userInfo": null,
            "showTags": null,
            "likedCount": null,
            "liked": null,
            "createTime": null,
            "status": null
          },
          "targetComment": {
            "id": null,
            "userInfo": null
          }
        }
      ],
      "childCount": "string",
      "childHasMore": true
    }
  ]
}

```

ListVideoCommentsResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|createdAt|string|true|none||none|
|userId|string|true|none||none|
|comments|[[ParentComment](#schemaparentcomment)]|true|none||none|

<h2 id="tocS_UserInfoReq">UserInfoReq</h2>

<a id="schemauserinforeq"></a>
<a id="schema_UserInfoReq"></a>
<a id="tocSuserinforeq"></a>
<a id="tocsuserinforeq"></a>

```json
{}

```

UserInfoReq

### 属性

*None*

<h2 id="tocS_ListVideoCommentsReq">ListVideoCommentsReq</h2>

<a id="schemalistvideocommentsreq"></a>
<a id="schema_ListVideoCommentsReq"></a>
<a id="tocSlistvideocommentsreq"></a>
<a id="tocslistvideocommentsreq"></a>

```json
{
  "videoId": 0
}

```

ListVideoCommentsReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|

<h2 id="tocS_User">User</h2>

<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": 0,
  "phone": "string",
  "nickName": "string",
  "sex": 0,
  "avatarUrl": "string",
  "info": "string"
}

```

User

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||none|
|phone|string|true|none||none|
|nickName|string|true|none||none|
|sex|integer(int64)|true|none||none|
|avatarUrl|string|true|none||none|
|info|string|true|none||none|

<h2 id="tocS_FavoriteVideoResp">FavoriteVideoResp</h2>

<a id="schemafavoritevideoresp"></a>
<a id="schema_FavoriteVideoResp"></a>
<a id="tocSfavoritevideoresp"></a>
<a id="tocsfavoritevideoresp"></a>

```json
{
  "favorite": true
}

```

FavoriteVideoResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|favorite|boolean(boolean)|true|none||none|

<h2 id="tocS_UpdateAvatarResp">UpdateAvatarResp</h2>

<a id="schemaupdateavatarresp"></a>
<a id="schema_UpdateAvatarResp"></a>
<a id="tocSupdateavatarresp"></a>
<a id="tocsupdateavatarresp"></a>

```json
{}

```

UpdateAvatarResp

### 属性

*None*

<h2 id="tocS_FavoriteVideoReq">FavoriteVideoReq</h2>

<a id="schemafavoritevideoreq"></a>
<a id="schema_FavoriteVideoReq"></a>
<a id="tocSfavoritevideoreq"></a>
<a id="tocsfavoritevideoreq"></a>

```json
{
  "videoId": 0,
  "isFavorite": true
}

```

FavoriteVideoReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|
|isFavorite|boolean(boolean)|true|none||true: favorite, false: unfavorite|

<h2 id="tocS_UpdateAvatarReq">UpdateAvatarReq</h2>

<a id="schemaupdateavatarreq"></a>
<a id="schema_UpdateAvatarReq"></a>
<a id="tocSupdateavatarreq"></a>
<a id="tocsupdateavatarreq"></a>

```json
{
  "image": "string"
}

```

UpdateAvatarReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|image|string|true|none||none|

<h2 id="tocS_DeleteVideoResp">DeleteVideoResp</h2>

<a id="schemadeletevideoresp"></a>
<a id="schema_DeleteVideoResp"></a>
<a id="tocSdeletevideoresp"></a>
<a id="tocsdeletevideoresp"></a>

```json
{}

```

DeleteVideoResp

### 属性

*None*

<h2 id="tocS_FavoriteCommentResp">FavoriteCommentResp</h2>

<a id="schemafavoritecommentresp"></a>
<a id="schema_FavoriteCommentResp"></a>
<a id="tocSfavoritecommentresp"></a>
<a id="tocsfavoritecommentresp"></a>

```json
{
  "favorite": true
}

```

FavoriteCommentResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|favorite|boolean(boolean)|true|none||none|

<h2 id="tocS_UnFollowResp">UnFollowResp</h2>

<a id="schemaunfollowresp"></a>
<a id="schema_UnFollowResp"></a>
<a id="tocSunfollowresp"></a>
<a id="tocsunfollowresp"></a>

```json
{}

```

UnFollowResp

### 属性

*None*

<h2 id="tocS_DeleteVideoReq">DeleteVideoReq</h2>

<a id="schemadeletevideoreq"></a>
<a id="schema_DeleteVideoReq"></a>
<a id="tocSdeletevideoreq"></a>
<a id="tocsdeletevideoreq"></a>

```json
{}

```

DeleteVideoReq

### 属性

*None*

<h2 id="tocS_FavoriteCommentReq">FavoriteCommentReq</h2>

<a id="schemafavoritecommentreq"></a>
<a id="schema_FavoriteCommentReq"></a>
<a id="tocSfavoritecommentreq"></a>
<a id="tocsfavoritecommentreq"></a>

```json
{
  "videoId": 0,
  "isFavorite": true
}

```

FavoriteCommentReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|
|isFavorite|boolean(boolean)|true|none||true: favorite, false: unfavorite|

<h2 id="tocS_UnFollowReq">UnFollowReq</h2>

<a id="schemaunfollowreq"></a>
<a id="schema_UnFollowReq"></a>
<a id="tocSunfollowreq"></a>
<a id="tocsunfollowreq"></a>

```json
{
  "user_id": 0
}

```

UnFollowReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|user_id|integer(int64)|true|none||none|

<h2 id="tocS_RegisterResp">RegisterResp</h2>

<a id="schemaregisterresp"></a>
<a id="schema_RegisterResp"></a>
<a id="tocSregisterresp"></a>
<a id="tocsregisterresp"></a>

```json
{
  "accessToken": "string",
  "accessExpire": 0,
  "refreshAfter": 0
}

```

RegisterResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|accessToken|string|true|none||none|
|accessExpire|integer(int64)|true|none||none|
|refreshAfter|integer(int64)|true|none||none|

<h2 id="tocS_RegisterReq">RegisterReq</h2>

<a id="schemaregisterreq"></a>
<a id="schema_RegisterReq"></a>
<a id="tocSregisterreq"></a>
<a id="tocsregisterreq"></a>

```json
{
  "phone": "string",
  "password": "string",
  "nickName": "string"
}

```

RegisterReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|phone|string|true|none||none|
|password|string|true|none||none|
|nickName|string|true|none||none|

<h2 id="tocS_CreateVideoResp">CreateVideoResp</h2>

<a id="schemacreatevideoresp"></a>
<a id="schema_CreateVideoResp"></a>
<a id="tocScreatevideoresp"></a>
<a id="tocscreatevideoresp"></a>

```json
{}

```

CreateVideoResp

### 属性

*None*

<h2 id="tocS_LoginResp">LoginResp</h2>

<a id="schemaloginresp"></a>
<a id="schema_LoginResp"></a>
<a id="tocSloginresp"></a>
<a id="tocsloginresp"></a>

```json
{
  "accessToken": "string",
  "accessExpire": 0,
  "refreshAfter": 0
}

```

LoginResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|accessToken|string|true|none||none|
|accessExpire|integer(int64)|true|none||none|
|refreshAfter|integer(int64)|true|none||none|

<h2 id="tocS_CreateVideoReq">CreateVideoReq</h2>

<a id="schemacreatevideoreq"></a>
<a id="schema_CreateVideoReq"></a>
<a id="tocScreatevideoreq"></a>
<a id="tocscreatevideoreq"></a>

```json
{
  "title": "string",
  "playUrl": "string",
  "thumbUrl": "string",
  "description": "string",
  "category": 0,
  "tags": [
    "string"
  ],
  "publishTime": 0,
  "videoKey": "string",
  "visibility": 0
}

```

CreateVideoReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|title|string|true|none||视频标题|
|playUrl|string|true|none||视频播放地址|
|thumbUrl|string|true|none||视频封面地址|
|description|string|true|none||视频描述|
|category|integer(int64)|true|none||视频分类|
|tags|[string]|true|none||视频标签|
|publishTime|integer(int64)|false|none||视频发布时间(毫秒时间戳)|
|videoKey|string|false|none||视频上传key|
|visibility|integer(int32)|true|none||视频可见性(1:公开,2:私密)|

<h2 id="tocS_LoginReq">LoginReq</h2>

<a id="schemaloginreq"></a>
<a id="schema_LoginReq"></a>
<a id="tocSloginreq"></a>
<a id="tocsloginreq"></a>

```json
{
  "phone": "string",
  "password": "string"
}

```

LoginReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|phone|string|true|none||none|
|password|string|true|none||none|

<h2 id="tocS_GetVideoCommentResp">GetVideoCommentResp</h2>

<a id="schemagetvideocommentresp"></a>
<a id="schema_GetVideoCommentResp"></a>
<a id="tocSgetvideocommentresp"></a>
<a id="tocsgetvideocommentresp"></a>

```json
{
  "comment": {
    "": {
      "id": "string",
      "content": "string",
      "atUsers": [
        {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      ],
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      },
      "showTags": [
        {
          "id": 0,
          "name": "string"
        }
      ],
      "likedCount": 0,
      "liked": true,
      "createTime": 0,
      "status": 0
    },
    "videoId": "string",
    "childComments": [
      {
        "": {
          "id": "string",
          "content": "string",
          "atUsers": [
            null
          ],
          "userInfo": {},
          "showTags": [
            null
          ],
          "likedCount": 0,
          "liked": true,
          "createTime": 0,
          "status": 0
        },
        "targetComment": {
          "id": "string",
          "userInfo": {}
        }
      }
    ],
    "childCount": "string",
    "childHasMore": true
  }
}

```

GetVideoCommentResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|comment|[ParentComment](#schemaparentcomment)|true|none||none|

<h2 id="tocS_ListMyFollowingResp">ListMyFollowingResp</h2>

<a id="schemalistmyfollowingresp"></a>
<a id="schema_ListMyFollowingResp"></a>
<a id="tocSlistmyfollowingresp"></a>
<a id="tocslistmyfollowingresp"></a>

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}

```

ListMyFollowingResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|users|[[FollowUser](#schemafollowuser)]|true|none||none|

<h2 id="tocS_GetVideoCommentReq">GetVideoCommentReq</h2>

<a id="schemagetvideocommentreq"></a>
<a id="schema_GetVideoCommentReq"></a>
<a id="tocSgetvideocommentreq"></a>
<a id="tocsgetvideocommentreq"></a>

```json
{
  "commentId": 0
}

```

GetVideoCommentReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|commentId|integer(int64)|true|none||none|

<h2 id="tocS_ListMyFollowingReq">ListMyFollowingReq</h2>

<a id="schemalistmyfollowingreq"></a>
<a id="schema_ListMyFollowingReq"></a>
<a id="tocSlistmyfollowingreq"></a>
<a id="tocslistmyfollowingreq"></a>

```json
{
  "": {
    "pageSize": "10",
    "page": "1",
    "listAll": "false"
  }
}

```

ListMyFollowingReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[ListUserOption](#schemalistuseroption)|false|none||none|

<h2 id="tocS_CurrentUserInfoResp">CurrentUserInfoResp</h2>

<a id="schemacurrentuserinforesp"></a>
<a id="schema_CurrentUserInfoResp"></a>
<a id="tocScurrentuserinforesp"></a>
<a id="tocscurrentuserinforesp"></a>

```json
{
  "": {
    "isFollow": true
  }
}

```

CurrentUserInfoResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[UserInteractionInfo](#schemauserinteractioninfo)|false|none||none|

<h2 id="tocS_ListMyFollowersResp">ListMyFollowersResp</h2>

<a id="schemalistmyfollowersresp"></a>
<a id="schema_ListMyFollowersResp"></a>
<a id="tocSlistmyfollowersresp"></a>
<a id="tocslistmyfollowersresp"></a>

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}

```

ListMyFollowersResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|users|[[FollowUser](#schemafollowuser)]|true|none||none|

<h2 id="tocS_CurrentUserInfoReq">CurrentUserInfoReq</h2>

<a id="schemacurrentuserinforeq"></a>
<a id="schema_CurrentUserInfoReq"></a>
<a id="tocScurrentuserinforeq"></a>
<a id="tocscurrentuserinforeq"></a>

```json
{}

```

CurrentUserInfoReq

### 属性

*None*

<h2 id="tocS_ListMyFollowersReq">ListMyFollowersReq</h2>

<a id="schemalistmyfollowersreq"></a>
<a id="schema_ListMyFollowersReq"></a>
<a id="tocSlistmyfollowersreq"></a>
<a id="tocslistmyfollowersreq"></a>

```json
{
  "": {
    "pageSize": "10",
    "page": "1",
    "listAll": "false"
  }
}

```

ListMyFollowersReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[ListUserOption](#schemalistuseroption)|false|none||none|

<h2 id="tocS_ListFollowingResp">ListFollowingResp</h2>

<a id="schemalistfollowingresp"></a>
<a id="schema_ListFollowingResp"></a>
<a id="tocSlistfollowingresp"></a>
<a id="tocslistfollowingresp"></a>

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}

```

ListFollowingResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|users|[[FollowUser](#schemafollowuser)]|true|none||none|

<h2 id="tocS_ListFollowingReq">ListFollowingReq</h2>

<a id="schemalistfollowingreq"></a>
<a id="schema_ListFollowingReq"></a>
<a id="tocSlistfollowingreq"></a>
<a id="tocslistfollowingreq"></a>

```json
{
  "user_id": 0,
  "": {
    "pageSize": "10",
    "page": "1",
    "listAll": "false"
  }
}

```

ListFollowingReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|user_id|integer(int64)|true|none||none|
|*anonymous*|[ListUserOption](#schemalistuseroption)|false|none||none|

<h2 id="tocS_CreateUpTokenResp">CreateUpTokenResp</h2>

<a id="schemacreateuptokenresp"></a>
<a id="schema_CreateUpTokenResp"></a>
<a id="tocScreateuptokenresp"></a>
<a id="tocscreateuptokenresp"></a>

```json
{
  "upToken": "string",
  "expires": 0
}

```

CreateUpTokenResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|upToken|string|true|none||上传凭证|
|expires|integer(int64)|true|none||上传凭证过期时间(秒)|

<h2 id="tocS_ListFollowersResp">ListFollowersResp</h2>

<a id="schemalistfollowersresp"></a>
<a id="schema_ListFollowersResp"></a>
<a id="tocSlistfollowersresp"></a>
<a id="tocslistfollowersresp"></a>

```json
{
  "users": [
    {
      "": {
        "isFollow": true
      }
    }
  ]
}

```

ListFollowersResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|users|[[FollowUser](#schemafollowuser)]|true|none||none|

<h2 id="tocS_CreateUpTokenReq">CreateUpTokenReq</h2>

<a id="schemacreateuptokenreq"></a>
<a id="schema_CreateUpTokenReq"></a>
<a id="tocScreateuptokenreq"></a>
<a id="tocscreateuptokenreq"></a>

```json
{
  "uploadType": "string"
}

```

CreateUpTokenReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|uploadType|string|true|none||上传类型(video:视频,cover:封面,avatar:头像)|

<h2 id="tocS_EditVideoCommentResp">EditVideoCommentResp</h2>

<a id="schemaeditvideocommentresp"></a>
<a id="schema_EditVideoCommentResp"></a>
<a id="tocSeditvideocommentresp"></a>
<a id="tocseditvideocommentresp"></a>

```json
{}

```

EditVideoCommentResp

### 属性

*None*

<h2 id="tocS_ListFollowersReq">ListFollowersReq</h2>

<a id="schemalistfollowersreq"></a>
<a id="schema_ListFollowersReq"></a>
<a id="tocSlistfollowersreq"></a>
<a id="tocslistfollowersreq"></a>

```json
{
  "user_id": 0,
  "": {
    "pageSize": "10",
    "page": "1",
    "listAll": "false"
  }
}

```

ListFollowersReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|user_id|integer(int64)|true|none||none|
|*anonymous*|[ListUserOption](#schemalistuseroption)|false|none||none|

<h2 id="tocS_CreateReplyCommentResp">CreateReplyCommentResp</h2>

<a id="schemacreatereplycommentresp"></a>
<a id="schema_CreateReplyCommentResp"></a>
<a id="tocScreatereplycommentresp"></a>
<a id="tocscreatereplycommentresp"></a>

```json
{
  "comment": {
    "": {
      "id": "string",
      "content": "string",
      "atUsers": [
        {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      ],
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      },
      "showTags": [
        {
          "id": 0,
          "name": "string"
        }
      ],
      "likedCount": 0,
      "liked": true,
      "createTime": 0,
      "status": 0
    },
    "videoId": "string",
    "childComments": [
      {
        "": {
          "id": "string",
          "content": "string",
          "atUsers": [
            null
          ],
          "userInfo": {},
          "showTags": [
            null
          ],
          "likedCount": 0,
          "liked": true,
          "createTime": 0,
          "status": 0
        },
        "targetComment": {
          "id": "string",
          "userInfo": {}
        }
      }
    ],
    "childCount": "string",
    "childHasMore": true
  }
}

```

CreateReplyCommentResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|comment|[ParentComment](#schemaparentcomment)|true|none||none|

<h2 id="tocS_EditVideoCommentReq">EditVideoCommentReq</h2>

<a id="schemaeditvideocommentreq"></a>
<a id="schema_EditVideoCommentReq"></a>
<a id="tocSeditvideocommentreq"></a>
<a id="tocseditvideocommentreq"></a>

```json
{
  "videoId": 0,
  "commentId": 0,
  "content": "string"
}

```

EditVideoCommentReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|
|commentId|integer(int64)|true|none||none|
|content|string|true|none||none|

<h2 id="tocS_FollowUser">FollowUser</h2>

<a id="schemafollowuser"></a>
<a id="schema_FollowUser"></a>
<a id="tocSfollowuser"></a>
<a id="tocsfollowuser"></a>

```json
{
  "": {
    "isFollow": true
  }
}

```

FollowUser

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[UserInteractionInfo](#schemauserinteractioninfo)|false|none||none|

<h2 id="tocS_CreateReplyCommentReq">CreateReplyCommentReq</h2>

<a id="schemacreatereplycommentreq"></a>
<a id="schema_CreateReplyCommentReq"></a>
<a id="tocScreatereplycommentreq"></a>
<a id="tocscreatereplycommentreq"></a>

```json
{
  "content": "string",
  "videoId": 0,
  "parentId": 0,
  "targetId": 0
}

```

CreateReplyCommentReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|content|string|true|none||none|
|videoId|integer(int64)|true|none||none|
|parentId|integer(int64)|true|none||none|
|targetId|integer(int64)|true|none||none|

<h2 id="tocS_DeleteVideoCommentResp">DeleteVideoCommentResp</h2>

<a id="schemadeletevideocommentresp"></a>
<a id="schema_DeleteVideoCommentResp"></a>
<a id="tocSdeletevideocommentresp"></a>
<a id="tocsdeletevideocommentresp"></a>

```json
{}

```

DeleteVideoCommentResp

### 属性

*None*

<h2 id="tocS_FollowResp">FollowResp</h2>

<a id="schemafollowresp"></a>
<a id="schema_FollowResp"></a>
<a id="tocSfollowresp"></a>
<a id="tocsfollowresp"></a>

```json
{}

```

FollowResp

### 属性

*None*

<h2 id="tocS_DeleteVideoCommentReq">DeleteVideoCommentReq</h2>

<a id="schemadeletevideocommentreq"></a>
<a id="schema_DeleteVideoCommentReq"></a>
<a id="tocSdeletevideocommentreq"></a>
<a id="tocsdeletevideocommentreq"></a>

```json
{
  "videoId": 0,
  "commentId": 0
}

```

DeleteVideoCommentReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|
|commentId|integer(int64)|true|none||none|

<h2 id="tocS_FollowReq">FollowReq</h2>

<a id="schemafollowreq"></a>
<a id="schema_FollowReq"></a>
<a id="tocSfollowreq"></a>
<a id="tocsfollowreq"></a>

```json
{
  "user_id": 0
}

```

FollowReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|user_id|integer(int64)|true|none||none|

<h2 id="tocS_CheckVideoFavoriteResp">CheckVideoFavoriteResp</h2>

<a id="schemacheckvideofavoriteresp"></a>
<a id="schema_CheckVideoFavoriteResp"></a>
<a id="tocScheckvideofavoriteresp"></a>
<a id="tocscheckvideofavoriteresp"></a>

```json
{
  "isFavorite": true
}

```

CheckVideoFavoriteResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|isFavorite|boolean(boolean)|true|none||none|

<h2 id="tocS_CreateVideoCommentResp">CreateVideoCommentResp</h2>

<a id="schemacreatevideocommentresp"></a>
<a id="schema_CreateVideoCommentResp"></a>
<a id="tocScreatevideocommentresp"></a>
<a id="tocscreatevideocommentresp"></a>

```json
{
  "comment": {
    "": {
      "id": "string",
      "content": "string",
      "atUsers": [
        {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      ],
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      },
      "showTags": [
        {
          "id": 0,
          "name": "string"
        }
      ],
      "likedCount": 0,
      "liked": true,
      "createTime": 0,
      "status": 0
    },
    "videoId": "string",
    "childComments": [
      {
        "": {
          "id": "string",
          "content": "string",
          "atUsers": [
            null
          ],
          "userInfo": {},
          "showTags": [
            null
          ],
          "likedCount": 0,
          "liked": true,
          "createTime": 0,
          "status": 0
        },
        "targetComment": {
          "id": "string",
          "userInfo": {}
        }
      }
    ],
    "childCount": "string",
    "childHasMore": true
  }
}

```

CreateVideoCommentResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|comment|[ParentComment](#schemaparentcomment)|true|none||none|

<h2 id="tocS_FeedResp">FeedResp</h2>

<a id="schemafeedresp"></a>
<a id="schema_FeedResp"></a>
<a id="tocSfeedresp"></a>
<a id="tocsfeedresp"></a>

```json
{
  "next": "string",
  "list": [
    {
      "": {
        "isFavorite": true
      }
    }
  ],
  "isEnd": true
}

```

FeedResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|next|string|true|none||请求游标|
|list|[[FeedVideoItem](#schemafeedvideoitem)]|true|none||视频列表|
|isEnd|boolean(boolean)|true|none||是否已到最后一页|

<h2 id="tocS_CommonTag">CommonTag</h2>

<a id="schemacommontag"></a>
<a id="schema_CommonTag"></a>
<a id="tocScommontag"></a>
<a id="tocscommontag"></a>

```json
{
  "id": 0,
  "name": "string"
}

```

CommonTag

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||none|
|name|string|true|none||none|

<h2 id="tocS_CheckVideoFavoriteReq">CheckVideoFavoriteReq</h2>

<a id="schemacheckvideofavoritereq"></a>
<a id="schema_CheckVideoFavoriteReq"></a>
<a id="tocScheckvideofavoritereq"></a>
<a id="tocscheckvideofavoritereq"></a>

```json
{
  "videoId": 0
}

```

CheckVideoFavoriteReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|

<h2 id="tocS_Commnent">Commnent</h2>

<a id="schemacommnent"></a>
<a id="schema_Commnent"></a>
<a id="tocScommnent"></a>
<a id="tocscommnent"></a>

```json
{
  "id": "string",
  "videoId": "string",
  "status": 0,
  "content": "string",
  "createTime": 0,
  "userInfo": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string"
  },
  "commentInteractInfo": {
    "liked": true,
    "likedCount": "string"
  },
  "subComments": [
    {
      "id": "string",
      "videoId": "string",
      "content": "string",
      "createTime": 0,
      "status": 0,
      "atUsers": [
        {
          "userId": "string",
          "nickName": "string",
          "avatarUrl": "string"
        }
      ],
      "userInfo": {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      },
      "commentInteractInfo": {
        "liked": true,
        "likedCount": "string"
      },
      "showTags": [
        "string"
      ]
    }
  ],
  "subCommentCount": "string",
  "subCommentHasMore": true,
  "subCommentCursor": "string",
  "atUsers": [
    {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string"
    }
  ],
  "showTags": [
    "string"
  ],
  "ipAddress": "string"
}

```

Commnent

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|videoId|string|true|none||none|
|status|integer(int32)|true|none||none|
|content|string|true|none||none|
|createTime|integer(int64)|true|none||none|
|userInfo|[CommentUserInfo](#schemacommentuserinfo)|true|none||none|
|commentInteractInfo|[CommentInteractInfo](#schemacommentinteractinfo)|true|none||none|
|subComments|[[SubComments](#schemasubcomments)]|true|none||none|
|subCommentCount|string|true|none||none|
|subCommentHasMore|boolean(boolean)|true|none||none|
|subCommentCursor|string|true|none||none|
|atUsers|[[CommentUserInfo](#schemacommentuserinfo)]|true|none||none|
|showTags|[string]|true|none||none|
|ipAddress|string|true|none||none|

<h2 id="tocS_CreateVideoCommentReq">CreateVideoCommentReq</h2>

<a id="schemacreatevideocommentreq"></a>
<a id="schema_CreateVideoCommentReq"></a>
<a id="tocScreatevideocommentreq"></a>
<a id="tocscreatevideocommentreq"></a>

```json
{
  "content": "string",
  "videoId": 0
}

```

CreateVideoCommentReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|content|string|true|none||none|
|videoId|integer(int64)|true|none||none|

<h2 id="tocS_FeedReq">FeedReq</h2>

<a id="schemafeedreq"></a>
<a id="schema_FeedReq"></a>
<a id="tocSfeedreq"></a>
<a id="tocsfeedreq"></a>

```json
{
  "cursor": "0",
  "limit": "10",
  "authorID": "0",
  "tag": "string",
  "categoryId": 0
}

```

FeedReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|cursor|integer(int64)|true|none||最新视频时间(毫秒时间戳)|
|limit|integer(int32)|true|none||请求数量|
|authorID|integer(int64)|true|none||作者ID(是否根据用户ID过滤)|
|tag|string|false|none||标签(是否根据标签过滤)|
|categoryId|integer(int64)|false|none||分类(是否根据分类过滤)|

<h2 id="tocS_CommentUserInfo">CommentUserInfo</h2>

<a id="schemacommentuserinfo"></a>
<a id="schema_CommentUserInfo"></a>
<a id="tocScommentuserinfo"></a>
<a id="tocscommentuserinfo"></a>

```json
{
  "userId": "string",
  "nickName": "string",
  "avatarUrl": "string"
}

```

CommentUserInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|string|true|none||用户ID|
|nickName|string|true|none||用户名|
|avatarUrl|string|true|none||头像|

<h2 id="tocS_CountFollowResp">CountFollowResp</h2>

<a id="schemacountfollowresp"></a>
<a id="schema_CountFollowResp"></a>
<a id="tocScountfollowresp"></a>
<a id="tocscountfollowresp"></a>

```json
{
  "followingCount": 0,
  "followerCount": 0
}

```

CountFollowResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|followingCount|integer(int64)|true|none||none|
|followerCount|integer(int64)|true|none||none|

<h2 id="tocS_CommentBasicInfo">CommentBasicInfo</h2>

<a id="schemacommentbasicinfo"></a>
<a id="schema_CommentBasicInfo"></a>
<a id="tocScommentbasicinfo"></a>
<a id="tocscommentbasicinfo"></a>

```json
{
  "id": "string",
  "content": "string",
  "atUsers": [
    {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string"
    }
  ],
  "userInfo": {
    "userId": "string",
    "nickName": "string",
    "avatarUrl": "string"
  },
  "showTags": [
    {
      "id": 0,
      "name": "string"
    }
  ],
  "likedCount": 0,
  "liked": true,
  "createTime": 0,
  "status": 0
}

```

CommentBasicInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||评论ID|
|content|string|true|none||评论内容|
|atUsers|[[CommentUserInfo](#schemacommentuserinfo)]|true|none||@用户列表(暂未实现)|
|userInfo|[CommentUserInfo](#schemacommentuserinfo)|true|none||none|
|showTags|[[CommonTag](#schemacommontag)]|true|none||标签列表(暂未实现)|
|likedCount|integer(int64)|true|none||点赞数|
|liked|boolean(boolean)|true|none||当前用户是否已点赞|
|createTime|integer(int64)|true|none||创建时间|
|status|integer(int32)|true|none||none|

<h2 id="tocS_CommentInteractInfo">CommentInteractInfo</h2>

<a id="schemacommentinteractinfo"></a>
<a id="schema_CommentInteractInfo"></a>
<a id="tocScommentinteractinfo"></a>
<a id="tocscommentinteractinfo"></a>

```json
{
  "liked": true,
  "likedCount": "string"
}

```

CommentInteractInfo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|liked|boolean(boolean)|true|none||当前用户是否已点赞|
|likedCount|string|true|none||点赞数|

<h2 id="tocS_CountFollowReq">CountFollowReq</h2>

<a id="schemacountfollowreq"></a>
<a id="schema_CountFollowReq"></a>
<a id="tocScountfollowreq"></a>
<a id="tocscountfollowreq"></a>

```json
{
  "userId": {}
}

```

CountFollowReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|

<h2 id="tocS_ChildComment">ChildComment</h2>

<a id="schemachildcomment"></a>
<a id="schema_ChildComment"></a>
<a id="tocSchildcomment"></a>
<a id="tocschildcomment"></a>

```json
{
  "": {
    "id": "string",
    "content": "string",
    "atUsers": [
      {
        "userId": "string",
        "nickName": "string",
        "avatarUrl": "string"
      }
    ],
    "userInfo": {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string"
    },
    "showTags": [
      {
        "id": 0,
        "name": "string"
      }
    ],
    "likedCount": 0,
    "liked": true,
    "createTime": 0,
    "status": 0
  },
  "targetComment": {
    "id": "string",
    "userInfo": {
      "userId": "string",
      "nickName": "string",
      "avatarUrl": "string"
    }
  }
}

```

ChildComment

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|[CommentBasicInfo](#schemacommentbasicinfo)|false|none||none|
|targetComment|[TargetComment](#schematargetcomment)|true|none||none|

<h2 id="tocS_CheckMyFollowingResp">CheckMyFollowingResp</h2>

<a id="schemacheckmyfollowingresp"></a>
<a id="schema_CheckMyFollowingResp"></a>
<a id="tocScheckmyfollowingresp"></a>
<a id="tocscheckmyfollowingresp"></a>

```json
{
  "status": true
}

```

CheckMyFollowingResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|boolean(boolean)|true|none||none|

<h2 id="tocS_CheckMyFollowingReq">CheckMyFollowingReq</h2>

<a id="schemacheckmyfollowingreq"></a>
<a id="schema_CheckMyFollowingReq"></a>
<a id="tocScheckmyfollowingreq"></a>
<a id="tocscheckmyfollowingreq"></a>

```json
{
  "userId": {}
}

```

CheckMyFollowingReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|userId|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|

<h2 id="tocS_ChangePasswordResp">ChangePasswordResp</h2>

<a id="schemachangepasswordresp"></a>
<a id="schema_ChangePasswordResp"></a>
<a id="tocSchangepasswordresp"></a>
<a id="tocschangepasswordresp"></a>

```json
{}

```

ChangePasswordResp

### 属性

*None*

<h2 id="tocS_CheckCommentFavoriteResp">CheckCommentFavoriteResp</h2>

<a id="schemacheckcommentfavoriteresp"></a>
<a id="schema_CheckCommentFavoriteResp"></a>
<a id="tocScheckcommentfavoriteresp"></a>
<a id="tocscheckcommentfavoriteresp"></a>

```json
{
  "isFavorite": true
}

```

CheckCommentFavoriteResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|isFavorite|boolean(boolean)|true|none||none|

<h2 id="tocS_CheckFollowingResp">CheckFollowingResp</h2>

<a id="schemacheckfollowingresp"></a>
<a id="schema_CheckFollowingResp"></a>
<a id="tocScheckfollowingresp"></a>
<a id="tocscheckfollowingresp"></a>

```json
{
  "status": true
}

```

CheckFollowingResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|status|boolean(boolean)|true|none||none|

<h2 id="tocS_ChangePasswordReq">ChangePasswordReq</h2>

<a id="schemachangepasswordreq"></a>
<a id="schema_ChangePasswordReq"></a>
<a id="tocSchangepasswordreq"></a>
<a id="tocschangepasswordreq"></a>

```json
{}

```

ChangePasswordReq

### 属性

*None*

<h2 id="tocS_CheckCommentFavoriteReq">CheckCommentFavoriteReq</h2>

<a id="schemacheckcommentfavoritereq"></a>
<a id="schema_CheckCommentFavoriteReq"></a>
<a id="tocScheckcommentfavoritereq"></a>
<a id="tocscheckcommentfavoritereq"></a>

```json
{
  "videoId": 0
}

```

CheckCommentFavoriteReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|videoId|integer(int64)|true|none||none|

<h2 id="tocS_CheckFollowingReq">CheckFollowingReq</h2>

<a id="schemacheckfollowingreq"></a>
<a id="schema_CheckFollowingReq"></a>
<a id="tocScheckfollowingreq"></a>
<a id="tocscheckfollowingreq"></a>

```json
{
  "doerUserId": {},
  "contextUserId": {}
}

```

CheckFollowingReq

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|doerUserId|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|
|contextUserId|object|true|none||[后端统一返回结果]<[Implementation of Excel T() function<br /><p/><br />If the argument is a text or error value it is returned unmodified.  All other argument types<br />cause an empty string result.  If the argument is an area, the first (top-left) cell is used<br />(regardless of the coordinates of the evaluating formula cell).]>|

<h2 id="tocS_CategoryResp">CategoryResp</h2>

<a id="schemacategoryresp"></a>
<a id="schema_CategoryResp"></a>
<a id="tocScategoryresp"></a>
<a id="tocscategoryresp"></a>

```json
{
  "categoryList": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}

```

CategoryResp

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|categoryList|[[Category](#schemacategory)]|true|none||none|

<h2 id="tocS_CategoryReq">CategoryReq</h2>

<a id="schemacategoryreq"></a>
<a id="schema_CategoryReq"></a>
<a id="tocScategoryreq"></a>
<a id="tocscategoryreq"></a>

```json
{}

```

CategoryReq

### 属性

*None*

<h2 id="tocS_Category">Category</h2>

<a id="schemacategory"></a>
<a id="schema_Category"></a>
<a id="tocScategory"></a>
<a id="tocscategory"></a>

```json
{
  "id": "string",
  "name": "string"
}

```

Category

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||分类ID|
|name|string|true|none||分类名称|

