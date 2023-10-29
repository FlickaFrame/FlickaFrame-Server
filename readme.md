# 网页短视频应用

使用七牛云存储、七牛视频相关产品（如视频截帧等）开发一款Web端短视频应用

## 基础功能（必须实现）

- 视频播放：播放、暂停、进度条拖拽
- 内容分类：视频内容分类页，如热门视频、体育频道
- ­视频切换：可通过上下键翻看视频

## 高级功能（可选实现）

- 账户系统：用户可登录，收藏视频
- 可参考常见短视频应用自由增加功能，提升完善度，如点赞、分享、关注、搜索等

## Deploy
- 数据库：Mysql
- 缓存：Redis
- 搜索引擎：Meilisearch
- 对象存储：七牛云存储
- 数据库和索引同步：Meilisync

### MYSQL

```bash
    docker run -itd \
    --restart=always \ 
    --name mysql-qiniuyun \
    -v $(pwd)/data/mysql:/var/lib/mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=$(mysql_root_password) \
    mysql
```

### REDIS

```bash
    docker run -itd \
    --restart=always \ 
    --name redis-qiniuyun \
    -v $(pwd)/data/redis:/data \
    -p 6379:6379 \
    redis
```

### meilisearch

```shell
    docker run -itd \
    --restart=always \ 
    -p 7700:7700 \
    -e MEILI_MASTER_KEY=$(master_key) \
    -v $(pwd)/data/meili_data:/meili_data \
    getmeili/meilisearch
```

### meilisearch-syn

```shell
    docker run -itd \
    --restart=always \ 
    -v ./etc/meilisync.yml:/meilisync/config.yml \
    long2ice/meilisync
```

## API Doc

```shell
# 生成openAPI文档, 生成的文档在docs/swagger
# 可以使用ApiFox订阅 http://localhost:8080/api/v1/swagger
make gen-api-swagger 
# 生成api markdown文档, 生成的文档在docs/api
make gen-api-doc
```
