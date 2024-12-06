package database

import (
	"sync"
)

type MemoryStorage struct {
	// you can ignore `mu` at the moment, it makes `MemoryStorage`
	// to be thread-safe for concurrency implementation later on
	mu        sync.RWMutex
	something map[string]int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		something: map[string]int{},
	}
}

func (s *MemoryStorage) GetTokenCount(name string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.something[name]
	if ok {
		return val
	}

	return 0
}

func (s *MemoryStorage) AddToken(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.something[name]++
}

func (s *MemoryStorage) UpdateTokenName(old string, new string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// find the old
	tmp := s.something[old]
	delete(s.something, old)
	s.something[new] = tmp
}

func (s *MemoryStorage) UpdateTokenCount(name string, count int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.something[name] = count
}

func (s *MemoryStorage) RemoveToken(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.something, name)
}

func (s *MemoryStorage) ResetToken(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.something[name] = 0
}

func (s *MemoryStorage) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.something)
}
