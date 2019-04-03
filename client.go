package main

import (
    "google.golang.org/grpc"
    "fmt"
	"context"
//	"strconv"
    pb "KVStore/API"
)
 
// 此处应与服务器端对应
const address  = "127.0.0.1:50051"
 
/**
    1. 创建groc连接器
    2. 创建grpc客户端,并将连接器赋值给客户端
    3. 向grpc服务端发起请求
    4. 获取grpc服务端返回的结果
 */

func main()  {
 
    // 创建一个grpc连接器
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil{
        fmt.Println(err)
    }
    // 当请求完毕后记得关闭连接,否则大量连接会占用资源
    defer conn.Close()
 
    // 创建grpc客户端
    c := pb.NewStoreServiceClient(conn)
 
    // 客户端向grpc服务端发起请求
    // 获取服务端返回的结果
    c.Put(context.Background(), &pb.PutRequest{Key:"key",Value:"val"})

    result, err := c.Get(context.Background(), &pb.GetRequest{Key:"key"})
    fmt.Println(result)
    if err != nil{
        fmt.Println("请求失败!!!")
        return
    }
    fmt.Println(result.Value)
}
