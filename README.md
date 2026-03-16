# 钱包服务 (Wallet Service)
一个简单的Go语言钱包服务，提供钱包查询和转账功能

# 项目结构
├── cmd/
│   └── server/          # 主程序入口
├── internal/            # 内部包
│   ├── handler/         # HTTP 处理器
│   └── service/         # 业务逻辑
├── data/                
│   └── model/           # 数据存储
├── docs/                # 接口文档
├── script/
│   └── sql/
│       └── init.sql     # 数据库初始化脚本
├── go.mod               # go mod
└── README.md            # 本文档

# 快速开始
1.创建MySQL容器
docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 mysql:8.0.28 --default-authentication-plugin=mysql_native_password --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci

2.初始化数据库 
docker exec -i mysql mysql -u root -p1234 < script/sql/init.sql

3.运行 Go 服务
go run ./cmd/server

# docker 构建镜像和运行
1. 构建镜像
   docker build -t wallet .
2. 运行容器，暴漏8080端口
   docker run -p 8080:8080 wallet

# swagger生成接口文档
1. 安装 swagger
go install github.com/swaggo/swag/cmd/swag@latest

2. 生成yaml格式接口文档
swag init -g internal/handler/wallet.go -o docs  --outputTypes yaml