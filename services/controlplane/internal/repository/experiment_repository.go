package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type ExperimentRepository interface {
	Save(ctx context.Context, experiment models.Experiment) error
	ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Experiment, error)
}

type experimentRepository struct {
	db *sql.DB
}

func NewExperimentRepository(db *sql.DB) ExperimentRepository {
	return &experimentRepository{db: db}
}

func (er *experimentRepository) Save(ctx context.Context, experiment models.Experiment) error {
	const query = `
		INSERT INTO experiments (id, org_id, route_id, name, status, started_at, ended_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := er.db.ExecContext(ctx, query, experiment.ID, experiment.OrgID, experiment.RouteID, experiment.Name, experiment.Status, experiment.StartedAt, experiment.EndedAt)
	return err
}

func (er *experimentRepository) ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Experiment, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, route_id, name, status, started_at, ended_at, created_at, updated_at
		FROM experiments
		WHERE org_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := er.db.QueryContext(ctx, query, orgID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	experiments := make([]models.Experiment, 0, limit)
	for rows.Next() {
		var experiment models.Experiment
		if err := rows.Scan(&experiment.ID, &experiment.OrgID, &experiment.RouteID, &experiment.Name, &experiment.Status, &experiment.StartedAt, &experiment.EndedAt, &experiment.CreatedAt, &experiment.UpdatedAt); err != nil {
			return nil, err
		}
		experiments = append(experiments, experiment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return experiments, nil
}
