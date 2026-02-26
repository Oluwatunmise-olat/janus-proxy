package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type RouteRepository interface {
	Save(ctx context.Context, route models.Route) error
	ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Route, error)
	ListByServiceID(ctx context.Context, serviceID string, limit int) ([]models.Route, error)
}

type routeRepository struct {
	db *sql.DB
}

func NewRouteRepository(db *sql.DB) RouteRepository {
	return &routeRepository{db: db}
}

func (rr *routeRepository) Save(ctx context.Context, route models.Route) error {
	const query = `
		INSERT INTO routes (id, org_id, service_id, method, path_pattern, is_shadow_enabled)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := rr.db.ExecContext(ctx, query, route.ID, route.OrgID, route.ServiceID, route.Method, route.PathPattern, route.IsShadowEnabled)
	return err
}

func (rr *routeRepository) ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Route, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, service_id, method, path_pattern, is_shadow_enabled, created_at, updated_at
		FROM routes
		WHERE org_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := rr.db.QueryContext(ctx, query, orgID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	routes := make([]models.Route, 0, limit)
	for rows.Next() {
		var route models.Route
		if err := rows.Scan(&route.ID, &route.OrgID, &route.ServiceID, &route.Method, &route.PathPattern, &route.IsShadowEnabled, &route.CreatedAt, &route.UpdatedAt); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}

func (rr *routeRepository) ListByServiceID(ctx context.Context, serviceID string, limit int) ([]models.Route, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, service_id, method, path_pattern, is_shadow_enabled, created_at, updated_at
		FROM routes
		WHERE service_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := rr.db.QueryContext(ctx, query, serviceID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	routes := make([]models.Route, 0, limit)
	for rows.Next() {
		var route models.Route
		if err := rows.Scan(&route.ID, &route.OrgID, &route.ServiceID, &route.Method, &route.PathPattern, &route.IsShadowEnabled, &route.CreatedAt, &route.UpdatedAt); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}
