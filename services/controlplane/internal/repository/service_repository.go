package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type ServiceRepository interface {
	Save(ctx context.Context, service models.Service) error
	GetByID(ctx context.Context, id string) (models.Service, error)
	ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Service, error)
}

type serviceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (sr *serviceRepository) Save(ctx context.Context, service models.Service) error {
	const query = `
		INSERT INTO services (id, org_id, owner_id, name)
		VALUES ($1, $2, $3, $4)
	`

	_, err := sr.db.ExecContext(ctx, query, service.ID, service.OrgID, service.OwnerID, service.Name)
	return err
}

func (sr *serviceRepository) GetByID(ctx context.Context, id string) (models.Service, error) {
	const query = `
		SELECT id, org_id, owner_id, name, created_at, updated_at
		FROM services
		WHERE id = $1
	`

	var service models.Service
	err := sr.db.QueryRowContext(ctx, query, id).Scan(
		&service.ID,
		&service.OrgID,
		&service.OwnerID,
		&service.Name,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		return models.Service{}, err
	}

	return service, nil
}

func (sr *serviceRepository) ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Service, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, owner_id, name, created_at, updated_at
		FROM services
		WHERE org_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := sr.db.QueryContext(ctx, query, orgID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	services := make([]models.Service, 0, limit)
	for rows.Next() {
		var service models.Service
		if err := rows.Scan(&service.ID, &service.OrgID, &service.OwnerID, &service.Name, &service.CreatedAt, &service.UpdatedAt); err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}
