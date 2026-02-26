package repository

import (
	"context"
	"database/sql"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/models"
)

type AuditLogRepository interface {
	Save(ctx context.Context, auditLog models.AuditLog) error
	ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.AuditLog, error)
}

type auditLogRepository struct {
	db *sql.DB
}

func NewAuditLogRepository(db *sql.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (al *auditLogRepository) Save(ctx context.Context, auditLog models.AuditLog) error {
	const query = `
		INSERT INTO audit_logs (id, org_id, actor_id, name, description, meta)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := al.db.ExecContext(
		ctx,
		query,
		auditLog.ID,
		auditLog.OrgID,
		auditLog.ActorID,
		auditLog.Name,
		auditLog.Description,
		auditLog.Meta,
	)

	return err
}

func (al *auditLogRepository) ListByOrgID(ctx context.Context, orgID string, limit int) ([]models.AuditLog, error) {
	if limit <= 0 {
		limit = 50
	}

	const query = `
		SELECT id, org_id, actor_id, name, description, meta, created_at, updated_at
		FROM audit_logs
		WHERE org_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := al.db.QueryContext(ctx, query, orgID, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	logs := make([]models.AuditLog, 0, limit)

	for rows.Next() {
		var log models.AuditLog

		if err := rows.Scan(
			&log.ID,
			&log.OrgID,
			&log.ActorID,
			&log.Name,
			&log.Description,
			&log.Meta,
			&log.CreatedAt,
			&log.UpdatedAt,
		); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
