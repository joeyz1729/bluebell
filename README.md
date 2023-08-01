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

# 技能清单
* Viper配置管理
* Zap日志库
* 雪花算法 
* JWT认证 
* 令牌桶限流
* Gin框架
* Go语言操作MySQL (sqlx)
* Go语言操作Redis (go-redis)

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

# API
1. 基础接口
   * 视频流接口，对应community列表和:id
   * 用户登陆和注册接口
   * 用户信息
   * 视频投稿，对应post
   * 发布列表，对应用户uid的所有帖子
2. 互动接口
   * 赞操作，对应join community
   * 喜欢列表，对应查询user join的所有community列表
   * 评论操作，与视频投稿类似，TODO
   * 视频评论列表，对应community 下的所有post列表
3. 社交接口
   * 关注操作
   * 获取关注列表
   * 获取粉丝列表
   * 获取好友列表 TODO


# TODO
* 修改数据查询流程，如果redis中不存在，则从数据库中查询
* controller用结构体form传值
* service/logic层用interface+method
* logic层添加rabbitmq，异步处理请求
* 首先检查redis是否存在，如果查询成功则更新expire，否则从mysql中查询，并加入redis中
