# WeChat API（iPad Protocol）

基于iPad协议开发的微信机器人，仅用于学习与开发，切勿商用，否则后果自负。

![Golang](https://img.shields.io/badge/language-Golang-blue.svg)
![Beego](https://img.shields.io/badge/framework-Beego-green.svg)

## 项目结构

```
wechat-api-ipad/
│
├── Cilent/                 # 客户端相关逻辑处理
│   ├── HybridEcdh.go      # 混合ECDH算法实现
│   ├── NewSpamData.go     # 新Spam数据处理
│   ├── common.go          # 公共数据处理
│   ├── device/            # 设备相关数据和操作
│   ├── mm/                # 微信相关原型定义
│   └── ...
│
├── Fun/                    # 功能函数
│   └── Hex2Int.go         # 十六进制与整型转换函数
│
├── Mmtls/                  # MMTLS协议实现
│   ├── GenNewHttpClient.go# 生成新的HTTP客户端
│   ├── InitMmtlsShake.go  # 初始化MMTLS握手
│   └── ...
│
├── controllers/            # MVC 控制器
│   ├── Favor.go           # 收藏相关操作
│   ├── Finder.go          # 发现页相关操作
│   └── ...
│
├── models/                 # 数据模型和业务逻辑
│   ├── Favor/             # 收藏相关模型
│   ├── Finder/            # 发现页相关模型
│   └── ...
│
├── conf/                   # 配置文件夹
│   └── app.conf
│
├── main.go                 # 程序入口文件
└── go.mod                  # Go模块定义
```


## 开始使用

### 1. 安装依赖

- 环境要求：Go 1.15+，MySQL 5.7+，Redis 5.0+
- 安装依赖：`go mod tidy`
- 配置文件：`conf/app.conf`
```go
go mod tidy

go run main.go
```

### 交流群

![微信](https://img.shields.io/badge/微信-扫码加群-green.svg)

<img src="./github/img/wechat.jpg" width="210px" style="border-radius: 20px;border: 3px solid #53e339">


## 许可和声明
<p style="color: rgba(224,9,206,0.74)">该项目是公开的，任何人都可以自由地使用和分发。作者不承担由于使用该代码而可能导致的任何后果。</p>