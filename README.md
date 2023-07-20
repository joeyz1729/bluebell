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
