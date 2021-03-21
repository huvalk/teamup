package repository

import (
	"context"
	"diplomaProject/application/auth"
	"diplomaProject/application/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) auth.Repository {
	return &Repository{conn: db}
}

func (r Repository) UpdateUserInfo(user *models.User) error {
	sql := `INSERT INTO users (id, firstname, lastname, email) 
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id)
			DO
				UPDATE SET firstname = EXCLUDED.firstname,
				lastname = EXCLUDED.lastname,
				email = EXCLUDED.email`

	_, err := r.conn.Exec(context.Background(),
		sql, user.Id, user.FirstName, user.LastName, user.Email)
	return err
}
