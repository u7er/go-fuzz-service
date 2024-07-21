package storage

import (
	"log"
	"os"
	"sync"
)

type Storage struct {
	mux      *sync.RWMutex
	Fd       *os.File
	FileName string
}

func NewStorage(fileName string) (*Storage, error) {
	fd, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil, err
	}

	return &Storage{
		mux:      &sync.RWMutex{},
		Fd:       fd,
		FileName: fileName,
	}, nil
}

func (s *Storage) Close() {
	err := s.Fd.Close()
	if err != nil {
		log.Printf("Error closing file: %v", err)
		return
	}
}

func (s *Storage) StoreValue(v string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	written, err := s.Fd.WriteString(v + "\n")
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return
	}
	log.Printf("Wrote %d bytes for %s", written, v)
}
