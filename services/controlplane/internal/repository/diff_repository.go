package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type DiffRepository interface {
	Save(ctx context.Context, diff models.DiffResult) error
	ListByExperimentID(ctx context.Context, experimentID string, limit int) ([]models.DiffResult, error)
}

type diffRepository struct {
	db *sql.DB
}

func NewDiffRepository(db *sql.DB) DiffRepository {
	return &diffRepository{db: db}
}

func (dr *diffRepository) Save(ctx context.Context, diff models.DiffResult) error {
	const query = `
		INSERT INTO diff_results (id, org_id, experiment_id, route_id, request_id, verdict, field, primary_value, shadow_value)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := dr.db.ExecContext(ctx, query, diff.ID, diff.OrgID, diff.ExperimentID, diff.RouteID, diff.RequestID, diff.Verdict, diff.Field, diff.PrimaryValue, diff.ShadowValue)
	return err
}

func (dr *diffRepository) ListByExperimentID(ctx context.Context, experimentID string, limit int) ([]models.DiffResult, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, experiment_id, route_id, request_id, verdict, field, primary_value, shadow_value, created_at
		FROM diff_results
		WHERE experiment_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := dr.db.QueryContext(ctx, query, experimentID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	diffs := make([]models.DiffResult, 0, limit)
	for rows.Next() {
		var diff models.DiffResult
		if err := rows.Scan(&diff.ID, &diff.OrgID, &diff.ExperimentID, &diff.RouteID, &diff.RequestID, &diff.Verdict, &diff.Field, &diff.PrimaryValue, &diff.ShadowValue, &diff.CreatedAt); err != nil {
			return nil, err
		}
		diffs = append(diffs, diff)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return diffs, nil
}
