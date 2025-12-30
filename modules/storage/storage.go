package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rostertar/give-me-source-bot/modules/log"
)

// Storage хранилище для данных куда складируются данные
type Storage struct {
	root string
}

func NewStorage(root string) (*Storage, error) {
	stat, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return &Storage{root}, nil
	}
	return nil, fmt.Errorf("Not a dir")
}

func (s *Storage) GetWriteCloser(category, id string) (io.WriteCloser, error) {
	rd := filepath.Join(s.root, category)
	if stat, err := os.Stat(rd); err != nil {
		return nil, err
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("Target is not dir: %s", rd)
	}
	path := filepath.Join(s.root, category, id)
	if _, err := os.Stat(path); err == nil {
		log.Errorf("Target %s already exists — replacing it", path)
	}
	return os.Create(path)
}

func (s *Storage) GetReadCloser(category, id string) (io.ReadCloser, error) {
	path := filepath.Join(s.root, category, id)
	return os.Open(path)
}
