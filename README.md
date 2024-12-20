<div align="center">
  <h3 align="center">go-easy</h3>

  <p align="center">
    基于Gin和DDD领域模型思想,搭配 <a href="https://github.com/google/wire">wire</a> 依赖注入模式可写出更简洁和安全的代码。
    <br />
    <a href="https://github.com/Guohuixixi/go-easy"><strong>Explore »</strong></a>
    <br />
  </p>
</div>

## Dependencies
- [Golang1.23](https://github.com/golang/go)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/jinzhu/gorm)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [Redis](https://github.com/go-redis/redis)

## Getting Started
使用 go-easy 快速构建一个新项目:
```bash
# 下载 go-easy
git clone https://github.com/Guohuixixi/go-easy

# 进入项目目录
cd go-easy

# 安装 go mod 依赖
go mod tidy
```

## Todo
- [ ] 集成 gorm
- [ ] 日志接口并集成 gorm
- [ ] 集成 JWT
- [ ] 基础用户模块实现
- [ ] 短信验证服务模块开发
- [ ] 集成 Kafka
- [ ] 集成 swagger
