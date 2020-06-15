# 框架文档
> 此文档是使用beego框架搭建web-api项目

## 技术架构
1. 项目中测试连接mysql、mongodb数据库和redis缓存
2. 项目使用swagger作为接口文档

## 项目代码结构
- conf：配置文件目录
        - app.conf：主配置配件
        - mysql.conf：mysql数据库配置文件，引入app.conf
- controller：web层接受请求
- models：Bean模型层，存放数据库对应的bean
- routers：url请求映射
- main.go：程序入口

## swagger API文档自动生成
1. 首先需要在app.conf配置文件中添加配置：`EnableDocs = true`
2. 在controller接口方法上面进行注解说明
3. 启动项目时使用下面的命令：`bee run -gendoc=true -downdoc=true`


## 关键模块
1. mysql数据库配置
    - conf/app.conf配置mysql数据库连接
    - main.go的init()方法中进行初始化mysql数据库连接池
2. jwt-go登陆校验
    - 依赖模块：https://github.com/dgrijalva/jwt-go
    - 相关代码包：jwtauth
        - 拦截器函数：JwtFilter
        - 自定义JwtUser用户实体
        - JwtUtil工具：token生成和解析token
    - 获取token和解析token的请求在controllers/jwt.go中
3. 系统日志模块
    - 日志配置文件：conf/logger.conf
    - 日志配置组件：component/logger.go
    - 在main.go的初始化函数Init()中进行初始化日志
    - 根据配置的日志级别log_level进行打印输出日志
    - 日志文件生成在配置的log_path目录下
4. Redis 缓存
    - 下载依赖：git clone https://github.com/gomodule/redigo
    - 在conf/redis.conf配置redis
    - 在模块component/redis.go中进行redis缓存配置
    - 在main.go的初始化函数Init()中进行初始化redis
