package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Repository
}

func newService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	acct := &Account{
		ID:   ksuid.New().String(),
		Name: name,
	}
	if err := s.repository.PutAccount(ctx, *acct); err != nil {
		return nil, err
	}

	return acct, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return s.repository.ListAccounts(ctx, skip, take)
}
