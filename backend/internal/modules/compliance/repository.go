package compliance

import (
	"context"
	"strings"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CampaignContent(ctx context.Context, workspaceID string, campaignID string) (string, error)
	CreateCampaignCheck(ctx context.Context, workspaceID string, campaignID string, findings []FindingInput) (Check, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) CampaignContent(ctx context.Context, workspaceID string, campaignID string) (string, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT COALESCE(cp.title, ''), COALESCE(cp.body, '')
		FROM campaigns c
		LEFT JOIN channel_packages cp ON cp.campaign_id = c.id AND cp.workspace_id = c.workspace_id
		WHERE c.id = $2 AND c.workspace_id = $1 AND c.archived_at IS NULL`,
		workspaceID, campaignID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var content strings.Builder
	found := false
	for rows.Next() {
		found = true
		var title string
		var body string
		if err := rows.Scan(&title, &body); err != nil {
			return "", err
		}
		content.WriteString(" ")
		content.WriteString(title)
		content.WriteString(" ")
		content.WriteString(body)
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	if !found {
		return "", common.ErrNotFound
	}
	return content.String(), nil
}

func (repo *PostgresRepository) CreateCampaignCheck(ctx context.Context, workspaceID string, campaignID string, findings []FindingInput) (Check, error) {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return Check{}, err
	}
	defer tx.Rollback(ctx)

	item := Check{ID: common.NewID("chk")}
	err = tx.QueryRow(ctx, `
		INSERT INTO compliance_checks (id, workspace_id, campaign_id, status)
		SELECT $1, $2, c.id, 'completed'
		FROM campaigns c
		WHERE c.id = $3 AND c.workspace_id = $2 AND c.archived_at IS NULL
		RETURNING id, workspace_id, COALESCE(campaign_id, ''), status, created_at`,
		item.ID, workspaceID, campaignID,
	).Scan(&item.ID, &item.WorkspaceID, &item.CampaignID, &item.Status, &item.CreatedAt)
	if err != nil {
		return item, common.NormalizePostgresErr(err)
	}

	for _, findingInput := range findings {
		finding := Finding{ID: common.NewID("fnd")}
		err = tx.QueryRow(ctx, `
			INSERT INTO compliance_findings (id, workspace_id, compliance_check_id, severity, code, message)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, workspace_id, compliance_check_id, severity, code, message, created_at`,
			finding.ID, workspaceID, item.ID, findingInput.Severity, findingInput.Code, findingInput.Message,
		).Scan(&finding.ID, &finding.WorkspaceID, &finding.ComplianceCheckID, &finding.Severity, &finding.Code, &finding.Message, &finding.CreatedAt)
		if err != nil {
			return item, common.NormalizePostgresErr(err)
		}
		item.Findings = append(item.Findings, finding)
	}
	if err := tx.Commit(ctx); err != nil {
		return item, err
	}
	return item, nil
}
