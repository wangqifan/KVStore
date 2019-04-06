# KVStore
key-value store

## 简介
 这是一个内存数据库，所有数据放在一个跳表中，使用grpc对外提供 get put delete scan四个函数的远程调用。使用proto buff对写记录进行序列化保存到硬盘当中。
 ## 运行与测试
 安装 protoc buff的go 包
 ~~~
 go get github.com/golang/protobuf/proto
 ~~~
 安装grpc 
 ~~~
 Git clone https://github.com/grpc/grpc-go
 ~~~
将grpc-go更名为grpc放入到google.golang.org中

运行服务
~~~
go run server.go
~~~
测试
~~~
go run client.go
~~~

## 关于WAL
Write-Ahead Logging 来保证数据持久化
* 写log---使用protoc buff 将record序列化，不满8比特打padding,使用uint64类型表示数据长度和padding长度，写入文件先写入uint64的长度数据，然后再写入对象序列化数据
* log 重放，先读出uint64 数据吗，然后再取出数据长度和padding长度，最后反序列号。

## 文件描述

* KVStore/API  定义了对外服务的get put delete scan四个函数的接口
* KVStore/Errors 自定义错误，对于客户端调用函数时，参数不合法会返回相应的错误类型
* KVStore/SkipList 跳表的实现
* KVStore/Util 基础设施，实现了类型转化函数
* KVStore/WAL WAL的实现
   * KVStore/WAL/pb 定义一条日志记录的格式
   * decoder.go 实现对日志的读
   * encoder.go 实现对日志的写
   * log.go 封装对单条日志append写
   * WalReplay.go 回放日志
* server.go 入口&& grpc服务的实现
* client.go  测试
