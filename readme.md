
## 项目介绍

Golang Learns Based On Gin

## 脚手架

> 基本方案
```bash
├── app
│   ├── controller # 控制器
│   │   └── post.go 
│   ├── middleware  # 中间件
│   │   └── example.go
│   └── model   # 数据处理
│       └── Posts.go
├── conf    # 配置中心
│   ├── conf.go
│   ├── mysql.go
│   └── redis.go
├── databases # 数据库实例
│   ├── migration.go
│   └── mysql.go
├── go.mod
├── go.sum
├── main.go
├── readme.md
├── routers # 路由中心
│   └── api.go
├── tmp # 临时目录
└── utils # 工具包
    ├── context.go
    ├── hyper.go
    ├── logger.go
    └── response.go

```

> 扩展

- 对gin的gin.Context中的JSON做一层封装，实现Success()、Fail()两个响应方法

```golang
// controller
func Test(ctx *utils.Context){}

// router
router.GET("/test", utils.Handle(controller.Test))

// 使用

// 正确响应
cxt.Success("OK",gin.H{"name":"walker"})

// 返回
{
    "code": 0,
    "msg": "OK",
    "data": {
        "name": "walker"
    }
}

// 错误响应
cxt.Fail(422,登录失败")

// 返回
{
    "code": 422,
    "msg": "登录失败"
}
    
```

## ENV 环境变量说明
```bash
# debug = 开发 release = 生产
GIN_MODE="debug"
LOG_LEVEL="debug"

# ON 开启数据连接
DB_CONNON="ON"
# 数据库地址
DB_HOST="localhost"
# 数据库端口
DB_PORT="3306"
# 数据库名称
DB_DATABASE="gin"
# 用户名
DB_USERNAME="root"
# 密码
DB_PASSWORD=""
# 编码
DB_CHARSET="utf8mb4"

# ON 开启Redis连接
REDIS_CONNON="ON"
# Redis地址
REDIS_HOST="127.0.0.1"
# Redis 端口
REDIS_PORT="6379"
# Redis密码
REDIS_AUTH=""
# Redis 存储
REDIS_DB="0"

```