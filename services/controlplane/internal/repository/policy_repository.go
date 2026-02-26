package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type PolicyRepository interface {
	Save(ctx context.Context, policy models.Policy) error
	ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Policy, error)
}

type policyRepository struct {
	db *sql.DB
}

func NewPolicyRepository(db *sql.DB) PolicyRepository {
	return &policyRepository{db: db}
}

func (pr *policyRepository) Save(ctx context.Context, policy models.Policy) error {
	const query = `
		INSERT INTO policies (id, org_id, name, mismatch_max, latency_max_ms, enabled)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := pr.db.ExecContext(ctx, query, policy.ID, policy.OrgID, policy.Name, policy.MismatchMax, policy.LatencyMaxMs, policy.Enabled)
	return err
}

func (pr *policyRepository) ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.Policy, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, name, mismatch_max, latency_max_ms, enabled, created_at, updated_at
		FROM policies
		WHERE org_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := pr.db.QueryContext(ctx, query, orgID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	policies := make([]models.Policy, 0, limit)
	for rows.Next() {
		var policy models.Policy
		if err := rows.Scan(&policy.ID, &policy.OrgID, &policy.Name, &policy.MismatchMax, &policy.LatencyMaxMs, &policy.Enabled, &policy.CreatedAt, &policy.UpdatedAt); err != nil {
			return nil, err
		}
		policies = append(policies, policy)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return policies, nil
}
