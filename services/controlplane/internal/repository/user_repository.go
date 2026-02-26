package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type UserRepository interface {
	Save(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id string) (models.User, error)
	ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Save(ctx context.Context, user models.User) error {
	const query = `
		INSERT INTO users (id, org_id, email, name, role)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := ur.db.ExecContext(ctx, query, user.ID, user.OrgID, user.Email, user.Name, user.Role)
	return err
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (models.User, error) {
	const query = `
		SELECT id, org_id, email, name, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := ur.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.OrgID,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *userRepository) ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.User, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, email, name, role, created_at, updated_at
		FROM users
		WHERE org_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := ur.db.QueryContext(ctx, query, orgID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0, limit)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.OrgID, &user.Email, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
