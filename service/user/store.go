package user

import (
	"context"
	"fmt"

	"github.com/irviner26/ecom/types"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) types.UserStore {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string, ctx context.Context) (*types.User, error) {
	user := types.User{}

	err := s.db.QueryRow(ctx, "select id, first_name, last_name, email from users where email = $1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return &user, err
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Store) CreateUser(user types.User) error {
	panic("not implemented") // TODO: Implement
}
