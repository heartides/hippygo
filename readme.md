
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