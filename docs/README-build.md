# ⚡🌈 💌EdgeSys Build

## 运行环境
go 1.19.4
mysql 8
redis 5及以上版本

## 编译方式
环境配置：需将go源调整到中国源
go env -w GOPROXY=https://goproxy.cn,direct

```依赖处理```
go mod tidy

```编译/调试源码```
go build -o edgesys main.go
或调试
go run main.go


