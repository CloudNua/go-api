package comment

import (
	"context"
	"fmt"
)

// Data structure of comment service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that the service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service = is a struct on which all logic is built on top of
type Service struct {
	Store Store
}

// NewService = return pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
