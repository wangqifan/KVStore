package main
  
import (
    "log"
    "net"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
     pb "KVStore/API"
     "KVStore/SkipList"
     "fmt"
     "KVStore/errors"
)
  
const (
    port = ":50051"
)
  
type server struct {
    skiplist *SkipList.ConcurrentSkipList 
}
  
func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
  
    chars := []byte(in.Key)
    if len(chars) ==0 || len(chars) >8 {
        return nil, errors.NewKeyInvaildError()
    } 
    var  num uint64 = 0
    for _,val := range chars {
        num = num*128 + uint64(val)
    }
    node, result := s.skiplist.Search(num)
    if !result {
      return   nil,errors.NewKeyNotFoundError()
    }
    return &pb.GetReply{Value:node.Value()}, nil
}

func (s *server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutReply, error) {

    chars := []byte(in.Key)
    if len(chars) ==0 || len(chars) >8 {
        return nil, errors.NewKeyInvaildError()
    } 

    var num uint64 = 0 
    for _,val := range chars {
        num = num*128 + uint64(val)
    }

    var value []byte = []byte(in.Value)
    if len(value) > 256 || len(value) ==0 {
        return &pb.PutReply{IsSuccess : false}, errors.NewValueInvaildError()
    }
    var temp [256]byte
    for i, item := range value {
        temp[i] = item
    }
    s.skiplist.Insert(num, temp)
    return &pb.PutReply{IsSuccess:true}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {

    chars := []byte(in.Key)  
    if len(chars) ==0 || len(chars) >8 {
        return nil, errors.NewKeyInvaildError()
    } 
    var num uint64 = 0 
    for val := range chars {
        num = num*128 + uint64(val)
    } 
    s.skiplist.Delete(num)
    return &pb.DeleteReply{IsSuccess:true}, nil
}

func (s *server) Scan(ctx context.Context, in *pb.ScanRequest) (*pb.ScanReply, error) {
    
    if in.Start < 0 || in.Limit < 1 {
        return nil, errors.NewScanParameterInvaildError()
    }
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