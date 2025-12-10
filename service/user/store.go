package user

import (
	"context"

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

	err := s.db.QueryRow(ctx, "select id, first_name, last_name, email, password from users where email = $1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *Store) GetUserByID(id int, ctx context.Context) (*types.User, error) {
	user := types.User{}

	err := s.db.QueryRow(ctx, "select id, first_name, last_name, email from users where id = $1", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *Store) CreateUser(user types.User, ctx context.Context) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "insert into users(first_name, last_name, email, password) values ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
