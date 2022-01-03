# 高级编程

高级编程课设后端，采用 go-gin 编写的 http 服务

## 项目结构

```txt
- conf          // 配置文件
- config        // 读取配置文件
- constant      // 常量
- dal           // 数据库访问层，负责所有数据库增删改查
- handlers      // 路由层，负责各个路由组的路由服务
    - blog
    - common
    - job
    - user
- models        // 模型层，负责所有和数据库打交道的底层数据结构
- saved_files   // 存放上传的简历
- schema        // 负责所有与请求响应有关的数据结构
- services      // 负责请求和响应的处理
- DDL           // 存放建表的语句
- handler.go    // 路由的注册
- main.go       // main 函数
```

## 运行

配置 go 环境，用 go mod tidy 和 go get 拉取依赖

配置 go path，把项目放到 go path 下运行即可

需要配置数据库，建表语句放在了 DDL 文件夹

