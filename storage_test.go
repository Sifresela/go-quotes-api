package main

import (
	"testing"
)

func TestAddAndGetAll(t *testing.T) {
	store := NewStore()
	q := Quote{Author: "Test", Text: "Hello"}
	added := store.Add(q)

	if added.ID != 1 {
		t.Errorf("expected ID 1, got %d", added.ID)
	}

	all := store.GetAll()
	if len(all) != 1 {
		t.Fatalf("expected 1 quote, got %d", len(all))
	}

	if all[0].Author != "Test" {
		t.Errorf("expected author Test, got %s", all[0].Author)
	}
}

func TestFilterByAuthor(t *testing.T) {
	store := NewStore()
	store.Add(Quote{Author: "A", Text: "1"})
	store.Add(Quote{Author: "B", Text: "2"})
	store.Add(Quote{Author: "A", Text: "3"})

	result := store.FilterByAuthor("A")
	if len(result) != 2 {
		t.Errorf("expected 2 results, got %d", len(result))
	}
}

func TestDeleteByID(t *testing.T) {
	store := NewStore()
	store.Add(Quote{Author: "X", Text: "Delete me"})

	deleted := store.DeleteByID(1)
	if !deleted {
		t.Error("expected deletion to succeed")
	}

	deletedAgain := store.DeleteByID(1)
	if deletedAgain {
		t.Error("expected second deletion to fail")
	}
}
