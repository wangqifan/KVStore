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
     RecordPb  "KVStore/WAL/pb"
     "KVStore/WAL"
     "KVStore/Util"
)
  
const (
    port = ":50051"
)
  
type server struct {
    skiplist *SkipList.ConcurrentSkipList
        log     *wal.Log
}
  
func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
  
    key, err := util.StingTounin64(in.Key)
    if err != nil {
        return nil, err
    }

    node, result := s.skiplist.Search(key)
    if !result {
      return   nil,errors.NewKeyNotFoundError()
    }
    return &pb.GetReply{Value:node.Value()}, nil
}

func (s *server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutReply, error) {

    key, err := util.StingTounin64(in.Key)
    if err != nil {
        return nil, err
    }
    
    value, err := util.StringToArray(in.Value)
    if err != nil {
        return nil, err
    }

    s.skiplist.Insert(key, value)

    record := &RecordPb.Record{
               Type : 1,
               Key : in.Key,
               Value : in.Value,
    }
    s.log.AppendLog(record)

    return &pb.PutReply{IsSuccess:true}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {

    key, err := util.StingTounin64(in.Key)
    if err != nil {
        return nil, err
    }

    s.skiplist.Delete(key)
    
    record := &RecordPb.Record{
        Type : 2,
        Key : in.Key,
    }
    s.log.AppendLog(record)

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
    server.log = wal.NewLog("data.dat")

    play := wal.NewWalReplay("data.dat")
    if play != nil {
        play.ReadAll(server.skiplist)
    }
   
    pb.RegisterStoreServiceServer(s, server)
    s.Serve(lis)
}