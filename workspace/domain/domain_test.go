package domain

import (
	"testing"
)

func TestNewOrderRepositoryDb(t *testing.T) {
	repo := NeworderRepositoryDb()
	if repo.client == nil {
		t.Error("Expected a non-nil client in the repository")
	}

}
