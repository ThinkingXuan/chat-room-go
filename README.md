# chat-room-go
🎨chat-room-go 简单聊天室go语言实现的🌮

技术点：
- Gin框架、Fiber框架、Gorm框架、JWT验证
- Redigo框架、Viper
- MySQL、Redis数据库、leveldb

开发计划：
### 1. 完成项目基本框架搭建
### 2. 完善功能
   
#### 2.1 用户
- 注册 ✔
- 登录 ✔
- 获取用户信息 ✔
#### 2.2 房间
- 创建新房间 ✔
- 进入房间  ✔
- 离开房间   ✔
- 获取房间信息
- 获取房间的用户列表 ✔
- 获取房间列表 ✔

#### 2.3 信息
- 发送信息 ✔
- 信息检索 ✔

### 3. 测试
- Wrk脚本编写 ✔
- JMeter测试 ✔

### 4. 优化
- Gin框架优化 ✔
- 本地缓存、压缩存储等 ✔
- 减少业务逻辑 ✔
- MySQL索引优化 ✔
- MySQL分页优化 ✔
- 操作系统内核调参 ✔
- MySQL和Redis配置调参 ✔
- 改用Redis实现 ✔
- Redis连接池配置调参 ✔
- 使用epoll实现http协议 ✖
- 用消息队列进行批处理 ✖
- Fiber框架替换gin框架 ✔ (效果不明显且稳定性差，弃用)
- 数据序列化存储Redis ✔
- Redis延迟优化(请求合并) ✔
- Redis读写分离 ✔
- 使用分布式雪花ID替换UUID ✔

### 5. 集群配置
- Redis master-slave主从模式 ✔
- Redis-Sentinel集群 ✔
- Sentinel探针(检查master宕机) ✔
