package analytics

import (
	"context"
	"database/sql"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

func (repo *PostgresRepository) ReconciliationSummary(ctx context.Context, workspaceID string, importID string) (ReconciliationSummary, error) {
	var item ReconciliationSummary
	err := repo.db.QueryRow(ctx, `
		SELECT COUNT(*)::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'pending')::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'matched')::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'unmatched')::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'ignored')::bigint
		FROM conversion_import_rows
		WHERE workspace_id = $1 AND conversion_import_id = $2`,
		workspaceID, importID,
	).Scan(&item.Total, &item.Pending, &item.Matched, &item.Unmatched, &item.Ignored)
	return item, err
}

func (repo *PostgresRepository) UpdateConversionImportRow(ctx context.Context, workspaceID string, importID string, rowID string, input UpdateConversionImportRowInput) (ConversionImportRow, error) {
	var item ConversionImportRow
	var occurredAt sql.NullTime
	var reconciledAt sql.NullTime
	err := repo.db.QueryRow(ctx, `
		WITH resolved AS (
			SELECT
				COALESCE(NULLIF($4, ''), (
					SELECT al.product_id
					FROM affiliate_links al
					WHERE al.id = NULLIF($5, '')
					  AND al.workspace_id = $1
					  AND al.archived_at IS NULL
				)) AS product_id,
				NULLIF($5, '') AS affiliate_link_id
		)
		UPDATE conversion_import_rows cir
		SET product_id = COALESCE(r.product_id, cir.product_id),
		    affiliate_link_id = COALESCE(r.affiliate_link_id, cir.affiliate_link_id),
		    reconciliation_status = COALESCE($6, CASE
		      WHEN COALESCE(r.product_id, cir.product_id) IS NOT NULL
		        OR COALESCE(r.affiliate_link_id, cir.affiliate_link_id) IS NOT NULL
		      THEN 'matched'
		      ELSE cir.reconciliation_status
		    END),
		    reconciliation_note = COALESCE($7, cir.reconciliation_note),
		    reconciled_at = CASE
		      WHEN $6 IS NOT NULL OR $4 IS NOT NULL OR $5 IS NOT NULL OR $7 IS NOT NULL THEN now()
		      ELSE cir.reconciled_at
		    END
		FROM resolved r
		WHERE cir.workspace_id = $1
		  AND cir.conversion_import_id = $2
		  AND cir.id = $3
		  AND (
		    r.product_id IS NULL
		    OR EXISTS (
		      SELECT 1 FROM products p
		      WHERE p.id = r.product_id AND p.workspace_id = $1 AND p.archived_at IS NULL
		    )
		  )
		  AND (
		    COALESCE(r.affiliate_link_id, cir.affiliate_link_id) IS NULL
		    OR EXISTS (
		      SELECT 1 FROM affiliate_links al
		      WHERE al.id = COALESCE(r.affiliate_link_id, cir.affiliate_link_id)
		        AND al.workspace_id = $1
		        AND al.archived_at IS NULL
		        AND al.product_id = COALESCE(r.product_id, cir.product_id)
		    )
		  )
		RETURNING id, workspace_id, conversion_import_id, COALESCE(product_id, ''),
		          COALESCE(affiliate_link_id, ''), occurred_at, COALESCE(order_reference, ''),
		          gross_amount_cents, commission_cents, COALESCE(currency, ''), raw_payload,
		          reconciliation_status, COALESCE(reconciliation_note, ''), reconciled_at, created_at`,
		workspaceID, importID, rowID, input.ProductID, input.AffiliateLinkID, input.ReconciliationStatus, input.ReconciliationNote,
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
