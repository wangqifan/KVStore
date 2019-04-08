package wal

import (
	"KVStore/WAL/pb"
	"bufio"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"io"
	"sync"
)

const minSectorSize = 512

// frameSizeBytes is frame size in bytes, including record size and padding size.
const frameSizeBytes = 8

type decoder struct {
	mu sync.Mutex
	br *bufio.Reader
}

func newDecoder(r *bufio.Reader) *decoder {
	return &decoder{br: r}
}

func (d *decoder) decode(rec *pb.Record) error {
	rec.Reset()
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.decodeRecord(rec)
}

func (d *decoder) decodeRecord(rec *pb.Record) error {
	if d.br == nil {
		return io.EOF
	}

	l, err := readInt64(d.br)
	if err == io.EOF || (err == nil && l == 0) {
		return io.EOF
	}
	if err != nil {
		return err
	}

	recBytes, padBytes := decodeFrameSize(l)

	data := make([]byte, recBytes+padBytes)
	if _, err = io.ReadFull(d.br, data); err != nil {
		// ReadFull returns io.EOF only if no bytes were read
		// the decoder should treat this as an ErrUnexpectedEOF instead.
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return err
	}
	if err := proto.Unmarshal(data[:recBytes], rec); err != nil {
		return err
	}

	return nil
}

func decodeFrameSize(lenField int64) (recBytes int64, padBytes int64) {
	// the record size is stored in the lower 56 bits of the 64-bit length
	recBytes = int64(uint64(lenField) & ^(uint64(0xff) << 56))
	// non-zero padding is indicated by set MSb / a negative length
	if lenField < 0 {
		// padding is stored in lower 3 bits of length MSB
		padBytes = int64((uint64(lenField) >> 56) & 0x7)
	}
	return recBytes, padBytes
}

func readInt64(r io.Reader) (int64, error) {
	var n int64
	err := binary.Read(r, binary.LittleEndian, &n)
	return n, err
}
