package router

import (
	"bytes"
	"encoding/json"
	"io"
)

type Entity interface {
	Type() string
	Data() io.Reader
}

type readerEntity struct {
	t string
	d io.Reader
}

func (r readerEntity) Type() string {
	return r.t
}

func (r readerEntity) Data() io.Reader {
	return r.d
}

func NewReaderEntity(t string, d io.Reader) (*readerEntity, error) {
	return &readerEntity{t, d}, nil
}

func NewBytesEntity(t string, d []byte) (*readerEntity, error) {
	return &readerEntity{t, bytes.NewReader(d)}, nil
}

func NewStringEntity(t string, d string) (*readerEntity, error) {
	return &readerEntity{t, bytes.NewReader([]byte(d))}, nil
}

func NewJSONEntity(e interface{}) (*readerEntity, error) {
	d, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return NewBytesEntity("application/json", d)
}
