# 电商系统

## 项目简介
个人觉得一个标准的Web项目应该都有了
- 用户模块
- 商品模块
- 购物车模块
- 订单模块
- 支付模块
- 收藏夹模块

## 使用技术
- Gorm
- Redis
- Gin
- MySQL
- Viper
- Logrus
- 其他不说了

## 项目结构
```
├── README.md
├── api
├── config
├── docs
├── global
├── cmd
├── types
├── pkg
├── respository
├── middelware
├── service
├── router
├── static
├── model
├── docker-compose.yml
├── Dockerfile
```
## 总结
- 项目结构清晰
- 项目功能完整

## 完善
- 项目日志感觉用zap写,并且日志的记录很乱
- Gorm的bug不知道怎么解决
- 许多查询其实可以使用缓存
- cmd只有一个main.go文件,感觉不够
- 权限的设置挺想用casbin的
- 没了，太懒了


