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


/*
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"hello",Value:"world"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"school",Value:"bit"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"age",Value:"24"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"city",Value:"beijing"})
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"country",Value:"china"})
*/
    fmt.Println("-------------------------------------------------------")
    fmt.Println("test for WAL")
    fmt.Println("school age country,the keys,had been inserted")
    fmt.Println("Get  school")
    GetResult, err := c.Get(context.Background(), &pb.GetRequest{Key:"school"})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(GetResult.Value)
    }
  
    fmt.Println("Get  age")
    GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key:"age"})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(GetResult.Value)
    }
    
    fmt.Println("Get country")
    GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key:"country"})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(GetResult.Value)
    }
    

    fmt.Println("-------------------------------------------------------")
    fmt.Println("test for Get")
    fmt.Println("Get  notexit")
    GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key:"notexit"})
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println("Get  akeyverylong")
    GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key:"akeyverylong"})
    if err != nil {
        fmt.Println(err)
    }
    

    fmt.Println("-------------------------------------------------------")
    fmt.Println("test for PUT")
    fmt.Println("PUT how fine")
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"how",Value:"fine"})
    if err == nil {
        fmt.Println("insert how success")
    } else {
        fmt.Println("insert how fail")
    }
    fmt.Println("Get how")
    GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key:"how"})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(GetResult.Value)
    }
    
    fmt.Println(`PUT longval sdaasdfjsdddddddddddaaggdsadfhhhhhhhhhhhhhhhhjsdfakjjjjjkjsdfajkasvs
    sfh3et22222gdddddddddddddddddddddddddddddddddddddddddascfeeeeeeeeeeeeeeeeedzdxvfeasfcweasdfcas
    dddddddddddddddddddddddddddddddddddddddddddddddddddddddddasddddddddddddddddddddasdasdaADASDSAD
    sadfannngoooooooooooooooooooooooooddddddddddddddddddddddddddfi`)
    _, err = c.Put(context.Background(), &pb.PutRequest{Key:"longval",Value:`sdaasdfjsdddddddddddaaggdsadfhhhhhhhhhhhhhhhhjsdfakjjjjjkjsdfajkasvs
    sfh3et22222gdddddddddddddddddddddddddddddddddddddddddascfeeeeeeeeeeeeeeeeedzdxvfeasfcweasdfcas
    dddddddddddddddddddddddddddddddddddddddddddddddddddddddddasddddddddddddddddddddasdasdaADASDSAD
    sadfannngoooooooooooooooooooooooooddddddddddddddddddddddddddfi`})
    if err != nil {
        fmt.Println(err)
    }
  

    fmt.Println("-------------------------------------------------------")
    fmt.Println("test for delete")
    fmt.Println("delete how")
    _,err = c.Delete(context.Background(), &pb.DeleteRequest{Key:"how"})
    if err == nil {
        fmt.Println("delete how success")
    } else {
        fmt.Println("delete how fail")
    }
    fmt.Println("Get how")
    GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key:"how"})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(GetResult.Value)
    }
    

    fmt.Println("-------------------------------------------------------")
    fmt.Println("test for scan")
    fmt.Println("scan 0 3")
    scanResult, err :=c.Scan(context.Background(),&pb.ScanRequest{Start:0,Limit:3})
    if err != nil {
        fmt.Println(err)
    } else {
        for _,str := range scanResult.Result {
            fmt.Println(str)
        }
    }
    fmt.Println("scan -1 3")
    scanResult, err =c.Scan(context.Background(),&pb.ScanRequest{Start:-1,Limit:3})
    if err != nil {
        fmt.Println(err)
    } else {
        for _,str := range scanResult.Result {
            fmt.Println(str)
        }
    }

   
}
