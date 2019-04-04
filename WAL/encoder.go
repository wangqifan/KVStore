package wal

import (
	"encoding/binary"
	"io"
	"sync"
	"KVStore/WAL/pb"
	"github.com/golang/protobuf/proto"
)


type encoder struct {
	mu sync.Mutex
	w io.Writer
	uint64buf []byte
}

func newEncoder(w io.Writer) *encoder {
	return &encoder{
		w:  w,
		uint64buf: make([]byte, 8),
	}
}


func (e *encoder) encode(rec *pb.Record) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	data, err := proto.Marshal(rec)
	lenField, padBytes := encodeFrameSize(len(data))
	
	if err = writeUint64(e.w, lenField, e.uint64buf); err != nil {
		return err
	}
	
	if padBytes != 0 {
		data = append(data, make([]byte, padBytes)...)
	}
	_, err = e.w.Write(data)
	return err
}

func encodeFrameSize(dataBytes int) (lenField uint64, padBytes int) {
	lenField = uint64(dataBytes)
	// force 8 byte alignment so length never gets a torn write
	padBytes = (8 - (dataBytes % 8)) % 8
	if padBytes != 0 {
		lenField |= uint64(0x80|padBytes) << 56
	}
	return lenField, padBytes
}

func writeUint64(w io.Writer, n uint64, buf []byte) error {
	// http://golang.org/src/encoding/binary/binary.go
	binary.LittleEndian.PutUint64(buf, n)
	_, err := w.Write(buf)
	return err
}
