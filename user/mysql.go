package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aarondl/authboss/v3"
)

type MysqlUserRepo struct {
	db *sql.DB
}

func NewMysqlUserRepo(db *sql.DB) *MysqlUserRepo {

	return &MysqlUserRepo{db: db}
}

func (repo *MysqlUserRepo) Load(ctx context.Context, key string) (authboss.User, error) {
	query := `SELECT email, password FROM users WHERE email = ?`

	row := repo.db.QueryRowContext(ctx, query, key)

	u := &User{}
	err := row.Scan(&u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, authboss.ErrUserNotFound
		}
		return nil, err
	}

	return u, nil
}

func (repo *MysqlUserRepo) Save(ctx context.Context, user authboss.User) error {
	u := user.(*User)

	query := `UPDATE users SET email = ?, password = ? where email=u.Email`

	_, err := repo.db.ExecContext(ctx, query, u.Email, u.Password)
	return err
}

func (repo *MysqlUserRepo) New(ctx context.Context) authboss.User {

	return NewUser()
}

func (repo *MysqlUserRepo) Create(ctx context.Context, user authboss.User) error {

	fmt.Println("here we are in the create ", user)

	u := user.(*User)

	query := `INSERT INTO users ( email, password) VALUES (?, ?)`

	if repo.db == nil {

		fmt.Println("here we go folks ")
	}

	_, err := repo.db.ExecContext(ctx, query, u.Email, u.Password)
	return err

}
