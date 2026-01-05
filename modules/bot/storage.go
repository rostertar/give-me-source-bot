package bot

import "io"

type Storage interface {
	GetWriteCloser(category, id string) (io.WriteCloser, error)
	GetReadCloser(category, id string) (io.ReadCloser, error)
}
