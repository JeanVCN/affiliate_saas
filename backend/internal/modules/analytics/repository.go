package analytics

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ClickMetrics(ctx context.Context, workspaceID string, groupBy string) ([]ClickMetric, error)
	Overview(ctx context.Context, workspaceID string) (Overview, error)
	TopProducts(ctx context.Context, workspaceID string) ([]TopProduct, error)
	CreateConversionImport(ctx context.Context, workspaceID string, input CreateConversionImportInput) (ConversionImport, error)
	CreateConversionImportRow(ctx context.Context, workspaceID string, importID string, input CreateConversionImportRowInput) (ConversionImportRow, error)
	CreateConversionImportRows(ctx context.Context, workspaceID string, importID string, inputs []CreateConversionImportRowInput) ([]ConversionImportRow, error)
	GetConversionImport(ctx context.Context, workspaceID string, importID string) (ConversionImport, error)
	ReconciliationSummary(ctx context.Context, workspaceID string, importID string) (ReconciliationSummary, error)
	UpdateConversionImportRow(ctx context.Context, workspaceID string, importID string, rowID string, input UpdateConversionImportRowInput) (ConversionImportRow, error)
}

type queryRower interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ClickMetrics(ctx context.Context, workspaceID string, groupBy string) ([]ClickMetric, error) {
	query := `
		SELECT 'product' AS group_name, p.id, p.name, COUNT(*)::bigint
		FROM click_events ce
		JOIN products p ON p.id = ce.product_id
		WHERE ce.workspace_id = $1
		GROUP BY p.id, p.name
		ORDER BY COUNT(*) DESC, p.name`
	if groupBy == "link" {
		query = `
			SELECT 'link' AS group_name, al.id, COALESCE(al.label, al.destination_url), COUNT(*)::bigint
			FROM click_events ce
			JOIN affiliate_links al ON al.id = ce.affiliate_link_id
			WHERE ce.workspace_id = $1
			GROUP BY al.id, al.label, al.destination_url
			ORDER BY COUNT(*) DESC, COALESCE(al.label, al.destination_url)`
	}

	rows, err := repo.db.Query(ctx, query, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ClickMetric
	for rows.Next() {
		var item ClickMetric
		if err := rows.Scan(&item.Group, &item.GroupID, &item.GroupLabel, &item.Clicks); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) Overview(ctx context.Context, workspaceID string) (Overview, error) {
	var item Overview
	err := repo.db.QueryRow(ctx, `
		SELECT
			(SELECT COUNT(*)::bigint FROM click_events WHERE workspace_id = $1),
			(SELECT COUNT(*)::bigint FROM conversion_import_rows WHERE workspace_id = $1),
			COALESCE((SELECT SUM(gross_amount_cents)::bigint FROM conversion_import_rows WHERE workspace_id = $1), 0),
			COALESCE((SELECT SUM(commission_cents)::bigint FROM conversion_import_rows WHERE workspace_id = $1), 0)`,
		workspaceID,
	).Scan(&item.Clicks, &item.ImportedConversions, &item.GrossAmountCents, &item.CommissionCents)
	return item, err
}

func (repo *PostgresRepository) TopProducts(ctx context.Context, workspaceID string) ([]TopProduct, error) {
	rows, err := repo.db.Query(ctx, `
		WITH click_counts AS (
			SELECT product_id, COUNT(*)::bigint AS clicks
			FROM click_events
			WHERE workspace_id = $1 AND product_id IS NOT NULL
			GROUP BY product_id
		),
		import_counts AS (
			SELECT product_id,
			       COUNT(*)::bigint AS imported_conversions,
			       COALESCE(SUM(gross_amount_cents), 0)::bigint AS gross_amount_cents,
			       COALESCE(SUM(commission_cents), 0)::bigint AS commission_cents
			FROM conversion_import_rows
			WHERE workspace_id = $1 AND product_id IS NOT NULL
			GROUP BY product_id
		)
		SELECT p.id, p.name, COALESCE(cc.clicks, 0), COALESCE(ic.imported_conversions, 0),
		       COALESCE(ic.gross_amount_cents, 0), COALESCE(ic.commission_cents, 0)
		FROM products p
		LEFT JOIN click_counts cc ON cc.product_id = p.id
		LEFT JOIN import_counts ic ON ic.product_id = p.id
		WHERE p.workspace_id = $1
		  AND p.archived_at IS NULL
		  AND (COALESCE(cc.clicks, 0) > 0 OR COALESCE(ic.imported_conversions, 0) > 0)
		ORDER BY COALESCE(ic.commission_cents, 0) DESC,
		         COALESCE(ic.imported_conversions, 0) DESC,
		         COALESCE(cc.clicks, 0) DESC,
		         p.name
		LIMIT 10`, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []TopProduct
	for rows.Next() {
		var item TopProduct
		if err := rows.Scan(&item.ProductID, &item.ProductName, &item.Clicks, &item.ImportedConversions, &item.GrossAmountCents, &item.CommissionCents); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) CreateConversionImport(ctx context.Context, workspaceID string, input CreateConversionImportInput) (ConversionImport, error) {
	item := ConversionImport{ID: common.NewID("imp")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO conversion_imports (id, workspace_id, source)
		VALUES ($1, $2, $3)
		RETURNING id, workspace_id, source, status, created_at, updated_at`,
		item.ID, workspaceID, input.Source,
	).Scan(&item.ID, &item.WorkspaceID, &item.Source, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) CreateConversionImportRow(ctx context.Context, workspaceID string, importID string, input CreateConversionImportRowInput) (ConversionImportRow, error) {
	return createConversionImportRow(ctx, repo.db, workspaceID, importID, input)
}

func (repo *PostgresRepository) CreateConversionImportRows(ctx context.Context, workspaceID string, importID string, inputs []CreateConversionImportRowInput) ([]ConversionImportRow, error) {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	items := make([]ConversionImportRow, 0, len(inputs))
	for _, input := range inputs {
		item, err := createConversionImportRow(ctx, tx, workspaceID, importID, input)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return items, nil
}

func createConversionImportRow(ctx context.Context, q queryRower, workspaceID string, importID string, input CreateConversionImportRowInput) (ConversionImportRow, error) {
	if input.RawPayload == nil {
		input.RawPayload = map[string]any{}
	}
	rawPayload, err := json.Marshal(input.RawPayload)
	if err != nil {
		return ConversionImportRow{}, err
	}
	item := ConversionImportRow{ID: common.NewID("cir")}
	var occurredAt sql.NullTime
	var reconciledAt sql.NullTime
	err = q.QueryRow(ctx, `
		WITH resolved AS (
			SELECT
				COALESCE(NULLIF($4, ''), (
					SELECT al.product_id
					FROM affiliate_links al
					WHERE al.id = NULLIF($5, '')
					  AND al.workspace_id = $2
					  AND al.archived_at IS NULL
				)) AS product_id,
				NULLIF($5, '') AS affiliate_link_id
		)
		INSERT INTO conversion_import_rows (
			id, workspace_id, conversion_import_id, product_id, affiliate_link_id,
			occurred_at, order_reference, gross_amount_cents, commission_cents, currency, raw_payload,
			reconciliation_status
		)
		SELECT $1, $2, ci.id, r.product_id, r.affiliate_link_id, $6, NULLIF($7, ''),
		       $8, $9, NULLIF($10, ''), $11,
		       CASE
		         WHEN r.product_id IS NOT NULL OR r.affiliate_link_id IS NOT NULL THEN 'matched'
		         ELSE 'pending'
		       END
		FROM conversion_imports ci
		CROSS JOIN resolved r
		WHERE ci.id = $3
		  AND ci.workspace_id = $2
		  AND ci.status <> 'archived'
		  AND (
		    r.product_id IS NULL
		    OR EXISTS (
		      SELECT 1 FROM products p
		      WHERE p.id = r.product_id AND p.workspace_id = $2 AND p.archived_at IS NULL
		    )
		  )
		  AND (
		    r.affiliate_link_id IS NULL
		    OR EXISTS (
		      SELECT 1 FROM affiliate_links al
		      WHERE al.id = r.affiliate_link_id
		        AND al.workspace_id = $2
		        AND al.archived_at IS NULL
		        AND (r.product_id IS NULL OR al.product_id = r.product_id)
		    )
		  )
		RETURNING id, workspace_id, conversion_import_id, COALESCE(product_id, ''),
		          COALESCE(affiliate_link_id, ''), occurred_at, COALESCE(order_reference, ''),
		          gross_amount_cents, commission_cents, COALESCE(currency, ''), raw_payload,
		          reconciliation_status, COALESCE(reconciliation_note, ''), reconciled_at, created_at`,
		item.ID, workspaceID, importID, input.ProductID, input.AffiliateLinkID, input.OccurredAt,
		input.OrderReference, input.GrossAmountCents, input.CommissionCents, input.Currency, rawPayload,
	).Scan(&item.ID, &item.WorkspaceID, &item.ConversionImportID, &item.ProductID, &item.AffiliateLinkID,
		&occurredAt, &item.OrderReference, &item.GrossAmountCents, &item.CommissionCents,
		&item.Currency, &item.RawPayload, &item.ReconciliationStatus, &item.ReconciliationNote,
		&reconciledAt, &item.CreatedAt)
	if occurredAt.Valid {
		item.OccurredAt = &occurredAt.Time
	}
	if reconciledAt.Valid {
		item.ReconciledAt = &reconciledAt.Time
	}
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetConversionImport(ctx context.Context, workspaceID string, importID string) (ConversionImport, error) {
	var item ConversionImport
	err := repo.db.QueryRow(ctx, `
		SELECT id, workspace_id, source, status, created_at, updated_at
		FROM conversion_imports
		WHERE workspace_id = $1 AND id = $2 AND status <> 'archived'`,
		workspaceID, importID,
	).Scan(&item.ID, &item.WorkspaceID, &item.Source, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return item, common.NormalizePostgresErr(err)
	}
	rows, err := repo.listConversionImportRows(ctx, workspaceID, importID)
	if err != nil {
		return item, err
	}
	item.Rows = rows
	summary, err := repo.ReconciliationSummary(ctx, workspaceID, importID)
	if err != nil {
		return item, err
	}
	item.Summary = &summary
	return item, nil
}

func (repo *PostgresRepository) listConversionImportRows(ctx context.Context, workspaceID string, importID string) ([]ConversionImportRow, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, conversion_import_id, COALESCE(product_id, ''),
		       COALESCE(affiliate_link_id, ''), occurred_at, COALESCE(order_reference, ''),
		       gross_amount_cents, commission_cents, COALESCE(currency, ''), raw_payload,
		       reconciliation_status, COALESCE(reconciliation_note, ''), reconciled_at, created_at
		FROM conversion_import_rows
		WHERE workspace_id = $1 AND conversion_import_id = $2
		ORDER BY created_at DESC`, workspaceID, importID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ConversionImportRow
	for rows.Next() {
		var item ConversionImportRow
		var occurredAt sql.NullTime
		var reconciledAt sql.NullTime
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.ConversionImportID, &item.ProductID, &item.AffiliateLinkID,
			&occurredAt, &item.OrderReference, &item.GrossAmountCents, &item.CommissionCents,
			&item.Currency, &item.RawPayload, &item.ReconciliationStatus, &item.ReconciliationNote,
			&reconciledAt, &item.CreatedAt); err != nil {
			return nil, err
		}
		if occurredAt.Valid {
			item.OccurredAt = &occurredAt.Time
		}
		if reconciledAt.Valid {
			item.ReconciledAt = &reconciledAt.Time
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
