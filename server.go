package main
  
// server.go
  
import (
    "log"
    "net"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
     pb "KVStore/API"
     "KVStore/SkipList"
     "fmt"
)
  
const (
    port = ":50051"
)
  
type server struct {
    skiplist *SkipList.ConcurrentSkipList 
}
  
func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
    var  num uint64 = 0
    chars := []byte(in.Key)  
    for _,val := range chars {
        num = num*128 + uint64(val)
    }
    node, result := s.skiplist.Search(num)
    if !result {
      return   &pb.GetReply{Value:""}, nil
    }
    return &pb.GetReply{Value:node.Value()}, nil
}

func (s *server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutReply, error) {
    var num uint64 = 0 
    chars := []byte(in.Key)  
    for _,val := range chars {
        num = num*128 + uint64(val)
    }
    var value []byte = []byte(in.Value)
    if len(value) > 256 {
        return &pb.PutReply{IsSuccess:false}, nil
    }
    var temp [256]byte
    for i, item := range value {
        temp[i] = item
    }
    s.skiplist.Insert(num, temp)
    return &pb.PutReply{IsSuccess:true}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {
    var num uint64 = 0 
    chars := []byte(in.Key)  
    for val := range chars {
        num = num*128 + uint64(val)
    } 
    s.skiplist.Delete(num)
    return &pb.DeleteReply{IsSuccess:true}, nil
}

func (s *server) Scan(ctx context.Context, in *pb.ScanRequest) (*pb.ScanReply, error) {
    nodes := s.skiplist.Sub(in.Start, in.Limit)
    var strs  []string
    for _,node :=range nodes {
        strs = append(strs, node.Value())
    }
    return &pb.ScanReply{Result:strs}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatal("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    skipList, err := SkipList.NewConcurrentSkipList(16)
    if err != nil {
         fmt.Println(err)
    }
    server := &server{}
    server.skiplist = skipList
    pb.RegisterStoreServiceServer(s, server)
    s.Serve(lis)
}