package repository

import (
	"context"

	"github.com/aaronlyy/go-api-example/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	DB *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) UsersRepository {
	return UsersRepository{DB: db}
}

func (r *UsersRepository) ListAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := pgxscan.Select(ctx, r.DB, &users,
		`SELECT uuid, username, active, created_at FROM users`,
	)
	return users, err
}
