# 项目结构
<pre>
<code>
bluebell
├─ .air.conf
├─ README.md
├─ conf
│    └─ config.yaml
├─ controller
│    ├─ code.go
│    ├─ community.go
│    ├─ post.go
│    ├─ request.go
│    ├─ response.go
│    ├─ user.go
│    ├─ validator.go
│    └─ vote.go
├─ dao
│    ├─ mysql
│    │    ├─ code.go
│    │    ├─ community.go
│    │    ├─ mysql.go
│    │    ├─ post.go
│    │    └─ user.go
│    └─ redis
│           ├─ key.go
│           ├─ post.go
│           ├─ redis.go
│           └─ vote.go
├─ go.mod
├─ go.sum
├─ log
│    └─ web-app.log
├─ logger
│    └─ logger.go
├─ logic
│    ├─ community.go
│    ├─ post.go
│    ├─ user.go
│    └─ vote.go
├─ main.go
├─ middleware
│    ├─ auth.go
│    └─ ratelimit.go
├─ model
│    ├─ community.go
│    ├─ form.go
│    ├─ post.go
│    ├─ sql
│    │    ├─ bluebell_community.sql
│    │    ├─ bluebell_post.sql
│    │    ├─ bluebell_user.sql
│    │    └─ init.sql
│    └─ user.go
├─ pkg
│    ├─ jwt
│    │    └─ jwt.go
│    └─ snowflake
│           └─ gen_id.go
├─ router
│    └─ router.go
└─ setting
└─ setting.go
</code>
</pre>

# API
1. User 相关
   * POST 用户注册操作
   * POST 用户登陆操作
   * GET 获取用户详细信息
   * GET jwt鉴权，生成access token和fresh token，通过fresh token刷新

2. Community 相关
   * GET 获取所有社区信息
   * GET 获取用户关注的社区列表
   * GET 获取指定社区的详细信息
   * POST 用户加入社区操作

3. Post 相关
   * GET 分页按顺序获取帖子信息
   * GET 获取社区所有帖子信息
   * GET 获取用户所有发帖信息
   * GET 获取指定帖子信息
   * POST 用户发帖操作
   * POST 用户投票操作

4. Follow 相关
   * POST 用户关注操作
   * GET 获取关注列表
   * GET 获取粉丝列表
   * GET 获取好友列表 TODO

5. Comment 相关 TODO

# 技能清单
* Viper配置管理
* Zap日志库
* 雪花算法 
* JWT鉴权
* 令牌桶限流
* Gin框架
* Validator参数校验
* Sqlx操作MySQL
* GO-redis操作Redis

# 启动流程
1. 修改conf/config.yaml 文件中 host，port，MySQL，Redis等配置
2. 连接上MySQL数据库，并建库建表和插入初始数据
    1. init.sql
    2. bluebell_user.sql
    3. bluebell_community.sql
    4. bluebell_post.sql
3. 执行 `go build -o ./bin/bluebell`，编译可执行文件至项目的bin目录
4. 执行 `./bin/bluebell conf/config.yaml`，启动程序
5. 打开浏览器或postman测试 http://127.0.0.1:8081




# 改进
* 数据获取流程，更新删除redis和延迟修改mysql
* 添加rabbitmq，join, post, follow等请求放入消息队列异步处理
* 规范controller层request，response form
* 修改logic层，使用interface+method
* 使用微服务架构
