# go-hmdp

使用 Go 重写黑马点评的学习项目，基于 Gin、GORM 和 MySQL，逐步实现业务功能。

目前处于基础工程搭建阶段：已接入配置加载、数据库连接和路由骨架；业务接口、数据库模型和初始化脚本仍在编写中。

## 目录结构

```text
go-hmdp/
├── main.go                     # 后端服务入口
├── config/                     # 环境变量、MySQL 和 Redis 配置
├── router/                     # 路由注册，按 user、shop 等模块拆分
├── handler/                    # 请求处理层
├── service/                    # 业务逻辑层
├── model/                      # 数据库模型及当前数据库连接初始化
│   ├── init.go                 # 初始化 GORM 数据库连接
│   └── tb_*.go                 # 数据表对应模型
├── dto/                        # 请求与返回数据结构
├── middleware/                 # 中间件
├── response/                   # 统一响应结构、业务状态码
├── utils/                      # 公共工具
├── scripts/                    # 独立的数据库初始化程序
│   ├── main.go                 # 初始化程序入口
│   ├── init_schema.go          # 建表初始化
│   ├── init_data.go            # 初始数据导入
│   └── sql/
│       ├── schema/             # 自行编写的建表 SQL
│       │   ├── user.sql
│       │   ├── shop.sql
│       │   └── blog.sql
│       └── data/               # 初始数据 SQL
├── docs/                       # 本地说明与接口参考
├── nginx-1.18.0/                # 本地前端和 Nginx
├── tmp/                        # 本地临时编译产物
├── .env.example                # 环境变量模板
├── .env                        # 本地环境变量
├── go.mod
└── go.sum
```
## 配置与启动

准备 MySQL，并先创建目标数据库。Go 版本要求以 `go.mod` 为准。


```powershell
go mod download
go run .
```

使用上面的端口配置后，访问 `http://localhost:8081/` 可查看基础响应。其他业务路由目前仍是占位，不代表已具备登录、商铺查询等能力。

## 数据库结构与初始化

表结构由自行编写的 `scripts/sql/schema/*.sql` 管理，初始数据放在 `scripts/sql/data/`。一个文件可以包含多张相关表，按 user、shop、blog 等业务模块划分。

- 当前 schema 下三个 SQL 文件仍为空，需要先补充建表语句。
- 当前 GORM 初始化仅连接数据库，未启用 `AutoMigrate`，也不会自动执行 SQL 文件。
- 可以先通过数据库工具执行自行编写的 SQL：先建表，再插入初始数据。
- 已有表的结构变更需要使用 `ALTER TABLE` 或后续迁移脚本；修改原始 `CREATE TABLE` 不会自动更新数据库。

`scripts` 是独立的 `main` 包，与根目录服务入口分开。初始化程序实现后，可从项目根目录执行：

```powershell
go run ./scripts
```