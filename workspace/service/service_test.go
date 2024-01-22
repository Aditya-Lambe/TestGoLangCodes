package service

import (
	"testing"

	"github.com/AdiLambe/TestGoLangCodes/workspace/domain"
)

func TestNewOrderService(t *testing.T) {
	repo := domain.NeworderRepositoryDb()
	service := NewOrderService(repo)

	if service.GetRepository() == nil {
		t.Error("Expected a non-nil repository in the service")
	}

}
