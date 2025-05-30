package main

import (
	"math/rand"
	"sync"
)

type QuoteStore struct {
	sync.Mutex
	quotes []Quote
	nextID int
}

func NewStore() *QuoteStore {
	return &QuoteStore{quotes: []Quote{}, nextID: 1}
}

func (s *QuoteStore) Add(q Quote) Quote {
	s.Lock()
	defer s.Unlock()
	q.ID = s.nextID
	s.nextID++
	s.quotes = append(s.quotes, q)
	return q
}

func (s *QuoteStore) GetAll() []Quote {
	s.Lock()
	defer s.Unlock()
	return append([]Quote(nil), s.quotes...)
}

func (s *QuoteStore) GetRandom() (Quote, bool) {
	s.Lock()
	defer s.Unlock()
	if len(s.quotes) == 0 {
		return Quote{}, false
	}
	return s.quotes[rand.Intn(len(s.quotes))], true
}

func (s *QuoteStore) FilterByAuthor(author string) []Quote {
	s.Lock()
	defer s.Unlock()
	var result []Quote
	for _, q := range s.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}
	return result
}

func (s *QuoteStore) DeleteByID(id int) bool {
	s.Lock()
	defer s.Unlock()
	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return true
		}
	}
	return false
}
