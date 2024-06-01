package storage

import (
	"fcompressor/env"
	"io"
)

type Metadata struct {
	Name string
	Size int64
	Mime string
	Path string
}

type Storage interface {
	Serve(filename string) (io.ReadCloser, error)
	Save(filename string, src io.Reader) (path string, err error)
	Remove(filename string) error
}

type Driver = string

var factories = map[Driver]Storage{}

func register(driver Driver, storage Storage) {
	factories[driver] = storage
}

func New(driver ...Driver) Storage {
	d := env.Get("STORAGE_DRIVER").String(Local)

	if len(driver) > 0 {
		d = driver[0]
	}

	return factories[d]
}