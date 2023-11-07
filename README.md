
#  FlickaFrame-Server | 2023 年 **七牛云校园马拉松**

## 快速启动

1. clone项目到本地

    ```shell
        git clone git@github.com:FlickaFrame/FlickaFrame-Server.git
        cd TikTok
    ```
2. 安装依赖(参考下面中间件介绍和docker-compose.yml)

3. 修改配置信息

    修改`etc/main.yaml`中的配置信息

4. 启动项目

    ```shell
        make run
    ```

## 二次开发 & 贡献

1. 安装golang依赖
    ```shell
        make install
    ```
2. 代码生成
    ```shell
        make gen
    ```
3. 代码格式化
    ```shell
        make fmt
    ```
4. 项目启动
    ```shell
        make run
    ```

## 技术栈

- HTTP框架: go-zero + gorm
- 数据库：Mysql
- 缓存：Redis
- 搜索引擎：MeiliSearch
- 对象存储：七牛云存储
- 数据库和索引同步：MeiliSync
- 消息队列: kafka
- 项目管理: makefile
- 前后端协同: goctl-swagger + apifox

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

## API Doc

```shell
# 生成openAPI文档, 生成的文档在docs/swagger
# 可以使用ApiFox订阅 http://localhost:8080/api/v1/swagger
make gen-api-swagger 
# 生成api markdown文档, 生成的文档在docs/api
make gen-api-doc
```
