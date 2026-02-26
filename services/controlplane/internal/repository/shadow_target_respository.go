package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type ShadowTargetRepository interface {
	Save(ctx context.Context, target models.ShadowTarget) error
	ListByRouteID(ctx context.Context, routeID string, limit int) ([]models.ShadowTarget, error)
}

type shadowTargetRepository struct {
	db *sql.DB
}

func NewShadowTargetRepository(db *sql.DB) ShadowTargetRepository {
	return &shadowTargetRepository{db: db}
}

func (sr *shadowTargetRepository) Save(ctx context.Context, target models.ShadowTarget) error {
	const query = `
		INSERT INTO shadow_targets (id, org_id, route_id, target_service, weight, timeout_ms, retries, enabled)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := sr.db.ExecContext(ctx, query, target.ID, target.OrgID, target.RouteID, target.TargetServiceID, target.Weight, target.TimeoutMs, target.Retries, target.Enabled)
	return err
}

func (sr *shadowTargetRepository) ListByRouteID(ctx context.Context, routeID string, limit int) ([]models.ShadowTarget, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, route_id, target_service, weight, timeout_ms, retries, enabled, created_at, updated_at
		FROM shadow_targets
		WHERE route_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := sr.db.QueryContext(ctx, query, routeID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	targets := make([]models.ShadowTarget, 0, limit)
	for rows.Next() {
		var target models.ShadowTarget
		if err := rows.Scan(&target.ID, &target.OrgID, &target.RouteID, &target.TargetServiceID, &target.Weight, &target.TimeoutMs, &target.Retries, &target.Enabled, &target.CreatedAt, &target.UpdatedAt); err != nil {
			return nil, err
		}
		targets = append(targets, target)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return targets, nil
}
