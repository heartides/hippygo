
## 项目介绍

Golang Learns Based On Gin

## 脚手架

> 基本方案
```bash
├── app # 应用中心
│   ├── controller # 控制器
│   │   └── post.go
│   └── model # 数据模型
│       └── user.go
├── conf # 配置中心
│   └── conf.go
├── databases # 数据库中心
│   ├── migration.go # 数据迁移生成
│   └── mysql.go
├── go.mod
├── go.sum
├── main.go
├── readme.md
├── routers # 路由
│   └── api.go
├── tmp # 临时目录
├── .env # 环境变量
├── .gitignore
└── utils # 工具包
    ├── hyper.go
    └── logger.go

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