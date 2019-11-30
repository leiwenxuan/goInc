# 简单调用GoInception 模块的测试代码
## go语言调用
服务地址： http://127.0.0.1/v1/rds/inception 
```json
 {
        "sql_str": "use test_01;",  //执行的SQL语句
        "sql_types": 1,       // 1 检查 2 执行
        "user": "root",       // 数据库账号
        "password": "123456", // 密码
        "host": "192.168.1.3",  // 数据库地址
        "port": "20001",      // 数据库端口
        "database": "test_01" // 数据库
 }
```