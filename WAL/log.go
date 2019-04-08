package wal

import (
	"KVStore/WAL/pb"
	"fmt"
	"os"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	encoder *encoder
	f       *os.File
}

func NewLog(path string) *Log {
	//  f, err := os.Open(path)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("open file err = ", err)
		f, err = os.Create(path)
		if err != nil {
			return nil
		}
	}
	encoder := newEncoder(f)
	return &Log{
		encoder: encoder,
	}
}

func (l *Log) AppendLog(rec *pb.Record) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.encoder.encode(rec)
}
