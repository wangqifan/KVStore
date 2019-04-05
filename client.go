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

    _, err = c.Get(context.Background(), &pb.GetRequest{Key:"dsf"})
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("get" + "hello")
    result, err := c.Get(context.Background(), &pb.GetRequest{Key:"hello"})
    fmt.Println(result)

/*
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"1",Value:"1+1"})
    _, err =c.Put(context.Background(), &pb.PutRequest{Key:"2",Value:"2+1"})
    _, err =c.Put(context.Background(), &pb.PutRequest{Key:"3",Value:"3+1"})
    _, err =c.Put(context.Background(), &pb.PutRequest{Key:"4",Value:"4+1"})
    _, err =c.Put(context.Background(), &pb.PutRequest{Key:"5",Value:"5+1"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"6",Value:"6+1"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"7",Value:"7+1"})
    _, err =c.Put(context.Background(), &pb.PutRequest{Key:"8",Value:"8+1"})
    _, err =c.Put(context.Background(), &pb.PutRequest{Key:"9",Value:"9+1"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"10",Value:"10+1"})
   */ 
  /*  fmt.Println("put " + "name" +"  wangqifan" )
    c.Put(context.Background(), &pb.PutRequest{Key:"name",Value:"wangqifan"})

    fmt.Println("put " + "age" +"  23" )
    c.Put(context.Background(), &pb.PutRequest{Key:"age",Value:"23"})
    fmt.Println("Get" + "name")
    result, err := c.Get(context.Background(), &pb.GetRequest{Key:"name"})
   // fmt.Println(result)
    if err != nil{
        fmt.Println("请求失败!!!")
        return
    }
    fmt.Println(result.Value)

    fmt.Println("get" + " age")
    resultage, err := c.Get(context.Background(), &pb.GetRequest{Key:"age"})
   // fmt.Println(result)
    if err != nil{
        fmt.Println("请求失败!!!")
        return
    }
    fmt.Println(resultage.Value)

    _,err = c.Delete(context.Background(), &pb.DeleteRequest{Key:"name"})
    fmt.Println("Get" + "name")
    resultd, err := c.Get(context.Background(), &pb.GetRequest{Key:"name"})
   // fmt.Println(result)
    if err != nil{
        fmt.Println("请求失败!!!")
        return
    }
    fmt.Println(resultd)
    */
    /*scanResult, err :=c.Scan(context.Background(),&pb.ScanRequest{Start:0,Limit:10})
    for _,str := range scanResult.Result {
        fmt.Println(str)
        fmt.Println("-------------")
    } */
   
}
