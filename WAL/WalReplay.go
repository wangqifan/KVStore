package wal


import (
	"os"
	"sync"
	"fmt"
	"bufio"
	"KVStore/WAL/pb"
	"KVStore/SkipList"
	"KVStore/Util"
)

type WalReplay struct {
	mu sync.Mutex
	decoder   *decoder
	f         *os.File
}

func NewWalReplay(path string) *WalReplay {
    f, err := os.Open(path)
    if err != nil {
      	 fmt.Println("open file err = ", err)
	     return nil
    }

	r := bufio.NewReader(f)
	return & WalReplay{
		decoder: newDecoder(r),
		f: f,
	}
}

func (play *WalReplay)ReadAll(skiplist *SkipList.ConcurrentSkipList ) {
	play.mu.Lock()
	defer play.mu.Unlock()
  
	fmt.Println("开始日志回放")
	record := &pb.Record{}
    count := 0
	for err := play.decoder.decode(record); err == nil; err = play.decoder.decode(record) {
		if record.Type == 1 {
			index, err:= util.StingTounin64(record.Key)
			if err != nil {
				continue;
			}
			value, err:= util.StringToArray(record.Value)
			if err != nil {
				continue;
			}
			skiplist.Insert(index, value)
		} else {
			index, err:= util.StingTounin64(record.Key)
			if err != nil {
				continue;
			}
			skiplist.Delete(index)
		}
		fmt.Println(record.Key + "  "+ record.Value)
		count++
	}
	fmt.Println("记录数")
	fmt.Println(count)
	play.f.Close()
}