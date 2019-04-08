package main

import (
	pb "KVStore/API"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

const address = "127.0.0.1:50051"

func TestWal(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewStoreServiceClient(conn)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("test for WAL")
	fmt.Println("school age country,the keys,had been inserted")
	fmt.Println("Get  school")
	GetResult, err := c.Get(context.Background(), &pb.GetRequest{Key: "school"})
	if err != nil {
		t.Log("should return the value of school")
	} else {
		fmt.Println(GetResult.Value)
	}

	fmt.Println("Get  age")
	GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key: "age"})
	if err != nil {
		t.Log("should return the value of age")
	} else {
		fmt.Println(GetResult.Value)
	}

	fmt.Println("Get country")
	GetResult, err = c.Get(context.Background(), &pb.GetRequest{Key: "country"})
	if err != nil {
		t.Log("should return the value of country")
	} else {
		fmt.Println(GetResult.Value)
	}
}

func TestGet(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Log("can not access service")
	}
	defer conn.Close()
	c := pb.NewStoreServiceClient(conn)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("test for Get")
	fmt.Println("Get  notexit")
	_, err = c.Get(context.Background(), &pb.GetRequest{Key: "notexit"})
	if err == nil {
		t.Log("should retrun error for  the key not exist")
	}

	fmt.Println("Get  akeyverylong")
	_, err = c.Get(context.Background(), &pb.GetRequest{Key: "akeyverylong"})
	if err == nil {
		t.Log("should retrun error for  the key is too long")
	}
}

func TestPut(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewStoreServiceClient(conn)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("test for PUT")
	fmt.Println("PUT how fine")
	_, err = c.Put(context.Background(), &pb.PutRequest{Key: "how", Value: "fine"})
	if err == nil {
		fmt.Println("insert how success")
	} else {
		fmt.Println("insert how fail")
	}
	fmt.Println("Get how")
	GetResult, err := c.Get(context.Background(), &pb.GetRequest{Key: "how"})
	if err != nil {
		t.Log("should retrun the value of how")
	} else {
		fmt.Println(GetResult.Value)
	}

	fmt.Println(`PUT longval sdaasdfjsdddddddddddaaggdsadfhhhhhhhhhhhhhhhhjsdfakjjjjjkjsdfajkasvs
    sfh3et22222gdddddddddddddddddddddddddddddddddddddddddascfeeeeeeeeeeeeeeeeedzdxvfeasfcweasdfcas
    dddddddddddddddddddddddddddddddddddddddddddddddddddddddddasddddddddddddddddddddasdasdaADASDSAD
    sadfannngoooooooooooooooooooooooooddddddddddddddddddddddddddfi`)
	_, err = c.Put(context.Background(), &pb.PutRequest{Key: "longval", Value: `sdaasdfjsdddddddddddaaggdsadfhhhhhhhhhhhhhhhhjsdfakjjjjjkjsdfajkasvs
    sfh3et22222gdddddddddddddddddddddddddddddddddddddddddascfeeeeeeeeeeeeeeeeedzdxvfeasfcweasdfcas
    dddddddddddddddddddddddddddddddddddddddddddddddddddddddddasddddddddddddddddddddasdasdaADASDSAD
    sadfannngoooooooooooooooooooooooooddddddddddddddddddddddddddfi`})
	if err == nil {
		t.Log("should retrun error for value is too long")
	}
}

func TestDelete(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewStoreServiceClient(conn)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("test for delete")
	fmt.Println("delete how")
	_, err = c.Delete(context.Background(), &pb.DeleteRequest{Key: "how"})
	if err == nil {
		fmt.Println("delete how success")
	} else {
		fmt.Println("delete how fail")
	}
	fmt.Println("Get how")
	_, err = c.Get(context.Background(), &pb.GetRequest{Key: "how"})
	if err != nil {
		fmt.Println(err)
	} else {
		t.Log("should retrun error for  the key had been delete")
	}
}

func TestScan(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewStoreServiceClient(conn)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("test for scan")
	fmt.Println("scan 0 3")
	scanResult, err := c.Scan(context.Background(), &pb.ScanRequest{Start: 0, Limit: 3})
	if err != nil {
		fmt.Println(err)
	} else {
		for _, str := range scanResult.Result {
			fmt.Println(str)
		}
	}
	fmt.Println("scan -1 3")
	scanResult, err = c.Scan(context.Background(), &pb.ScanRequest{Start: -1, Limit: 3})
	if err != nil {
		fmt.Println(err)
	} else {
		t.Log("should return error for start is -1")
	}

}
